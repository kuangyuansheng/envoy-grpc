package main
import (
	"net"
	"log"
	"strconv"
	"syscall"
	"fmt"
	"context"
	"os/signal"
	"os"
	"path/filepath"
	"google.golang.org/grpc"
	"envoy-grpc/app/handler"
	"envoy-grpc/app/health"

	cli "gopkg.in/urfave/cli.v1"
	pb "envoy-grpc/protos"
	pbh "google.golang.org/grpc/health/grpc_health_v1"
)

const (
	version = "3.0.1"
	usage   = "envoy proxy test"
)
var app *cli.App

func init(){
	app 	    = cli.NewApp()
	app.Name 	= filepath.Base(os.Args[0])
	app.Version = version
	app.Usage 	= usage

	app.Flags = []cli.Flag{
		cli.UintFlag{Name: "port, p", 	Usage: "端口"},
	}

	app.Action 	= func(ctx *cli.Context) error {
		p := ctx.GlobalUint("port")
		if p == 0 {
			log.Fatalf("Missing port!")
		}

		grpcServer := grpc.NewServer(			
			grpc.StreamInterceptor(StreamServerInterceptor),
			grpc.UnaryInterceptor(UnaryServerInterceptor),
		)
		lis, err := net.Listen("tcp", ":"+strconv.Itoa(int(p)))
		if err != nil {
			log.Fatalf("Failed to listen:%+v",err)
			return err
		}
	
		pb.RegisterGreeterServer(grpcServer, handler.New())
		pbh.RegisterHealthServer(grpcServer, health.New())
	
		go func() {
			sigs := make(chan os.Signal, 1)
			signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
			_ = <-sigs
			grpcServer.GracefulStop()
		}()	

		log.Printf("service started")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %+v", err)
			return err
		}
		return nil
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}


func UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// log.Printf("before handling. Info: %+v", info)
	resp, err := handler(ctx, req)
	// log.Printf("after handling. resp: %+v", resp)
	return resp, err
}

func StreamServerInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	// log.Printf("before handling. Info: %+v", info)
	err := handler(srv, ss)
	// log.Printf("after handling. err: %v", err)
	return err
}