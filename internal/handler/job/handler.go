package job

import (
	"context"

	"github.com/nayanprasad/jobQ-go/internal/domain/job"
)

type Handler interface {
	Type() job.JobType
	Execute(ctx context.Context, payload []byte) error
}
