package main

import (
	"fmt"
	"github.com/micro/go-micro/v2/config"
	"log"
)

type SjtConfig struct {
	Db struct{
		Ip string `json:"ip"`
		Port int `json:"port"`
	} `json:"db"`
	Redis struct{
		Ip string `json:"ip"`
		Port int `json:"port"`
	} `json:"redis"`
}
func main()  {
	// 读取json
	configFile := "application.json"
	err := config.LoadFile(configFile)
	if err !=nil {
		log.Fatal(err)
	}
	sjtConfig := &SjtConfig{}
	err = config.Get("sjt").Scan(sjtConfig)
	if err !=nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", sjtConfig)
}
