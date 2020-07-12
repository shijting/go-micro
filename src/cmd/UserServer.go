package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	"github.com/shijting/go-micro/src/Users"
	"log"
)

type UserServices struct {
}

func NewUser() *UserServices {
	return &UserServices{}
}

func (*UserServices) Test(ctx context.Context, req *Users.UserRequest, resp *Users.UserResponse) error {
	resp.Ret = "users" + req.Id
	return nil
}

func main() {
	// create a new service
	service := micro.NewService(
		micro.Name("go.micro.api.sjt.users"),
	)
	// initialise flags
	service.Init()
	err := Users.RegisterUserServiceHandler(service.Server(), NewUser())
	if err != nil {
		log.Fatal(err)
	}
	// start the service
	service.Run()
}
