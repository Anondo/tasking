package main

import (
	"log"
	"time"

	machinery "github.com/RichardKnop/machinery/v1"
	mconfig "github.com/RichardKnop/machinery/v1/config"
)

var (
	consumer *machinery.Server
)

func loadConsumer() {
	var err error
	consumer, err = machinery.NewServer(&mconfig.Config{
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
		log.Fatalf("Could not create consumer server: %s\n", err.Error())
		return
	}

	if err := consumer.RegisterTasks(tsks); err != nil {
		log.Fatalf("Failed to register task: %s\n", err.Error())
		return
	}
}

func consume() {

	if err := consumer.NewWorker("tasking_worker", 1).Launch(); err != nil { // 1 is the concurrency number i.e the number of workers
		log.Fatalf("Worker is dead: %s", err.Error())
		return
	}

}
