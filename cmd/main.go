package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ProArash/go-movie-reseller/config"
	"github.com/ProArash/go-movie-reseller/internal/proto"
	"github.com/ProArash/go-movie-reseller/internal/route"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.LoadConfig()

	if cfg.Production {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.GrpcPort))

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	proto.RegisterGrpcServers(grpcServer)

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.SetTrustedProxies([]string{"127.0.0.1"})

	route.RegisterRoutes(router)

	addr := fmt.Sprintf(":%v", cfg.Port)

	go func() {
		log.Printf("gRPC server running on :%v", cfg.GrpcPort)
		if err := grpcServer.Serve(lis); err != nil {
			panic(err)
		}
	}()

	log.Printf("REST server running on :%v", cfg.Port)

	if err := router.Run(addr); err != nil {
		panic(err)
	}
}
