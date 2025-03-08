package main

import (
	"context"
	"log"
	authpb "lucar/auth/api/gen/v1"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseEnumNumbers: true,
				UseProtoNames:  true,
			},
		},
	))
	err := authpb.RegisterAuthServiceHandlerFromEndpoint(
		c, mux, "localhost:8088",
		[]grpc.DialOption{grpc.WithInsecure()},
	)
	if err != nil {
		log.Fatalf("can not register auth service:%v\n", err)
	}
	log.Fatal(http.ListenAndServe(":8080", mux))
}
