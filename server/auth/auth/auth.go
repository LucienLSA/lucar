package auth

import (
	"context"
	authpb "lucar/auth/api/gen/v1"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Service implements auth service
type Service struct {
	OpenIDResolver OpenIDResolver
	Logger         *zap.Logger
}

// OpenIDResolver resolves an authorization code to an open_id
type OpenIDResolver interface {
	Resolve(code string) (string, error)
}

// Login logs a user in
func (s *Service) Login(c context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	s.Logger.Info("received code", zap.String("code", req.Code))
	openID, err := s.OpenIDResolver.Resolve(req.Code)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "can not resolve openid:%v\n", err)
	}
	return &authpb.LoginResponse{
		// AccessToken: "access_token for: " + req.Code,
		AccessToken: "token for open id" + openID,
		ExpiresIn:   7200,
	}, nil
}
