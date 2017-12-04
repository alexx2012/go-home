package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"flag"

	"github.com/alexx2012/go-home/service/device/persistence/repository/mongo"
	"github.com/alexx2012/go-home/service/device/proto/device"
	"github.com/alexx2012/go-home/service/device/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/mgo.v2"
)

const (
	defaultMongoUrl = "127.0.0.1:27018"
	mongoDbName     = "hs"
)

var mongoUrl string

func init() {
	mongoUrl = GetVar("MONGO_URL", defaultMongoUrl)
}

func main() {
	var address, port = parseFlags()

	if err := start(address, port); err != nil {
		log.Fatal(err)
	}
}

func start(address string, port string) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", address, port))

	if err != nil {
		return err
	}

	mongoSession, err := mgo.Dial(mongoUrl)

	if err != nil {
		return fmt.Errorf("mongo err: %v", err)
	}

	mongoDb := mongoSession.DB(mongoDbName)
	repository := mongo.NewRawDeviceRepository(mongoDb)

	errChan := make(chan error, 1)

	go func() {
		server := service.NewDeviceService(repository)

		s := grpc.NewServer()
		device.RegisterDeviceServiceServer(s, server)

		reflection.Register(s)

		log.Printf("Started device service on [%s:%s]", address, port)

		errChan <- s.Serve(lis)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		log.Printf("Signal %v", <-c)

		errChan <- nil
	}()

	return <-errChan
}

func parseFlags() (string, string) {
	var a, p string

	flag.StringVar(&a, "address", "", "address")
	flag.StringVar(&p, "port", "50071", "port")
	flag.Parse()

	return a, p
}

func GetVar(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
