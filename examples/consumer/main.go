package main

import (
	"context"
	"fmt"
	"github.com/Kotodian/rbcli/client"
	"github.com/Kotodian/rbcli/examples/util"
	"github.com/Kotodian/rbcli/naming"
	"log"
)

func failOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	rbClient := client.NewClient([]string{util.URL})
	err := rbClient.Conn(naming.NewExchange("ocpp", naming.Direct))
	failOnError(err)
	routingKey := naming.NewRoutingKey("ocpp", util.NewOcppTopic("ocpp01", "register"))
	err = rbClient.Subscribe(naming.NewQueue("ocpp", routingKey), func(ctx context.Context) error {
		fmt.Println("consume")
		delivery := client.DeliveryFromContext(ctx)
		fmt.Println(string(delivery.Body))
		return nil
	})
	fmt.Println("subscribe routingKey:", routingKey.String(false))
	failOnError(err)
	select {}
}
