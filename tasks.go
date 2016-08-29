package ntnxAPI

import (
	log "github.com/Sirupsen/logrus"

	"encoding/json"
	"fmt"
	"strconv"
)

const (
	pullTimeout = "5"
	maxTimeout  = "1800"
	succeeded   = "Succeeded"
)

// TaskUUID ...
type TaskUUID struct {
	TaskUUID string `json:"taskUuid"`
}

// TaskPollTimeoutJSONREST ...
type TaskPollTimeoutJSONREST struct {
	TimedOut       bool `json:"timedOut"`
	IsUnrecognized bool `json:"isUnrecognized"`
}

// TaskJSONREST ...
type TaskJSONREST struct {
	UUID        string `json:"uuid"`
	MetaRequest struct {
		MethodName string `json:"methodName"`
	} `json:"metaRequest"`
	MetaResponse struct {
		Error       string `json:"error"`
		ErrorDetail string `json:"errorDetail"`
	} `json:"metaResponse"`
	CreateTime      int64 `json:"createTime"`
	StartTime       int64 `json:"startTime"`
	CompleteTime    int64 `json:"completeTime"`
	LastUpdatedTime int64 `json:"lastUpdatedTime"`
	EntityList      []struct {
		UUID       string `json:"uuid"`
		EntityType string `json:"entityType"`
		EntityName string `json:"entityName"`
	} `json:"entityList"`
	OperationType      string `json:"operationType"`
	Message            string `json:"message"`
	PercentageComplete int    `json:"percentageComplete"`
	ProgressStatus     string `json:"progressStatus"`
}

// TaskPollJSONREST ...
type TaskPollJSONREST struct {
	TaskInfo struct {
		UUID        string `json:"uuid"`
		MetaRequest struct {
			MethodName string `json:"methodName"`
		} `json:"metaRequest"`
		MetaResponse struct {
			Error       string `json:"error"`
			ErrorDetail string `json:"errorDetail"`
		} `json:"metaResponse"`
		CreateTime      int64 `json:"createTime"`
		StartTime       int64 `json:"startTime"`
		CompleteTime    int64 `json:"completeTime"`
		LastUpdatedTime int64 `json:"lastUpdatedTime"`
		EntityList      []struct {
			UUID       string `json:"uuid"`
			EntityType string `json:"entityType"`
			EntityName string `json:"entityName"`
		} `json:"entityList"`
		OperationType      string `json:"operationType"`
		Message            string `json:"message"`
		PercentageComplete int    `json:"percentageComplete"`
		ProgressStatus     string `json:"progressStatus"`
	} `json:"taskInfo"`
	IsUnrecognized bool `json:"isUnrecognized"`
}

// TaskListJSONREST ...
type TaskListJSONREST struct {
	Metadata struct {
		GrandTotalEntities int `json:"grandTotalEntities"`
		TotalEntities      int `json:"totalEntities"`
	} `json:"metadata"`
	Entities []struct {
		UUID        string `json:"uuid"`
		MetaRequest struct {
			MethodName string `json:"methodName"`
		} `json:"metaRequest"`
		CreateTime      int64 `json:"createTime"`
		StartTime       int64 `json:"startTime"`
		LastUpdatedTime int64 `json:"lastUpdatedTime"`
		EntityList      []struct {
			UUID       string `json:"uuid"`
			EntityType string `json:"entityType"`
			EntityName string `json:"entityName"`
		} `json:"entityList"`
		OperationType      string `json:"operationType"`
		Message            string `json:"message"`
		PercentageComplete int    `json:"percentageComplete"`
		ProgressStatus     string `json:"progressStatus"`
	} `json:"entities"`
}

// GetTaskbyTaskUUID ...
func GetTaskbyTaskUUID(n *NTNXConnection, taskUUID string) TaskJSONREST {

	resp, _ := NutanixAPIGet(n, NutanixAHVurl(n), "tasks/"+taskUUID)

	var t TaskJSONREST

	json.Unmarshal(resp, &t)

	return t

}

// PollTaskbyTaskUUID ...
func PollTaskbyTaskUUID(n *NTNXConnection, taskUUID string) (TaskJSONREST, bool) {

	var t TaskPollJSONREST
	var tp TaskPollTimeoutJSONREST

	resp, statusCode := NutanixAPIGet(n, NutanixAHVurl(n), "tasks/"+taskUUID+"/poll?timeoutseconds="+pullTimeout)

	if statusCode == 200 {

		json.Unmarshal(resp, &t)

		if t.TaskInfo.UUID == "" {
			json.Unmarshal(resp, &tp)
			if tp.TimedOut {
				return t.TaskInfo, false
			}
		}

		return t.TaskInfo, true
	}

	return t.TaskInfo, false

}

// GetTaskPercentageCompletebyTaskUUID ...
func GetTaskPercentageCompletebyTaskUUID(n *NTNXConnection, taskUUID string) (string, error) {

	var t TaskJSONREST

	resp, statusCode := NutanixAPIGet(n, NutanixAHVurl(n), "tasks/"+taskUUID)

	if statusCode == 200 {

		json.Unmarshal(resp, &t)

		return string(t.PercentageComplete), nil
	}

	log.Warn("Task " + taskUUID + " not found")
	return string(t.PercentageComplete), fmt.Errorf("Task " + taskUUID + " not found")

}

// WaitUntilTaskFinished ...
func WaitUntilTaskFinished(n *NTNXConnection, taskUUID string) (TaskJSONREST, error) {

	pullTimeoutInt, _ := strconv.Atoi(pullTimeout)
	maxTimeoutInt, _ := strconv.Atoi(maxTimeout)
	var i int
	var finished bool

	var task TaskJSONREST

	i = pullTimeoutInt

	for i < maxTimeoutInt {

		task, finished = PollTaskbyTaskUUID(n, taskUUID)

		log.Debug(task)

		i = i + pullTimeoutInt

		if finished {
			return task, nil
		}

		percentageComplete, err := GetTaskPercentageCompletebyTaskUUID(n, taskUUID)
		if err != nil {
			log.Info("Task has " + percentageComplete + "% completed")
		}
	}

	log.Warn("Task " + taskUUID + " not found or timedout")
	return task, fmt.Errorf("Task " + taskUUID + " not found or timedout")
}

// WrappWaitUntilTaskFinished ...
func WrappWaitUntilTaskFinished(n *NTNXConnection, taskUUID string, successStr string) {

	_, err := WaitUntilTaskFinished(n, taskUUID)

	if err != nil {
		log.Fatal("Task does not exist")
	} else {
		log.Info(successStr)
	}

}
