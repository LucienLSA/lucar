package trip

import (
	"context"
	rentalpb "lucar/rental/api/gen/v1"
	"lucar/shared/auth"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Service implements a trip service.
type Service struct {
	Logger *zap.Logger
}

// CreateTrip creates a trip.
func (s *Service) CreateTrip(c context.Context, req *rentalpb.CreateTripRequest) (*rentalpb.CreateTripResponse, error) {
	// get accountID from context
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}
	s.Logger.Info("create trip", zap.String("start", req.Start), zap.String("account_id", aid))
	return nil, status.Error(codes.Unimplemented, "")
}
