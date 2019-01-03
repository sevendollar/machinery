ckage main

import (
        "fmt"
        "time"

        machinery "github.com/RichardKnop/machinery/v1"
        "github.com/RichardKnop/machinery/v1/config"
        "github.com/RichardKnop/machinery/v1/tasks"
        "github.com/jasonlvhit/gocron"
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
                fmt.Println(err, "Can not create server!")
        }

        getShitTask := tasks.Signature{
                Name: "getShit",
                Args: []tasks.Arg{
                        tasks.Arg{
                                Type:  "string",
                                Value: "jef rocks!",
                        },
                },
        }

        server.SendTask(&getShitTask)
}

