package task

import (
	"github.com/hibiken/asynq"
	"github.com/joint-online-judge/go-horse/pkg/json"
)

// A list of task types.
const (
	TypeSubmitRecord = "submit:record"
)

type SubmitRecordPayload struct{}

func NewSubmitRecordTask() (*asynq.Task, error) {
	payload, err := json.Marshal(SubmitRecordPayload{})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeSubmitRecord, payload), nil
}
