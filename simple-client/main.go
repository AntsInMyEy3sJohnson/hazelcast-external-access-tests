package main

import (
	"context"
	"fmt"
	"github.com/hazelcast/hazelcast-go-client"
	"time"
)

func main() {
	fmt.Println("hello, hazelcast world!")

	config := hazelcast.Config{}
	cc := &config.Cluster

	cc.Network.SetAddresses("192.168.44.128")
	cc.Discovery.UsePublicIP = true
	ctx := context.TODO()

	client, err := hazelcast.StartNewClientWithConfig(ctx, config)

	if err != nil {
		panic(err)
	}

	m, err := client.GetMap(ctx, "awesome-map")
	if err != nil {
		panic(err)
	}

	if size, err := m.Size(ctx); err != nil {
		if size > 0 {
			fmt.Println("clearing map")
			_ = m.Clear(ctx)
		}
	}

	for i := 0; ; i++ {
		key := fmt.Sprintf("mykey-%d", i)
		value := fmt.Sprintf("myvalue-%d", i)
		fmt.Printf("writing key-value pair: %s / %s\n", key, value)

		err = m.Set(ctx, key, value)
		if err != nil {
			fmt.Printf("set unsuccessful with error: %v\n", err)
		}

		time.Sleep(500 * time.Millisecond)
	}

}
