package v1

import (
	"context"

	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.module_name }}-proto/pb"

	"github.com/sanservices/apilogger/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"{{ cookiecutter.module_name }}/internal/api/v1/dto"
)

func (h *Handler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	validation := &dto.GetUserRequestValidation{
		UserID: req.UserID,
	}

	err := h.validate.Struct(validation)
	if err != nil {
		apilogger.Error(ctx, apilogger.LogCatInputValidation, err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	user, err := h.service.GetUser(ctx, validation.UserID)
	if err != nil {
		apilogger.Error(ctx, apilogger.LogCatInputValidation, err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	return user, nil

}
