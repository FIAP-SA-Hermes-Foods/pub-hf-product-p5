package main

import (
	"context"
	"log"
	"net"
	"os"
	"pub-hf-product-p5/external/broker"
	l "pub-hf-product-p5/external/logger"
	productBroker "pub-hf-product-p5/internal/adapters/broker"
	"pub-hf-product-p5/internal/core/application"
	grpcH "pub-hf-product-p5/internal/handler/rpc"
	cp "pub-hf-product-p5/product_pub_proto"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/marcos-dev88/genv"
	"google.golang.org/grpc"
)

func init() {
	if err := genv.New(".env.local"); err != nil {
		l.Errorf("", "error set envs %v", " | ", err)
	}
}

func main() {

	listener, err := net.Listen("tcp", ":"+os.Getenv("PUB_HF_PRODUCT_RPC_PORT"))

	if err != nil {
		l.Errorf("", "error to create connection %v", " | ", err)
	}

	defer listener.Close()

	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("", "unable to load SDK config, %v", err)
	}

	b := broker.NewSQSBroker(cfg)

	pb := productBroker.NewProductBroker(b, os.Getenv("PUB_HF_PRODUCT_QUEUE"))

	app := application.NewApplication(ctx, pb)

	h := grpcH.NewHandler(app)

	grpcServer := grpc.NewServer()

	cp.RegisterProductServer(grpcServer, h.Handler())

	if err := grpcServer.Serve(listener); err != nil {
		l.Errorf("", "error in server: %v", " | ", err)
	}
}
