package ntnxAPI

import (

	"encoding/json"
	"fmt"
	"bytes"
)

type VDisk_json_REST struct {
		ClusterUUID                    string        `json:"clusterUuid"`
		ContainerID                    string        `json:"containerId"`
		ContainerName                  string        `json:"containerName"`
		CreationTimeInMicrosSinceEpoch int           `json:"creationTimeInMicrosSinceEpoch"`
		Disabled                       bool          `json:"disabled"`
		ErasureCode                    string        `json:"erasureCode"`
		ErasureCodeDelaySecs           interface{}   `json:"erasureCodeDelaySecs"`
		FingerPrintOnWrite             string        `json:"fingerPrintOnWrite"`
		Immutable                      bool          `json:"immutable"`
		IscsiLun                       int           `json:"iscsiLun"`
		IscsiTargetName                string        `json:"iscsiTargetName"`
		MarkedForRemoval               bool          `json:"markedForRemoval"`
		MaxCapacityBytes               int           `json:"maxCapacityBytes"`
		Name                           string        `json:"name"`
		NfsFile                        bool          `json:"nfsFile"`
		NfsFileName                    string        `json:"nfsFileName"`
		OnDiskDedup                    string        `json:"onDiskDedup"`
		ParentNfsFileName              string        `json:"parentNfsFileName"`
		QosFairshare                   interface{}   `json:"qosFairshare"`
		QosPriority                    interface{}   `json:"qosPriority"`
		Shared                         bool          `json:"shared"`
		Snapshot                       bool          `json:"snapshot"`
		Snapshots                      []interface{} `json:"snapshots"`
		StoragePoolID                  string        `json:"storagePoolId"`
		StoragePoolName                string        `json:"storagePoolName"`
		TotalReservedCapacityBytes     interface{}   `json:"totalReservedCapacityBytes"`
		VdiskUUID                      string        `json:"vdiskUuid"`
}

func GetVDiskIDbyName(n *NTNXConnection,Name string) string {
	
	
	resp := NutanixAPIGet(n,NutanixRestURL(n),`vdisks/?vdiskNames=`+Name)
	
	//remove "[" at begin and end "]" before Unmarshal	
	r := resp[1:len(resp)-1] 
	
	fmt.Println(string(r)) 
	
	var dl VDisk_json_REST
	
	json.Unmarshal(r, &dl)
	
	return dl.VdiskUUID
		
}

func CreateVDisk(n *NTNXConnection,d *VDisk) {	
	
	var jsonStr = []byte(`{"containerId": "`+d.ContainerID+`", "name": "`+d.Name+`", "maxCapacityBytes": "`+d.MaxCapacityBytes+`"}`)
		
	resp := NutanixAPIPost(n,NutanixRestURL(n),"vdisks",bytes.NewBuffer(jsonStr))
	
	fmt.Println(resp);
} 


