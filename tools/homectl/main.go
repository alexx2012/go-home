package main

import (
	"flag"
	"log"
	"os"

	"strconv"

	"github.com/alexx2012/go-home/service/device/proto/device"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var address string

	flag.StringVar(&address, "a", "127.0.0.1:50071", "device service address")
	flag.Parse()

	args := os.Args[1:]

	if len(args) < 1 {
		log.Fatal("invalid arguments: provide action arg")
	}

	action := args[0]

	switch action {
	case "add":
		if len(args) < 4 {
			log.Fatal("invalid arguments: you should provide device id, subdevice id and type for new device")
		}

		deviceId, _ := strconv.Atoi(args[1])
		subDeviceId, _ := strconv.Atoi(args[2])
		deviceType, _ := strconv.Atoi(args[3])

		_, err := getClient(address).AddRawDevice(context.Background(), &device.RawDevice{
			ParentDeviceId: int32(deviceId),
			SubDeviceId:    int32(subDeviceId),
			SubDeviceType:  int32(deviceType),
		})

		if err != nil {
			log.Fatalf("can't create device, err: %v", err)
		}

		log.Println("device created")

		break

	case "del":
		if len(args) < 2 {
			log.Fatal("invalid arguments: you should provide raw device id")
		}

		rawDeviceId, _ := strconv.Atoi(args[1])

		_, err := getClient(address).RemoveRawDevice(context.Background(), &device.RawDeviceId{
			Value: int32(rawDeviceId),
		})

		if err != nil {
			log.Fatalf("can't delete device, err: %v", err)
		}

		log.Println("device deleted")

		break
	}

	//log.Printf(action)
}

func getClient(address string) device.DeviceServiceClient {
	var opts = []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(address, opts...)

	if err != nil {
		log.Fatalf("can't connect to server, err: %v", err)
	}

	return device.NewDeviceServiceClient(conn)
}
