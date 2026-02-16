package repository

import (
	"context"

	"github.com/nayanprasad/jobq-go/internal/domain/job"
)

// this interface defines the contract for transactions (independent of storage layer)
type JobRepository interface {
	Create(ctx context.Context, job *job.Job) (*job.Job, error)
	Get(ctx context.Context, id int32) (*job.Job, error)
}
