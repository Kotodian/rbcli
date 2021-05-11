package main

import (
	"fmt"
	"github.com/Kotodian/rbcli/client"
	"github.com/Kotodian/rbcli/examples/util"
	"github.com/Kotodian/rbcli/naming"
	"log"
	"time"
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
	err = rbClient.Publish(routingKey, []byte("Hello World"))
	failOnError(err)
	fmt.Println("publish routingKey:", routingKey.String(false))
	time.Sleep(1 * time.Second)
}
