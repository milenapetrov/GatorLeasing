package service

import "GatorLeasing/gator-leasing-server/repository"

type LeaseService struct {
	repository *repository.LeaseRepository
}

func NewLeaseService(repository *repository.LeaseRepository) *LeaseService {
	return &LeaseService{repository: repository}
}
