package main

import (
	"github.com/Rakanixu/flag/srv/elastic"
	"github.com/Rakanixu/flag/srv/handler"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"log"
	"strings"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.flag"),
		micro.Version("latest"),
		micro.Flags(
			cli.StringFlag{
				Name:   "elasticsearch_hosts",
				EnvVar: "ELASTICSEARCH_HOSTS",
				Usage:  "Comma separated list of elasticsearch hosts",
				Value:  "localhost:9200",
			},
		),
		micro.Action(func(c *cli.Context) {
			parts := strings.Split(c.String("elasticsearch_hosts"), ",")
			elastic.Hosts = parts
		}),
	)

	// Initialise service
	service.Init()

	// Initialise elasticsearch
	elastic.Init()

	// Register Handler
	service.Server().Handle(
		service.Server().NewHandler(new(handler.Flag)),
	)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
