package main

import (
	"context"
	"flag"
	"fmt"
	"gRPCProject/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type Command struct {
	Port    int
	Host    string
	Cmd     string
	Name    string
	NameNew string
	Amount  int
}

func Use(vals ...interface{}) { _ = vals }

func (c *Command) Do() error {
	switch c.Cmd {
	case "create":
		return c.create()
	default:
		return fmt.Errorf("unknown command: %s", c.Cmd)
	}
}

func (c *Command) create() error {
	panic("implement me")
}

func main() {
	Use(Use)
	portVal := flag.Int("port", 4567, "server port")
	hostVal := flag.String("host", "0.0.0.0", "server host")
	cmdVal := flag.String("cmd", "", "command to execute")
	nameVal := flag.String("name", "c", "name of account")
	nameNewVal := flag.String("name_new", "", "new name of account")
	amountVal := flag.Int("amount", 0, "amount of account")

	flag.Parse()

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", *hostVal, *portVal),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = conn.Close()
	}()

	c := proto.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	cmd := Command{
		Port:    *portVal,
		Host:    *hostVal,
		Cmd:     *cmdVal,
		Name:    *nameVal,
		NameNew: *nameNewVal,
		Amount:  *amountVal,
	}

	if err = do(cmd, c, ctx); err != nil {
		panic(err)
	}

}

func do(cmd Command, c proto.GreeterClient, ctx context.Context) error {
	switch cmd.Cmd {
	case "create":
		if _, err := c.CreateAccount(ctx, &proto.CreateAccountRequest{Name: cmd.Name, Amount: int32(cmd.Amount)}); err != nil {
			return fmt.Errorf("create account failed: %w", err)
		}
		return nil
	case "get":
		response, err := c.GetAccount(ctx, &proto.GetAccountRequest{Name: cmd.Name})
		if err != nil {
			return fmt.Errorf("create account failed: %w", err)
		}
		fmt.Printf("response account name: %s and amount: %d", response.Name, response.Amount)
		return nil
	case "delete":
		if _, err := c.DeleteAccount(ctx, &proto.DeleteAccountRequest{Name: cmd.Name}); err != nil {
			return fmt.Errorf("delete account failed: %w", err)
		}
		return nil
	case "patch":
		if _, err := c.PatchAccount(ctx, &proto.PatchAccountRequest{Name: cmd.Name, Amount: int32(cmd.Amount)}); err != nil {
			return fmt.Errorf("patch account failed: %w", err)
		}
		return nil
	case "change":
		if _, err := c.ChangeAccount(ctx, &proto.ChangeAccountRequest{Name: cmd.Name, NameNew: cmd.NameNew}); err != nil {
			return fmt.Errorf("change account failed: %w", err)
		}
		return nil

	default:
		return fmt.Errorf("unknown command %s", cmd.Cmd)
	}
}
