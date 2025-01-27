package main

import (
	"awesomeProject/accounts/dto"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
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
	portVal := flag.Int("port", 8080, "server port")
	hostVal := flag.String("host", "0.0.0.0", "server host")
	cmdVal := flag.String("cmd", "", "command to execute")
	nameVal := flag.String("name", "c", "name of account")
	nameNewVal := flag.String("name_new", "", "new name of account")
	amountVal := flag.Int("amount", 0, "amount of account")

	flag.Parse()

	c := http.Client{Timeout: time.Duration(1) * time.Second}

	cmd := Command{
		Port:    *portVal,
		Host:    *hostVal,
		Cmd:     *cmdVal,
		Name:    *nameVal,
		NameNew: *nameNewVal,
		Amount:  *amountVal,
	}

	if err := do(cmd, c); err != nil {
		panic(err)
	}
}

func do(cmd Command, c http.Client) error {
	switch cmd.Cmd {
	case "create":
		if err := create(cmd, c); err != nil {
			return fmt.Errorf("create account failed: %w", err)
		}

		return nil
	case "get":
		if err := get(cmd, c); err != nil {
			return fmt.Errorf("get account failed: %w", err)
		}

		return nil
	case "delete":
		if err := srvDelete(cmd, c); err != nil {
			return fmt.Errorf("get account failed: %w", err)
		}

		return nil
	case "patch":
		if err := patch(cmd, c); err != nil {
			return fmt.Errorf("get account failed: %w", err)
		}

		return nil
	case "change":
		if err := change(cmd, c); err != nil {
			return fmt.Errorf("get account failed: %w", err)
		}

		return nil

	default:
		return fmt.Errorf("unknown command %s", cmd.Cmd)
	}
}

func get(cmd Command, c http.Client) error {
	resp, err := c.Get(
		fmt.Sprintf("http://%s:%d/account?name=%s", cmd.Host, cmd.Port, cmd.Name),
	)
	if err != nil {
		return fmt.Errorf("http get failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("read body failed: %w", err)
		}

		return fmt.Errorf("resp error %s", string(body))
	}

	var response dto.GetAccountResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("json decode failed: %w", err)
	}

	fmt.Printf("response account name: %s and amount: %d", response.Name, response.Amount)

	return nil
}

func create(cmd Command, c http.Client) error {
	request := dto.CreateAccountRequest{
		Name:   cmd.Name,
		Amount: cmd.Amount,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	resp, err := c.Post(
		fmt.Sprintf("http://%s:%d/account/create", cmd.Host, cmd.Port),
		"application/json",
		bytes.NewReader(data),
	)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func srvDelete(cmd Command, c http.Client) error {
	request := dto.DeleteAccountRequest{
		Name: cmd.Name,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	resp, err := c.Post(
		fmt.Sprintf("http://%s:%d/account/delete", cmd.Host, cmd.Port),
		"application/json",
		bytes.NewReader(data),
	)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func patch(cmd Command, c http.Client) error {
	request := dto.PatchAccountRequest{
		Name:   cmd.Name,
		Amount: cmd.Amount,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	resp, err := c.Post(
		fmt.Sprintf("http://%s:%d/account/patch", cmd.Host, cmd.Port),
		"application/json",
		bytes.NewReader(data),
	)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func change(cmd Command, c http.Client) error {
	request := dto.ChangeAccountRequest{
		Name:    cmd.Name,
		NameNew: cmd.NameNew,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	resp, err := c.Post(
		fmt.Sprintf("http://%s:%d/account/change", cmd.Host, cmd.Port),
		"application/json",
		bytes.NewReader(data),
	)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}
