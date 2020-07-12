package main

import (
	"fmt"
	"github.com/micro/go-micro/v2/config"
	"log"
)

type SjtYamlConfig struct {
	Config struct{
		Address string `json:"address"`
		Port int `json:"port"`
		Path string `json:"path"`
	} `json:"config"`
}
func main()  {
	// 读取json
	configFile := "application.yaml"
	err := config.LoadFile(configFile)
	if err !=nil {
		log.Fatal(err)
	}
	sjtConfig := &SjtYamlConfig{}
	err = config.Get("sjt").Scan(sjtConfig)
	if err !=nil {
		log.Fatal(err)
	}
	fmt.Println("res")
	fmt.Printf("%v\n", sjtConfig)
}
