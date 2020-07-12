package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/shijting/go-micro/src/courses"
	"github.com/shijting/go-micro/src/service"
	"log"
)



func main() {
	// create a new service
	serv := micro.NewService(
		micro.Name("go.micro.api.sjt.courses"),
	)

	// initialise flags
	serv.Init()
	err := courses.RegisterCoursesServiceHandler(serv.Server(), service.NewCourse())
	if err != nil {
		log.Fatal(err)
	}
	// start the service
	serv.Run()
}
