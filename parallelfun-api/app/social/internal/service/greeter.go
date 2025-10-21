package service

import (
	//"context"

	v1 "parallelfun-api/api/social/v1"
	"parallelfun-api/app/social/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedSocialServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
