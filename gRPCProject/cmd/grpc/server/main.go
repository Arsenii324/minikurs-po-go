package main

import (
	"context"
	"errors"
	"fmt"
	"gRPCProject/accounts"
	"gRPCProject/accounts/models"
	"gRPCProject/proto"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	proto.UnimplementedGreeterServer
}

var (
	h accounts.Handler
)

func (s *server) CreateAccount(_ context.Context, request *proto.CreateAccountRequest) (*proto.BaseReply, error) {

	if len(request.Name) == 0 {
		return nil, errors.New("empty name")
	}

	h.Guard.Lock()

	if _, ok := h.Accounts[request.Name]; ok {
		h.Guard.Unlock()

		return nil, errors.New("account already exists")
	}

	h.Accounts[request.Name] = &models.Account{
		Name:   request.Name,
		Amount: int(request.Amount),
	}

	h.Guard.Unlock()

	return &proto.BaseReply{State: "OK"}, nil
}

func (s *server) PatchAccount(_ context.Context, request *proto.PatchAccountRequest) (*proto.BaseReply, error) {

	if len(request.Name) == 0 {
		return nil, errors.New("empty name")
	}

	h.Guard.Lock()

	if _, ok := h.Accounts[request.Name]; !ok {
		h.Guard.Unlock()

		return nil, errors.New("account does not exist")
	}

	h.Accounts[request.Name].Amount = int(request.Amount)

	h.Guard.Unlock()

	return &proto.BaseReply{State: "OK"}, nil
}

func (s *server) DeleteAccount(_ context.Context, request *proto.DeleteAccountRequest) (*proto.BaseReply, error) {

	if len(request.Name) == 0 {
		return nil, errors.New("empty name")
	}

	h.Guard.Lock()

	if _, ok := h.Accounts[request.Name]; !ok {
		h.Guard.Unlock()

		return nil, errors.New("account does not exist")
	}

	delete(h.Accounts, request.Name)

	h.Guard.Unlock()

	return &proto.BaseReply{State: "OK"}, nil
}

func (s *server) ChangeAccount(_ context.Context, request *proto.ChangeAccountRequest) (*proto.BaseReply, error) {

	if len(request.Name) == 0 {
		return nil, errors.New("empty name")
	}

	if len(request.NameNew) == 0 {
		return nil, errors.New("empty new name")
	}

	h.Guard.Lock()

	if _, ok := h.Accounts[request.Name]; !ok {
		h.Guard.Unlock()

		return nil, errors.New("account does not exist")
	}

	if _, ok := h.Accounts[request.NameNew]; ok {
		h.Guard.Unlock()

		return nil, errors.New("account with new name already exists")
	}

	h.Accounts[request.NameNew] = h.Accounts[request.Name]
	delete(h.Accounts, request.Name)

	h.Guard.Unlock()

	return &proto.BaseReply{State: "OK"}, nil
}

func (s *server) GetAccount(_ context.Context, request *proto.GetAccountRequest) (*proto.GetAccountResponse, error) {

	if len(request.Name) == 0 {
		return nil, errors.New("empty name")
	}

	h.Guard.RLock()

	account, ok := h.Accounts[request.Name]

	h.Guard.RUnlock()

	if !ok {
		return nil, errors.New("account does not exist")
	}

	response := proto.GetAccountResponse{
		Name:   account.Name,
		Amount: int32(account.Amount),
	}

	return &response, nil
}

func main() {
	h = *accounts.New()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 4567))
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

//func main() {
//	accountsHandler := accounts.New()
//
//	// Echo instance
//	e := echo.New()
//
//	// Middleware
//	e.Use(middleware.Logger())
//	e.Use(middleware.Recover())
//
//	e.GET("/account", accountsHandler.GetAccount)
//	e.POST("/account/create", accountsHandler.CreateAccount)
//	e.POST("/account/patch", accountsHandler.PatchAccount)
//	e.POST("/account/change", accountsHandler.ChangeAccount)
//	e.POST("/account/delete", accountsHandler.DeleteAccount)
//
//	// Start server
//	e.Logger.Fatal(e.Start(":1323"))
//}
