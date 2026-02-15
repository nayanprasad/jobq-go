package service

import (
	"context"

	"github.com/nayanprasad/jobQ-go/internal/domain/job/command"
)

type JobService interface {
	CreateJob(ctx context.Context, args command.CreateJobCommand) error
}

type svc struct {
	//repo
	//db
}

func NewService() *svc {
	return &svc{}
}

func (s *svc) CreateJob(ctx context.Context, cmd command.CreateJobCommand) error {
	//query and create job here

	return nil
}
