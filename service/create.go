package service

import (
	model_kit "grpc-kit-service/model/kit"

	"github.com/espitman/protos-kit/kit"
	"github.com/go-playground/validator/v10"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Create(ctx context.Context, req *kit.RequestCreate) (*kit.ResponseDetails, error) {

	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	res, err := model_kit.CreateNewkit(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return res, nil
}
