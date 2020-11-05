package main

import (
	"fmt"

	"github.com/hashicorp/consul/api"

	_ "github.com/hashicorp/consul/connect"
)

func main() {
	// Get a new client
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	// Get a handle to the KV API
	kv := client.KV()

	// PUT a new KV pair
	p := &api.KVPair{Key: "REDIS_MAXCLIENTS", Value: []byte("1000")}
	_, err = kv.Put(p, nil)
	if err != nil {
		panic(err)
	}

	// Lookup the pair
	pair, _, err := kv.Get("REDIS_MAXCLIENTs", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("res:%#v\n", pair)
	fmt.Printf("KV: %v %s\n", pair.Key, string(pair.Value))
}
