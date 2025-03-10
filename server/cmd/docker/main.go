package main

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func main() {
	c, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	resp, err := c.ContainerCreate(ctx, &container.Config{
		Image: "mongo:4.4",
		ExposedPorts: nat.PortSet{
			"27017/tcp": {},
		},
	}, &container.HostConfig{
		PortBindings: nat.PortMap{
			"27017/tcp": []nat.PortBinding{
				{
					HostIP:   "127.0.0.1",
					HostPort: "0", // 自动选择一个端口
				},
			},
		},
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}
	// 启动容器
	err = c.ContainerStart(ctx, resp.ID, container.StartOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println("container started")
	time.Sleep(10 * time.Second)

	// 查看容器id 监听的端口
	inspRes, err := c.ContainerInspect(ctx, resp.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("listening at %+v\n",
		inspRes.NetworkSettings.Ports["27017/tcp"][0])

	// 容器终止
	fmt.Println("killing container")
	err = c.ContainerRemove(ctx, resp.ID, container.RemoveOptions{
		Force: true,
	})
	if err != nil {
		panic(err)
	}
}
