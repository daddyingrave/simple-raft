package main

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"simpleraft/pkg/raft"
)

var Nodes []string
var Port string

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	pflag.StringSliceVar(&Nodes, "nodes", []string{}, "")
	pflag.StringVar(&Port, "port", "", "")
	pflag.Parse()

	app := fiber.New()

	entryChannel := make(chan int64)

	fsm := raft.NewFsm(Port, Nodes, entryChannel)
	service := raft.NewService(entryChannel)
	raft.AppendEntryRoutes(app, service)

	go func() {
		fsm.Run()
	}()

	err := app.Listen(":" + Port)
	if err != nil {
		log.Panicf("Error: %s", err)
	}
}
