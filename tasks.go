package ntnxAPI

import (

	log "github.com/Sirupsen/logrus"

	"encoding/json"
	"strconv"
	"fmt"
)

const (

pullTimeout = "5"
maxTimeout = "1800"
succeeded = "Succeeded"

)


type TaskUUID struct {
	TaskUUID string `json:"taskUuid"`
}

type TaskPoll_Timeout_json_REST struct {
	TimedOut bool `json:"timedOut"`
	IsUnrecognized bool `json:"isUnrecognized"`
}

type Task_json_REST struct {
	UUID string `json:"uuid"`
	MetaRequest struct {
		MethodName string `json:"methodName"`
	} `json:"metaRequest"`
	MetaResponse struct {
		Error string `json:"error"`
		ErrorDetail string `json:"errorDetail"`
	} `json:"metaResponse"`
	CreateTime int64 `json:"createTime"`
	StartTime int64 `json:"startTime"`
	CompleteTime int64 `json:"completeTime"`
	LastUpdatedTime int64 `json:"lastUpdatedTime"`
	EntityList []struct {
		UUID string `json:"uuid"`
		EntityType string `json:"entityType"`
		EntityName string `json:"entityName"`
	} `json:"entityList"`
	OperationType string `json:"operationType"`
	Message string `json:"message"`
	PercentageComplete int `json:"percentageComplete"`
	ProgressStatus string `json:"progressStatus"`
}

type TaskPoll_json_REST struct {
	TaskInfo struct {
		UUID string `json:"uuid"`
		MetaRequest struct {
			MethodName string `json:"methodName"`
		} `json:"metaRequest"`
		MetaResponse struct {
			Error string `json:"error"`
			ErrorDetail string `json:"errorDetail"`
		} `json:"metaResponse"`
		CreateTime int64 `json:"createTime"`
		StartTime int64 `json:"startTime"`
		CompleteTime int64 `json:"completeTime"`
		LastUpdatedTime int64 `json:"lastUpdatedTime"`
		EntityList []struct {
			UUID string `json:"uuid"`
			EntityType string `json:"entityType"`
			EntityName string `json:"entityName"`
		} `json:"entityList"`
		OperationType string `json:"operationType"`
		Message string `json:"message"`
		PercentageComplete int `json:"percentageComplete"`
		ProgressStatus string `json:"progressStatus"`
	} `json:"taskInfo"`
	IsUnrecognized bool `json:"isUnrecognized"`
}

type TaskList_json_REST struct {
	Metadata struct {
		GrandTotalEntities int `json:"grandTotalEntities"`
		TotalEntities int `json:"totalEntities"`
	} `json:"metadata"`
	Entities []struct {
		UUID string `json:"uuid"`
		MetaRequest struct {
			MethodName string `json:"methodName"`
		} `json:"metaRequest"`
		CreateTime int64 `json:"createTime"`
		StartTime int64 `json:"startTime"`
		LastUpdatedTime int64 `json:"lastUpdatedTime"`
		EntityList []struct {
			UUID string `json:"uuid"`
			EntityType string `json:"entityType"`
			EntityName string `json:"entityName"`
		} `json:"entityList"`
		OperationType string `json:"operationType"`
		Message string `json:"message"`
		PercentageComplete int `json:"percentageComplete"`
		ProgressStatus string `json:"progressStatus"`
	} `json:"entities"`
}

func GetTaskbyTaskUUID(n *NTNXConnection, taskUUID string) Task_json_REST {
	
	resp, _ := NutanixAPIGet(n, NutanixAHVurl(n), "tasks/"+taskUUID)
	
	var t Task_json_REST

	json.Unmarshal(resp, &t)
	
	return t
	
}

func PollTaskbyTaskUUID(n *NTNXConnection, taskUUID string) (Task_json_REST,bool) {
	
	var t TaskPoll_json_REST
	var tp TaskPoll_Timeout_json_REST
	
	resp, statusCode := NutanixAPIGet(n, NutanixAHVurl(n), "tasks/"+taskUUID+"/poll?timeoutseconds="+pullTimeout)
		
	if ( statusCode == 200 ) {
		
		json.Unmarshal(resp, &t)
		
		if ( t.TaskInfo.UUID == "" ) {
		  json.Unmarshal(resp, &tp)
		  if ( tp.TimedOut ) {
			  return t.TaskInfo, false
			}
		}
			
       return t.TaskInfo, true
    } 
	
	return t.TaskInfo, false
	
}

func GetTaskPercentageCompletebyTaskUUID(n *NTNXConnection, taskUUID string) (string,error) {
	
	var t Task_json_REST
		
	resp, statusCode := NutanixAPIGet(n, NutanixAHVurl(n), "tasks/"+taskUUID)
		
	if ( statusCode == 200 ) {
		
		json.Unmarshal(resp, &t)
		
		 return string(t.PercentageComplete), nil		
		}			    
	
	log.Warn("Task "+taskUUID+" not found")
	return string(t.PercentageComplete), fmt.Errorf("Task "+taskUUID+" not found")	
	
}



func WaitUntilTaskFinished(n *NTNXConnection, taskUUID string) (Task_json_REST,error) {
	
	pullTimeoutInt, _ := strconv.Atoi(pullTimeout)
	maxTimeoutInt, _ := strconv.Atoi(maxTimeout)
	var i int 
	var finished bool
	
	var task Task_json_REST
	
	i = pullTimeoutInt
	
	for  i < maxTimeoutInt {
		
	 task, finished = PollTaskbyTaskUUID(n,taskUUID)
	 
	 log.Debug(task)	
	 	 
	 i = i + pullTimeoutInt
	 
	 if ( finished ) { 
		 return task, nil
	 } 
	 
	 percentageComplete, err := GetTaskPercentageCompletebyTaskUUID(n,taskUUID)
	 if ( err != nil) { 
	  log.Info("Task has "+percentageComplete+"% completed")
	 }
	}
	  
	log.Warn("Task "+taskUUID+" not found or timedout")
	return task, fmt.Errorf("Task "+taskUUID+" not found or timedout")	
}





