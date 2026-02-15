package job

type JobStatus string

const (
	StatusPending   JobStatus = "pendig"
	StatusRunning   JobStatus = "running"
	StatusCompleted JobStatus = "completed"
	StatusFailed    JobStatus = "failed"
	StatudDead      JobStatus = "dead"
)
