package main

import (
	"log"
	"net"
	"time"
	"google.golang.org/grpc"
	pb "lock/lock"
	"google.golang.org/grpc/reflection"

	_ "myapp/routers"
	"github.com/astaxie/beego"

)

const(
	port = "localhost:8082"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}


func (s *server) Lock(in *pb.Request,stream pb.Hello_LockServer) error {
	
	x := time.Now().Format("2006-01-02 15:04:05")
	
	for{
		timer := time.NewTimer(time.Minute*20)
			x = time.Now().Format("2006-01-02 15:04:05")
			stream.Send(&pb.Reply{Message:x})
		<-timer.C
	
	}

	return nil

}

func showTime(){
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("开启服务失败: %v", err)
	}
}


func main() {

	go showTime()

	beego.SetStaticPath("/base","/view/base")
	beego.SetStaticPath("/img","/static/image")
	beego.SetStaticPath("/css","/static/css")

	beego.Run()
}