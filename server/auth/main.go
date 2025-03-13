package main

import (
	"io/ioutil"
	"log"
	authpb "lucar/auth/api/gen/v1"
	"lucar/auth/auth"
	"lucar/auth/dao"
	"lucar/auth/token"
	"lucar/auth/wechat"
	"net"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger, err := newZapLogger()
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
	pkFile, err := os.Open("auth/private.key")
	if err != nil {
		logger.Fatal("can not open private key", zap.Error(err))
	}
	pkBytes, err := ioutil.ReadAll(pkFile)
	if err != nil {
		logger.Fatal("can not read private key", zap.Error(err))
	}
	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(pkBytes)
	if err != nil {
		logger.Fatal("can not parse private key", zap.Error(err))
	}
	appSecretFile, err := os.Open("auth/appSecret.key")
	if err != nil {
		logger.Fatal("can not open appSecret", zap.Error(err))
	}
	aSBytes, err := ioutil.ReadAll(appSecretFile)
	if err != nil {
		logger.Fatal("can not read appSecret", zap.Error(err))
	}
	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &auth.Service{
		OpenIDResolver: &wechat.Service{
			AppID:     "wxe05234e87fd5ccb1",
			AppSecret: string(aSBytes),
		},
		Mongo:          dao.NewMongo(mdb.Database("lucar")),
		Logger:         logger,
		TokenExpire:    2 * time.Hour,
		TokenGenerator: token.NewJwtTokenGen("lucar/auth", privKey),
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
