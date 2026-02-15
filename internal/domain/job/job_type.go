package job

type JobType string

const (
	JobTypeEmail   JobType = "email.send"
	JobTypeWebhook JobType = "webhook.deliver"
)

func (j JobType) IsValid() bool {
	switch j {
	case JobTypeEmail, JobTypeWebhook:
		return true
	default:
		return false
	}
}
