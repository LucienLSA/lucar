package main

import (
	"log"

	rentalpb "lucar/rental/api/gen/v1"
	"lucar/rental/trip"
	"lucar/shared/auth"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger, err := newZapLogger()
	if err != nil {
		log.Fatalf("can not create logger: %v\n", zap.Error(err))
	}
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		logger.Fatal("can not listen: %v\n", zap.Error(err))
	}
	in, err := auth.Interceptor("shared/auth/public.key")
	if err != nil {
		logger.Fatal("can not create auth interceptor", zap.Error(err))
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(in))
	rentalpb.RegisterTripServiceServer(s, &trip.Service{
		Logger: logger,
	})
	err = s.Serve(lis)
	if err != nil {
		logger.Fatal("can not server", zap.Error(err))
	}
}
func newZapLogger() (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.TimeKey = ""
	return cfg.Build()
}
