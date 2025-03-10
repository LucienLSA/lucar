package main

import (
	"log"
	authpb "lucar/auth/api/gen/v1"
	"lucar/auth/auth"
	"lucar/auth/dao"
	"lucar/auth/wechat"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can not create logger: %v\n", zap.Error(err))
	}
	lis, err := net.Listen("tcp", ":8088")
	if err != nil {
		logger.Fatal("can not listen: %v\n", zap.Error(err))
	}
	mdb, err := dao.InitMongo()
	if err != nil {
		logger.Fatal("dao.InitMongo failed: %v\n", zap.Error(err))
	}

	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &auth.Service{
		OpenIDResolver: &wechat.Service{
			AppID:     "wxe05234e87fd5ccb1",
			AppSecret: "0d207a45cc64dec4436e74739f8ef563",
		},
		Mongo:  dao.NewMongo(mdb.Database("lucar")),
		Logger: logger,
	})
	err = s.Serve(lis)
	if err != nil {
		logger.Fatal("can not server", zap.Error(err))
	}
}

// func newZapLogger() (*zap.Logger, error) {
// 	cfg := zap.NewDevelopmentConfig()
// 	cfg.EncoderConfig.TimeKey = ""
// 	return cfg.Build()
// }
