package ntnxAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
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

func GetImageIDbyName(n *NTNXConnection, Name string) string {

	fmt.Println(NutanixAHVurl(n), "images/?filterCriteria=name%3D%3D"+Name)

	resp := NutanixAPIGet(n,NutanixAHVurl(n), "images/?filterCriteria=name%3D%3D"+Name)

	var iml ImageList_AHV

	json.Unmarshal(resp, &iml)

	s := iml.Entities

	for i := 0; i < len(s); i++ {
		if s[i].Name == Name {
			//im.UUID = s[i].UUID
			//im.VMDiskID = s[i].VMDiskID
			return s[i].VMDiskID
		}

	}

	return "NOT FOUND"
}

func ImageExist(n *NTNXConnection, im *Image) bool {

	// Image names are not unique so using FilterCriteria could return > 1 value
	resp := NutanixAPIGet(n, NutanixAHVurl(n), "images")

	var iml ImageList_AHV

	json.Unmarshal(resp, &iml)

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

	fmt.Println(string(jsonStr))

	resp := NutanixAPIPost(n, NutanixAHVurl(n), "vms/"+v.VmId+"/disks/", bytes.NewBuffer(jsonStr))

	fmt.Println(resp)

}

func GenerateNFSURIfromVDisk(host string,container_name string, VMDiskID string) string {
 
  return "nfs://"+host+"/"+container_name+NFSvmdiskPath+VMDiskID

}

func CreateImageFromURL(n *NTNXConnection, d *VDisk, im *Image) error {
	
	SourceContainerName, err := GetContainerNamebyUUID(n,d.ContainerUUID)
	
	if err != nil {
    log.Fatal(err)
	}	
		
	var jsonStr = []byte(`{ "name": "`+im.Name+`","annotation": "`+im.Annotation+`", "imageType":"DISK_IMAGE", "imageImportSpec": {"containerName": "`+im.ContainerName+`","url":"`+GenerateNFSURIfromVDisk(n.NutanixHost,SourceContainerName,d.VdiskUuid)+`"} }`)

	fmt.Println(string(jsonStr))

	resp := NutanixAPIPost(n, NutanixAHVurl(n), "images", bytes.NewBuffer(jsonStr))

	fmt.Println(resp)	
	
 return nil
}
