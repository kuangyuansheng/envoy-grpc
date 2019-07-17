package main
import (
	"net"
	"log"
	"strconv"
	"syscall"
	"fmt"
	"os/signal"
	"os"
	"path/filepath"
	"google.golang.org/grpc"
	pb "envoy-grpc/protos"
	"envoy-grpc/app/handler"
	"envoy-grpc/app/health"

	cli "gopkg.in/urfave/cli.v1"
	pbh "google.golang.org/grpc/health/grpc_health_v1"
)

const (
	version = "3.0.1"
	usage   = "envoy proxy test"
)
var (
	app *cli.App
)

func init(){
	app 	    = cli.NewApp()
	app.Name 	= filepath.Base(os.Args[0])
	app.Version = version
	app.Usage 	= usage

	// 定义命令行参数
	app.Flags = []cli.Flag{
		cli.UintFlag{Name: "port, p", 	Usage: "端口"},
	}

	// Run执行动作
	app.Action 	= func(ctx *cli.Context) error {
		p := ctx.GlobalUint("port")
		if p == 0 {
			log.Fatalf("Missing port!")
		}

		grpcServer := grpc.NewServer()
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
		// go func(){
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
