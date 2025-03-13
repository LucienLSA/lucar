package main

import (
	"log"

	rentalpb "lucar/rental/api/gen/v1"
	"lucar/rental/trip"
	"lucar/shared/server"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger, err := server.NewZapLogger()
	if err != nil {
		log.Fatalf("can not create logger: %v\n", zap.Error(err))
	}
	err = server.RunGRPCServer(&server.GRPCConfig{
		Name:              "rental",
		Addr:              ":8089",
		AuthPublicKeyFile: "shared/auth/public.key",
		Logger:            logger,
		RegisterFunc: func(s *grpc.Server) {
			rentalpb.RegisterTripServiceServer(s, &trip.Service{
				Logger: logger,
			})
		},
	})
	// if err != nil {
	// 	logger.Fatal("can not server", zap.Error(err))
	// }
	logger.Sugar().Fatal(err)
}
