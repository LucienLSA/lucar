package server

import (
	"lucar/shared/auth"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// GRPCConfig defines a grpc server
type GRPCConfig struct {
	Name              string
	Addr              string
	AuthPublicKeyFile string
	RegisterFunc      func(*grpc.Server)
	Logger            *zap.Logger
}

// RunGRPCServer runs a grpc server
func RunGRPCServer(c *GRPCConfig) error {
	nameField := zap.String("name", c.Name)
	lis, err := net.Listen("tcp", c.Addr)
	if err != nil {
		c.Logger.Fatal("can not listen", nameField, zap.Error(err))
	}
	var opts []grpc.ServerOption
	if c.AuthPublicKeyFile != "" {
		in, err := auth.Interceptor(c.AuthPublicKeyFile)
		if err != nil {
			c.Logger.Fatal("can not create auth interceptor", nameField, zap.Error(err))
		}
		opts = append(opts, grpc.UnaryInterceptor(in))
	}

	s := grpc.NewServer(opts...)
	c.RegisterFunc(s)
	// rentalpb.RegisterTripServiceServer(s, &trip.Service{
	// 	Logger: logger,
	// })
	c.Logger.Info("server started", nameField, zap.String("addr", c.Addr))
	return s.Serve(lis)
	// err = s.Serve(lis)
	// if err != nil {
	// 	logger.Fatal("can not server", zap.Error(err))
	// }
}
