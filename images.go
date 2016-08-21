package ntnxAPI

import (

	log "github.com/Sirupsen/logrus"

	"bytes"
	"encoding/json"
	"fmt"		
	"time"
)

const (
NFSvmdiskPath = "/.acropolis/vmdisk/"
Active = "ACTIVE"
)
		
	
type Image_json_AHV struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Annotation string `json:"annotation"`
	Deleted bool `json:"deleted"`
	ContainerID int `json:"containerId"`
	ContainerUUID string `json:"containerUuid"`
	LogicalTimestamp int `json:"logicalTimestamp"`
	ImageType string `json:"imageType"`
	VMDiskID string `json:"vmDiskId"`
	ImageState string `json:"imageState"`
	CreatedTimeInUsecs int64 `json:"createdTimeInUsecs"`
	UpdatedTimeInUsecs int64 `json:"updatedTimeInUsecs"`
}
	

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

func GetImageVMDiskIDbyName(n *NTNXConnection, ImageName string) (string,error) {

	resp, _ := NutanixAPIGet(n,NutanixAHVurl(n), "images")

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

func GetImagebyName(n *NTNXConnection, ImageName string) (Image_json_AHV,error) {

	resp, _ := NutanixAPIGet(n,NutanixAHVurl(n), "images")

	var iml ImageList_AHV
	var im Image_json_AHV

	json.Unmarshal(resp, &iml)

	// TO-DO: Handle if more than one images found
	
	s := iml.Entities

	for i := 0; i < len(s); i++ {
		if s[i].Name == ImageName {
						
			resp_image, _ := NutanixAPIGet(n,NutanixAHVurl(n), "images/"+s[i].UUID)
			json.Unmarshal(resp_image, &im)
			
			return im, nil
		}

	}
	
	log.Warn("Image "+ImageName+" not found")
	return im, fmt.Errorf("Image "+ImageName+" not found")
}


func GetImageStatebyUUID(n *NTNXConnection, UUID string) (string,error) {

	resp, _ := NutanixAPIGet(n,NutanixAHVurl(n), "images/"+UUID)

	var im Image_json_AHV

	json.Unmarshal(resp, &im)
	
	if ( im.UUID == UUID ) {
		return im.ImageState, nil
		}
	
	log.Warn("Image ID "+UUID+" not found")
	return "", fmt.Errorf("Image ID "+UUID+" not found")
}

// Checks if Image exists and fills helper struct with UUID and VMDiskID
func ImageExistbyName(n *NTNXConnection, im *Image_json_AHV) bool {

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

func CloneCDforVM(n *NTNXConnection, v *VM_json_AHV, im *Image_json_AHV) (TaskUUID,error) {

	var jsonStr = []byte(`{ "disks": [ { "vmDiskClone":  { "vmDiskUuid": "` + im.VMDiskID + `" } , "isCdrom" : "true"} ] }`)
	var task TaskUUID

    log.Debug("Post body: "+string(jsonStr))

	resp , statusCode := NutanixAPIPost(n, NutanixAHVurl(n), "vms/"+v.UUID+"/disks/", bytes.NewBuffer(jsonStr))
		
	if ( statusCode == 200) {
		json.Unmarshal(resp, &task)
		return task, nil
	 }	
 log.Warn("CD ISO Image "+im.Name+" could not cloned for VM "+v.Config.Name)
 return task,fmt.Errorf("CD ISO Image "+im.Name+" could not cloned for VM "+v.Config.Name) 
}

func CloneDiskforVM(n *NTNXConnection, v *VM_json_AHV, im *Image_json_AHV) (TaskUUID,error) {

	var jsonStr = []byte(`{ "disks": [ { "vmDiskClone":  { "vmDiskUuid": "` + im.VMDiskID + `" } } ] }`)
	var task TaskUUID	
	
	log.Debug("Post body: "+string(jsonStr))

	resp , statusCode := NutanixAPIPost(n, NutanixAHVurl(n), "vms/"+v.UUID+"/disks/", bytes.NewBuffer(jsonStr))
		
	if ( statusCode == 200) {
		json.Unmarshal(resp, &task)
		return task, nil
	 }	
 log.Warn("Image "+im.Name+" could not cloned for VM "+v.Config.Name)
 return task,fmt.Errorf("Image "+im.Name+" could not cloned for VM "+v.Config.Name)
	 
}


// Waits until an Image is actice - timeouts afters 30s
func WaitUntilImageIsActive(n *NTNXConnection, im *Image_json_AHV) (bool,error) {
	
	starttime := time.Now()	
	
	var imageState string
	
	
	for  time.Since(starttime)/1000/1000/1000 < 30 {
		
	 imageState, _ = GetImageStatebyUUID(n,im.UUID)
	 
	 log.Debug(imageState)	
	 	 
	 
	 if ( imageState == Active ) { 
		 return true, nil
	 } 
	}
	  
	log.Warn("Image "+im.UUID+" is not active and timedout")
	return false, fmt.Errorf("Image "+im.UUID+" is not active and timedout")	
}

func GenerateNFSURIfromVDisk(host string,container_name string, VMDiskID string) string {
 
  return "nfs://"+host+"/"+container_name+NFSvmdiskPath+VMDiskID

}

func CreateImageFromURL(n *NTNXConnection, d *VDisk_json_REST, im *Image_json_AHV) (TaskUUID,error) {
	
	SourceContainerName, err := GetContainerNamebyUUID(n,d.ContainerID)
	if err != nil {
    log.Fatal(err)
	}	
		
	var jsonStr = []byte(`{ "name": "`+im.Name+`","annotation": "`+im.Annotation+`", "imageType":"DISK_IMAGE", "imageImportSpec": {"containerUuid ": "`+im.ContainerUUID+`","url":"`+GenerateNFSURIfromVDisk(n.NutanixHost,SourceContainerName,d.VdiskUUID)+`"} }`)
	var task TaskUUID

	resp , statusCode := NutanixAPIPost(n, NutanixAHVurl(n), "images", bytes.NewBuffer(jsonStr))	
	
	if ( statusCode == 200) {
		json.Unmarshal(resp, &task)
		return task, nil
	 }	
  
  log.Warn("Image "+im.Name+" could not created on container ID "+im.ContainerUUID+" from "+GenerateNFSURIfromVDisk(n.NutanixHost,SourceContainerName,d.VdiskUUID))
  return task,fmt.Errorf("Image "+im.Name+" could not created on container ID "+im.ContainerUUID+" from "+GenerateNFSURIfromVDisk(n.NutanixHost,SourceContainerName,d.VdiskUUID))
}

func CreateImageObject(n *NTNXConnection, im *Image_json_AHV) (TaskUUID,error) {
				
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

func GetImageUUIDbyTask(n *NTNXConnection, t *Task_json_REST) string {
	
	return t.EntityList[0].UUID
		
}




