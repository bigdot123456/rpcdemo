package main

import (
	"fmt"

	pb "rpcdemo/api"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	Address = "127.0.0.1:50052"
)

func main() {

	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	c := pb.NewDemoClient(conn)

	req := new(pb.HelloReq)
	req.Name = "kratos grpc111"
	r, err := c.SayHelloURL(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r.Content)

}

