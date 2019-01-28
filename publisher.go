package main

import (
	"log"
	"time"

	machinery "github.com/RichardKnop/machinery/v1"
	mconfig "github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
)

var (
	publisher *machinery.Server
)

func loadPublisher() {
	var err error
	publisher, err = machinery.NewServer(&mconfig.Config{
		Broker:          "amqp://guest:guest@localhost:5672/",
		DefaultQueue:    "tasking_",
		ResultBackend:   redis(),
		ResultsExpireIn: int(3600 / time.Second),
		AMQP: &mconfig.AMQPConfig{
			Exchange:      "tasking_exchange",
			ExchangeType:  "direct",
			BindingKey:    "tasking_bind",
			PrefetchCount: 1,
		},
	},
	)

	if err != nil {
		log.Fatalf("Could not create publisher server: %s\n", err.Error())
		return
	}
}

func publish(taskName string, arg ...float64) {

	_, err := publisher.SendTask(&tasks.Signature{
		Name:         taskName,
		RetryCount:   1,
		RetryTimeout: 10,
		Args: []tasks.Arg{
			{
				Type:  "float64",
				Value: arg[0],
			},
			{
				Type:  "float64",
				Value: arg[1],
			},
		},
	})

	if err != nil {
		log.Fatalf("Could not push task to queue: %s\n", err.Error())
		return
	}

	log.Println("Successfully Pushed Task To The Queue")
}
