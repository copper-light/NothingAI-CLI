package common

const (
	RESOURCE_URL        = "%v/api/v1/%v"
	RESOURCE_DETAIL_URL = RESOURCE_URL + "/%v"
	RESOURCE_FILES_URL  = RESOURCE_DETAIL_URL + "/files"
	EXEC_PERIEMENT_URL  = "%v/api/v1/experiments/%v/exec"
	LOG_URL             = "%v/api/v1/tasks/%v/logs"
)
