package main

import (
        "fmt"

        machinery "github.com/RichardKnop/machinery/v1"
        "github.com/RichardKnop/machinery/v1/config"
)

func main() {

        var cnf = config.Config{
                Broker:        "amqp://jef:rabbitpw@10.5.1.114:5672",
                DefaultQueue:  "machinery_tasks",
                ResultBackend: "mongodb://jef:jefpw@10.5.1.114:27017",
                AMQP: &config.AMQPConfig{
                        Exchange:     "machinery_exchange",
                        ExchangeType: "direct",
                        BindingKey:   "machinery_task",
                },
        }

        server, err := machinery.NewServer(&cnf)
        if err != nil {
                fmt.Println(err, "Could not create server")
        }

        server.RegisterTask("Say", Say)

        err = server.NewWorker("worker-1", 10).Launch()
        if err != nil {
                fmt.Println(err, "Could not launch worker!")
        }
}
