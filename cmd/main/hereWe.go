package main

import (
	"context"
	"fmt"
	"net"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {

	p, err := net.ListenPacket("udp", ":1053")
	if err != nil {
		panic(err)
	}
	defer p.Close()


	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Println(container.ID)
		fmt.Println(container.Names)
		fmt.Println(container.NetworkSettings.Networks["bridge"].IPAddress)
	}

	for {
		buf := make([]byte, 512)
		n, addr, err := p.ReadFrom(buf)
		if err != nil {
			fmt.Printf("Connection error [%s]: %s\n", addr.String(), err)
			continue
		}

		fmt.Println(n)
		fmt.Println(addr)
	}


}