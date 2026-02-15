package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/nayanprasad/jobq-go/internal/domain/job"
)

type JobRepository struct {
	db *pgx.Conn
}

func NewJobRepository(db *pgx.Conn) *JobRepository {
	return &JobRepository{
		db: db,
	}
}

func (r *JobRepository) Create(ctx context.Context, j *job.Job) (*job.Job, error) {

	//query here

	return &job.Job{}, nil
}
