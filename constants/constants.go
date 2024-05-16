package constants

const (
	VERSION  = "0.0.1"
	APP_NAME = "Nothing AI"

	MODELS_GET_URL      = "/api/v1/models"
	DATASETS_GET_URL    = "/api/v1/datasets"
	EXPERIMENTS_GET_URL = "/api/v1/experiments"
	TASKS_GET_URL       = "/api/v1/tasks"

	MODEL_DESC_URL       = MODELS_GET_URL + "/%s"
	DATASET_DESC_URL     = DATASETS_GET_URL + "/%s"
	EXPERIEMENT_DESC_URL = EXPERIMENTS_GET_URL + "/%s"
	TASK_DESC_URL        = TASKS_GET_URL + "/%s"
)
