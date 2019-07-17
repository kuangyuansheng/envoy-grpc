package main

import(
	"time"
	"context"
	"os"
	"fmt"
	"log"
	"google.golang.org/grpc"
	pb "envoy-grpc/protos"
	"path/filepath"
	cli "gopkg.in/urfave/cli.v1"

)


const (
	version = "3.0.1"
	usage   = "envoy-grpc client test"
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
		cli.StringFlag{Name: "address, a", 	Usage: "请求地址"},
	}

	// Run执行动作
	app.Action 	= func(ctx *cli.Context) error {

		a := ctx.GlobalString("address")
		if a == "" {
			log.Fatalf("Missing address!")
		}

		conn, err := grpc.Dial(a, grpc.WithInsecure())
		// conn, err := grpc.Dial("prelease-consul.pointsmart.cn:42323", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
	
		f := pb.NewGreeterClient(conn)
		c, cancel := context.WithTimeout(context.Background(), time.Second * 30)
		defer cancel()
	
	
		r, err := f.SayHello(c, &pb.HelloRequest{
			Name:"test_name",
		})
	
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
	
		log.Println(r)
		return nil
	}

}


func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}