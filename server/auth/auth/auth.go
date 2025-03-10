package auth

import (
	"context"
	authpb "lucar/auth/api/gen/v1"
	"lucar/auth/dao"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Service implements auth service
type Service struct {
	OpenIDResolver OpenIDResolver
	Mongo          *dao.Mongo
	Logger         *zap.Logger
}

// OpenIDResolver resolves an authorization code to an open_id
type OpenIDResolver interface {
	Resolve(code string) (string, error)
}

// Login logs a user in
func (s *Service) Login(c context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	// s.Logger.Info("received code", zap.String("code", req.Code))
	openID, err := s.OpenIDResolver.Resolve(req.Code)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "can not resolve openid:%v\n", err)
	}
	accountID, err := s.Mongo.ResolveAccountID(c, openID)
	if err != nil {
		s.Logger.Error("can not resolve account id", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}
	return &authpb.LoginResponse{
		// AccessToken: "access_token for: " + req.Code,
		AccessToken: "token for account id" + accountID,
		ExpiresIn:   7200,
	}, nil
}
