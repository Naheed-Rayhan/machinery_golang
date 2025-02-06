package connection

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
)

func InitServer_v1() (*machinery.Server, error) {
	cfg := &config.Config{
		Broker:          "amqp://guest:guest@localhost:5672",
		DefaultQueue:    "develop_machinery_task",
		ResultBackend:   "amqp://guest:guest@localhost:5672",
		ResultsExpireIn: 100,
		AMQP: &config.AMQPConfig{
			Exchange:      "develop_machinery_exchange",
			ExchangeType:  "direct",
			BindingKey:    "develop_machinery_task",
			PrefetchCount: 20,
		},
	}
	return machinery.NewServer(cfg)
}
