package main

import (

	"log"
	"net"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "echo/echo"
	"google.golang.org/grpc/reflection"

	_ "myapp/routers"
	"github.com/astaxie/beego"

)

const(
	port = "localhost:8082"
)

type server struct{}


func (s *server) Echo(ctx context.Context, in *pb.Request) (*pb.Reply, error) {
	return &pb.Reply{Message:in.Message}, nil
}

func bee_echo(){
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("开启服务失败: %v", err)
	}
	
}

func main() {

	go bee_echo()

	beego.SetStaticPath("/base","/view/base")
	beego.SetStaticPath("/img","/static/image")
	beego.SetStaticPath("/css","/static/css")

	beego.Run()
}