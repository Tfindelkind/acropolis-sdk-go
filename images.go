package ntnxAPI

import (

	log "github.com/Sirupsen/logrus"

	"bytes"
	"encoding/json"
	"fmt"	
)

const (
NFSvmdiskPath = "/.acropolis/vmdisk/"
	)

type ImageList_AHV struct {
	Entities []struct {
		Annotation         string `json:"annotation"`
		ContainerID        int    `json:"containerId"`
		CreatedTimeInUsecs int    `json:"createdTimeInUsecs"`
		Deleted            bool   `json:"deleted"`
		ImageState         string `json:"imageState"`
		ImageType          string `json:"imageType"`
		LogicalTimestamp   int    `json:"logicalTimestamp"`
		Name               string `json:"name"`
		UpdatedTimeInUsecs int    `json:"updatedTimeInUsecs"`
		UUID               string `json:"uuid"`
		VMDiskID           string `json:"vmDiskId"`
	} `json:"entities"`
	Metadata struct {
		GrandTotalEntities int `json:"grandTotalEntities"`
		TotalEntities      int `json:"totalEntities"`
	} `json:"metadata"`
}

func GetImageIDbyName(n *NTNXConnection, ImageName string) (string,error) {

	resp, _ := NutanixAPIGet(n,NutanixAHVurl(n), "images/?filterCriteria=name%3D%3D"+ImageName)

	var iml ImageList_AHV

	json.Unmarshal(resp, &iml)

	// TO-DO: Handle if more than one images found
	
	s := iml.Entities

	for i := 0; i < len(s); i++ {
		if s[i].Name == ImageName {
			
			return s[i].VMDiskID, nil
		}

	}
	
	log.Warn("Image "+ImageName+" not found")
	return "", fmt.Errorf("Image "+ImageName+" not found")
}

// Checks if Image exists and fills helper struct with UUID and VMDiskID
func ImageExistbyName(n *NTNXConnection, im *Image) bool {

	// Image names are not unique so could return > 1 value
	resp, _ := NutanixAPIGet(n, NutanixAHVurl(n), "images")

	var iml ImageList_AHV

	json.Unmarshal(resp, &iml)

	// TO-DO: Handle if more than one images found

	s := iml.Entities

	for i := 0; i < len(s); i++ {
		if s[i].Name == im.Name {
			im.UUID = s[i].UUID
			im.VMDiskID = s[i].VMDiskID
			return true
		}
	}
	return false
}

func CloneCDforVM(n *NTNXConnection, v *VM, im *Image) {

	var jsonStr = []byte(`{ "disks": [ { "vmDiskClone":  { "vmDiskUuid": "` + im.VMDiskID + `" } , "isCdrom" : "true"} ] }`)

	NutanixAPIPost(n, NutanixAHVurl(n), "vms/"+v.VmId+"/disks/", bytes.NewBuffer(jsonStr))

    
}

func GenerateNFSURIfromVDisk(host string,container_name string, VMDiskID string) string {
 
  return "nfs://"+host+"/"+container_name+NFSvmdiskPath+VMDiskID

}

func CreateImageFromURL(n *NTNXConnection, d *VDisk, im *Image) (TaskUUID,error) {
	
	SourceContainerName, err := GetContainerNamebyUUID(n,d.ContainerUUID)
	if err != nil {
    log.Fatal(err)
	}	
		
	var jsonStr = []byte(`{ "name": "`+im.Name+`","annotation": "`+im.Annotation+`", "imageType":"DISK_IMAGE", "imageImportSpec": {"containerName": "`+im.ContainerName+`","url":"`+GenerateNFSURIfromVDisk(n.NutanixHost,SourceContainerName,d.VdiskUuid)+`"} }`)
	var task TaskUUID

	resp , statusCode := NutanixAPIPost(n, NutanixAHVurl(n), "images", bytes.NewBuffer(jsonStr))	
	
	if ( statusCode == 200) {
		json.Unmarshal(resp, &task)
		return task, nil
	 }	
  
  log.Warn("Image "+im.Name+" could not created on container "+im.ContainerName+" from "+GenerateNFSURIfromVDisk(n.NutanixHost,SourceContainerName,d.VdiskUuid))
  return task,fmt.Errorf("Image "+im.Name+" could not created on container "+im.ContainerName+" from "+GenerateNFSURIfromVDisk(n.NutanixHost,SourceContainerName,d.VdiskUuid))
}

func CreateImageObject(n *NTNXConnection, im *Image) (TaskUUID,error) {
				
	var jsonStr = []byte(`{ "name": "`+im.Name+`","annotation": "`+im.Annotation+`", "imageType":"`+im.ImageType+`" }`)
	var task TaskUUID
	
	resp, statusCode := NutanixAPIPost(n, NutanixAHVurl(n), "images", bytes.NewBuffer(jsonStr))
	
	if ( statusCode == 200 ) {
	   json.Unmarshal(resp, &task)	
       return task, nil
    } 
  
  log.Warn("Image "+im.Name+" could not created")
  return task, fmt.Errorf("Image "+im.Name+" could not created")

}

func GetImageIDbyTask(n *NTNXConnection, t *Task_json_REST) string {
	
	return t.EntityList[0].UUID
		
}




