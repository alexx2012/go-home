package main

import (
	"github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/cluster/consul"
	"github.com/AsynkronIT/protoactor-go/cluster"
	"github.com/alexx2012/go-home/shared"
	"log"
	"time"
	"fmt"
)

func main() {
	//remote.Register("Hello", actor.FromProducer(func() actor.Actor {
	//	return &shared.HelloActor{}
	//}))

	cp, err := consul.New()

	if err != nil {
		log.Fatal(err)
	}

	cluster.Start("go-home-cluster", "127.0.0.1:8081", cp)

	sync()
	//async()

	console.ReadLine()
}

func sync() {
	hello := shared.GetHelloGrain("abc2")
	options := []cluster.GrainCallOption{cluster.WithTimeout(5 * time.Second), cluster.WithRetry(5)}
	res, err := hello.SayHello(&shared.HelloRequest{Name: "GAM"}, options...)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Message from SayHello: %v", res.Message)

	for i := 0; i < 10000; i++ {
		x := shared.GetHelloGrain(fmt.Sprintf("hello%v", i))
		x.SayHello(&shared.HelloRequest{Name: "GAM2"})
	}

	log.Println("Done")
}