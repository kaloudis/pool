package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/btcsuite/btcutil"
	"github.com/lightninglabs/llm"
	"github.com/lightninglabs/llm/clmrpc"
	"github.com/lightninglabs/protobuf-hex-display/jsonpb"
	"github.com/lightninglabs/protobuf-hex-display/proto"
	"github.com/lightningnetwork/lnd/macaroons"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"gopkg.in/macaroon.v2"
)

var (
	// maxMsgRecvSize is the largest message our client will receive. We
	// set this to 200MiB atm.
	maxMsgRecvSize = grpc.MaxCallRecvMsgSize(1 * 1024 * 1024 * 200)

	// defaultMacaroonTimeout is the default macaroon timeout in seconds
	// that we set when sending it over the line.
	defaultMacaroonTimeout int64 = 60

	tlsCertFlag = cli.StringFlag{
		Name: "tlscertpath",
		Usage: "path to llm's TLS certificate, only needed if llm " +
			"runs in the same process as lnd",
	}
	macaroonPathFlag = cli.StringFlag{
		Name: "macaroonpath",
		Usage: "path to macaroon file, only needed if llm runs " +
			"in the same process as lnd",
	}
)

type invalidUsageError struct {
	ctx     *cli.Context
	command string
}

func (e *invalidUsageError) Error() string {
	return fmt.Sprintf("invalid usage of command %s", e.command)
}

func printJSON(resp interface{}) {
	b, err := json.Marshal(resp)
	if err != nil {
		fatal(err)
	}

	var out bytes.Buffer
	_ = json.Indent(&out, b, "", "\t")
	out.WriteString("\n")
	_, _ = out.WriteTo(os.Stdout)
}

func printRespJSON(resp proto.Message) { // nolint
	jsonMarshaler := &jsonpb.Marshaler{
		EmitDefaults: true,
		OrigName:     true,
		Indent:       "\t", // Matches indentation of printJSON.
	}

	jsonStr, err := jsonMarshaler.MarshalToString(resp)
	if err != nil {
		fmt.Println("unable to decode response: ", err)
		return
	}

	fmt.Println(jsonStr)
}

func fatal(err error) {
	var e *invalidUsageError
	if errors.As(err, &e) {
		_ = cli.ShowCommandHelp(e.ctx, e.command)
	} else {
		_, _ = fmt.Fprintf(os.Stderr, "[llm] %v\n", err)
	}
	os.Exit(1)
}

func main() {
	app := cli.NewApp()

	app.Version = llm.Version()
	app.Name = "llm"
	app.Usage = "control plane for your llmd"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "rpcserver",
			Value: "localhost:12010",
			Usage: "llmd daemon address host:port",
		},
		tlsCertFlag,
		macaroonPathFlag,
	}
	app.Commands = append(app.Commands, accountsCommands...)
	app.Commands = append(app.Commands, ordersCommands...)
	app.Commands = append(app.Commands, auctionCommands...)
	app.Commands = append(app.Commands, listAuthCommand)

	err := app.Run(os.Args)
	if err != nil {
		fatal(err)
	}
}

func getClient(ctx *cli.Context) (clmrpc.TraderClient, func(),
	error) {

	rpcServer := ctx.GlobalString("rpcserver")
	tlsCertPath := ctx.GlobalString(tlsCertFlag.Name)
	macaroonPath := ctx.GlobalString(macaroonPathFlag.Name)
	conn, err := getClientConn(rpcServer, tlsCertPath, macaroonPath)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() { _ = conn.Close() }

	traderClient := clmrpc.NewTraderClient(conn)
	return traderClient, cleanup, nil
}

func parseAmt(text string) (btcutil.Amount, error) {
	amtInt64, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid amt value: %v", err)
	}
	return btcutil.Amount(amtInt64), nil
}

func getClientConn(address, tlsCertPath, macaroonPath string) (*grpc.ClientConn,
	error) {

	opts := []grpc.DialOption{
		grpc.WithDefaultCallOptions(maxMsgRecvSize),
	}

	switch {
	// If a TLS certificate file is specified, we need to load it and build
	// transport credentials with it.
	case tlsCertPath != "":
		creds, err := credentials.NewClientTLSFromFile(tlsCertPath, "")
		if err != nil {
			fatal(err)
		}

		// Macaroons are only allowed to be transmitted over a TLS
		// enabled connection.
		if macaroonPath != "" {
			opts = append(opts, readMacaroon(macaroonPath))
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))

	// By default, if no certificate is supplied, we assume the RPC server
	// runs without TLS.
	default:
		opts = append(opts, grpc.WithInsecure())
	}

	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to RPC server: %v",
			err)
	}

	return conn, nil
}

func parseStr(ctx *cli.Context, argIdx int, flag, cmd string) (string, error) {
	var str string
	switch {
	case ctx.IsSet(flag):
		str = ctx.String(flag)
	case ctx.Args().Get(argIdx) != "":
		str = ctx.Args().Get(argIdx)
	default:
		return "", &invalidUsageError{ctx, cmd}
	}
	return str, nil
}

func parseHexStr(ctx *cli.Context, argIdx int, flag, cmd string) ([]byte, error) { // nolint:unparam
	hexStr, err := parseStr(ctx, argIdx, flag, cmd)
	if err != nil {
		return nil, err
	}
	return hex.DecodeString(hexStr)
}

func parseUint64(ctx *cli.Context, argIdx int, flag, cmd string) (uint64, error) {
	str, err := parseStr(ctx, argIdx, flag, cmd)
	if err != nil {
		return 0, err
	}
	return strconv.ParseUint(str, 10, 64)
}

// readMacaroon tries to read the macaroon file at the specified path and create
// gRPC dial options from it.
func readMacaroon(macPath string) grpc.DialOption {
	// Load the specified macaroon file.
	macBytes, err := ioutil.ReadFile(macPath)
	if err != nil {
		fatal(fmt.Errorf("unable to read macaroon path : %v", err))
	}

	mac := &macaroon.Macaroon{}
	if err = mac.UnmarshalBinary(macBytes); err != nil {
		fatal(fmt.Errorf("unable to decode macaroon: %v", err))
	}

	macConstraints := []macaroons.Constraint{
		// We add a time-based constraint to prevent replay of the
		// macaroon. It's good for 60 seconds by default to make up for
		// any discrepancy between client and server clocks, but leaking
		// the macaroon before it becomes invalid makes it possible for
		// an attacker to reuse the macaroon. In addition, the validity
		// time of the macaroon is extended by the time the server clock
		// is behind the client clock, or shortened by the time the
		// server clock is ahead of the client clock (or invalid
		// altogether if, in the latter case, this time is more than 60
		// seconds).
		macaroons.TimeoutConstraint(defaultMacaroonTimeout),
	}

	// Apply constraints to the macaroon.
	constrainedMac, err := macaroons.AddConstraints(mac, macConstraints...)
	if err != nil {
		fatal(err)
	}

	// Now we append the macaroon credentials to the dial options.
	cred := macaroons.NewMacaroonCredential(constrainedMac)
	return grpc.WithPerRPCCredentials(cred)
}
