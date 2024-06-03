package main

import (
	"context"
	"testing"

	"github.com/lukanzx/DouServer/cmd/user/dal"
	"github.com/lukanzx/DouServer/cmd/user/rpc"
	"github.com/lukanzx/DouServer/cmd/user/service"
	"github.com/lukanzx/DouServer/config"
)

var (
	username string
	password string
	token    string
	id       int64

	userService *service.UserService
)

func TestMain(m *testing.M) {
	config.InitForTest()
	dal.Init()
	rpc.Init()

	userService = service.NewUserService(context.Background())

	username = "xiaoming"
	password = "123456"

	m.Run()
}

func TestMainOrder(t *testing.T) {
	t.Run("register", testRegister)

	t.Run("login", testLogin)

	// t.Run("info", testGetUserInfo)

	t.Run("RPC Test", testRPC)
}
