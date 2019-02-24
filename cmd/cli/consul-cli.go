package main

import (
	"amadeus-go/cmd/cli/utils"

	"fmt"

	consul "github.com/hashicorp/consul/api"
)

type cli struct {
	Name        string
	ConsulAgent *consul.Agent
}

func main() {
	c, err := consul.NewClient(consul.DefaultConfig())
	if err != nil {
		panic(err)
	}

	client := cli{
		Name:        "amadeus-go",
		ConsulAgent: c.Agent(),
	}

	_, resp, err := client.ConsulAgent.AgentHealthServiceByName("amadeus-go")
	if err != nil {
		panic(err)
	}

	grpcAddr := fmt.Sprintf("%s:%v", resp[0].Service.Address, resp[0].Service.Port)
	utils.SendReq(&grpcAddr, FLIGHT_LOW_FARE_SEARCH)
}
