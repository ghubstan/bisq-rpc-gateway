package client

import (
	"fmt"
	pb "golang.org/x/bisq-grpc-gateway/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"time"
)

type passwordHeader struct {
	password string
}

// Return value is mapped to request headers.
func (h passwordHeader) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	return map[string]string{
		"password": h.password,
	}, nil
}

func (passwordHeader) RequireTransportSecurity() bool {
	return false
}

func runCommand(ctx context.Context, client pb.MessageServiceClient, cmd string) error {
	fmt.Printf("RPC Call with cmd tokens (%q)\n", cmd)
	r, err := client.Call(ctx, &pb.Command{Params: cmd})
	if err != nil {
		log.Fatalf("Could not call RPC server: %v", err)
		return err
	}
	log.Printf("Response.Result: %s", r.GetResult())
	return nil
}

func RunClient(serverAddress string, params string) error {

	// Set up connection to the server with the bisq daemon password (hardcoded for now).
	conn, err := grpc.DialContext(context.TODO(), serverAddress, grpc.WithInsecure(),
		grpc.WithPerRPCCredentials(passwordHeader{
			password: "xyz",
		}),
	)

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMessageServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if len(params) > 0 {
		err := runCommand(ctx, c, params)
		if err != nil {
			return err
		}
	}

	// TODO gRPC client should not wrap request cmd tokens in json, but http client has to, via http req.body.
	var cmd = [8]string{
		"",
		"getversion",
		"getbalance",
		"setwalletpassword oldwalletpassword \"a brand new walletpassword\"",
		"unlockwallet \"oldwalletpassword\" 60",
		"lockwallet",
		"setwalletpassword oldwalletpassword \"a brand new walletpassword\"",
		"removewalletpassword \"a brand new walletpassword\"",
	}

	var i int
	for i = 0; i < len(cmd); i++ {
		err = runCommand(ctx, c, cmd[i])
		if err != nil {
			return err
		}
	}

	return nil
}
