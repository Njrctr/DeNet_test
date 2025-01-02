package service

import "github.com/Njrctr/DeNet_test/pkg/repository"

type Service struct{}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
