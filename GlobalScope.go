package AmazonSessionAPIClient

import "log"

type LOGS_LEVEL_TYPE int

const (
	DISABLE_LOGS LOGS_LEVEL_TYPE = -1
	ALL_LOGS     LOGS_LEVEL_TYPE = 0
	WARNING_LOGS LOGS_LEVEL_TYPE = 1
	ERROR_LOGS   LOGS_LEVEL_TYPE = 2
)

var LOGS_LEVEL LOGS_LEVEL_TYPE = ALL_LOGS

func PrintLog(logs_level LOGS_LEVEL_TYPE, message ...any) {
	if LOGS_LEVEL == DISABLE_LOGS {
		return
	}

	if LOGS_LEVEL == ALL_LOGS || logs_level == LOGS_LEVEL {
		log.Println(message...)
	}
}
