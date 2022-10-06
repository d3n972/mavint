package models

type TrainDetailsResponse struct {
	ExceptionMessage      string                  `json:"exceptionMessage,omitempty"`
	TrainSchedulerDetails []TrainSchedulerDetails `json:"trainSchedulerDetails"`
}
