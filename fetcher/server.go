package fetcher

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
)

var cnf = &config.Config{
	Broker:        "amqp://guest:guest@rabbit:5672/",
	DefaultQueue:  "default",
	ResultBackend: "amqp://guest:guest@rabbit:5672/",
	AMQP: &config.AMQPConfig{
		Exchange:     "fetcher",
		ExchangeType: "direct",
		BindingKey:   "fetcher",
	},
}

func StartServer() (*machinery.Server, error) {
	var tasks = map[string]interface{}{
		"schedule_fetch": ScheduleFetch,
		"stop_fetch":     StopFetch,
	}
	server, err := machinery.NewServer(cnf)
	if err != nil {
		return server, err
	}
	return server, server.RegisterTasks(tasks)
}
