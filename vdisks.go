package ntnxAPI

import (

	log "github.com/Sirupsen/logrus"
	
	"bytes"
	"encoding/json"
	"fmt"
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

type VDiskList_json_REST []struct {
	Name string `json:"name"`
	ClusterUUID string `json:"clusterUuid"`
	ContainerID string `json:"containerId"`
	ContainerUUID string `json:"containerUuid"`
	ContainerName string `json:"containerName"`
	StoragePoolID string `json:"storagePoolId"`
	StoragePoolUUID string `json:"storagePoolUuid"`
	StoragePoolName string `json:"storagePoolName"`
	Shared bool `json:"shared"`
	NfsFile bool `json:"nfsFile"`
	Immutable bool `json:"immutable"`
	Snapshot bool `json:"snapshot"`
	QosPriority interface{} `json:"qosPriority"`
	QosFairshare interface{} `json:"qosFairshare"`
	ErasureCode string `json:"erasureCode"`
	ErasureCodeDelaySecs interface{} `json:"erasureCodeDelaySecs"`
	FingerPrintOnWrite string `json:"fingerPrintOnWrite"`
	OnDiskDedup string `json:"onDiskDedup"`
	TotalReservedCapacityBytes interface{} `json:"totalReservedCapacityBytes"`
	MaxCapacityBytes int64 `json:"maxCapacityBytes"`
	IscsiTargetName string `json:"iscsiTargetName"`
	IscsiLun int `json:"iscsiLun"`
	NfsFileName string `json:"nfsFileName"`
	ParentNfsFileName string `json:"parentNfsFileName"`
	Disabled bool `json:"disabled"`
	MarkedForRemoval bool `json:"markedForRemoval"`
	CreationTimeInMicrosSinceEpoch int64 `json:"creationTimeInMicrosSinceEpoch"`
	Snapshots []interface{} `json:"snapshots"`
	Stats struct {
		HypervisorAvgIoLatencyUsecs string `json:"hypervisor_avg_io_latency_usecs"`
		NumReadIops string `json:"num_read_iops"`
		HypervisorWriteIoBandwidthKBps string `json:"hypervisor_write_io_bandwidth_kBps"`
		TimespanUsecs string `json:"timespan_usecs"`
		ControllerNumReadIops string `json:"controller_num_read_iops"`
		ReadIoPpm string `json:"read_io_ppm"`
		ControllerNumIops string `json:"controller_num_iops"`
		TotalReadIoTimeUsecs string `json:"total_read_io_time_usecs"`
		ControllerTotalReadIoTimeUsecs string `json:"controller_total_read_io_time_usecs"`
		HypervisorNumIo string `json:"hypervisor_num_io"`
		ControllerTotalTransformedUsageBytes string `json:"controller_total_transformed_usage_bytes"`
		ControllerNumWriteIo string `json:"controller_num_write_io"`
		AvgReadIoLatencyUsecs string `json:"avg_read_io_latency_usecs"`
		ControllerTotalIoTimeUsecs string `json:"controller_total_io_time_usecs"`
		ControllerTotalReadIoSizeKbytes string `json:"controller_total_read_io_size_kbytes"`
		ControllerNumSeqIo string `json:"controller_num_seq_io"`
		ControllerReadIoPpm string `json:"controller_read_io_ppm"`
		ControllerTotalIoSizeKbytes string `json:"controller_total_io_size_kbytes"`
		ControllerNumIo string `json:"controller_num_io"`
		HypervisorAvgReadIoLatencyUsecs string `json:"hypervisor_avg_read_io_latency_usecs"`
		NumWriteIops string `json:"num_write_iops"`
		ControllerNumRandomIo string `json:"controller_num_random_io"`
		NumIops string `json:"num_iops"`
		HypervisorNumReadIo string `json:"hypervisor_num_read_io"`
		HypervisorTotalReadIoTimeUsecs string `json:"hypervisor_total_read_io_time_usecs"`
		ControllerAvgIoLatencyUsecs string `json:"controller_avg_io_latency_usecs"`
		NumIo string `json:"num_io"`
		ControllerNumReadIo string `json:"controller_num_read_io"`
		HypervisorNumWriteIo string `json:"hypervisor_num_write_io"`
		ControllerSeqIoPpm string `json:"controller_seq_io_ppm"`
		ControllerReadIoBandwidthKBps string `json:"controller_read_io_bandwidth_kBps"`
		ControllerIoBandwidthKBps string `json:"controller_io_bandwidth_kBps"`
		HypervisorTimespanUsecs string `json:"hypervisor_timespan_usecs"`
		HypervisorNumWriteIops string `json:"hypervisor_num_write_iops"`
		TotalReadIoSizeKbytes string `json:"total_read_io_size_kbytes"`
		HypervisorTotalIoSizeKbytes string `json:"hypervisor_total_io_size_kbytes"`
		AvgIoLatencyUsecs string `json:"avg_io_latency_usecs"`
		HypervisorNumReadIops string `json:"hypervisor_num_read_iops"`
		ControllerWriteIoBandwidthKBps string `json:"controller_write_io_bandwidth_kBps"`
		ControllerWriteIoPpm string `json:"controller_write_io_ppm"`
		HypervisorAvgWriteIoLatencyUsecs string `json:"hypervisor_avg_write_io_latency_usecs"`
		HypervisorTotalReadIoSizeKbytes string `json:"hypervisor_total_read_io_size_kbytes"`
		ReadIoBandwidthKBps string `json:"read_io_bandwidth_kBps"`
		HypervisorNumIops string `json:"hypervisor_num_iops"`
		HypervisorIoBandwidthKBps string `json:"hypervisor_io_bandwidth_kBps"`
		ControllerNumWriteIops string `json:"controller_num_write_iops"`
		TotalIoTimeUsecs string `json:"total_io_time_usecs"`
		ControllerRandomIoPpm string `json:"controller_random_io_ppm"`
		ControllerAvgReadIoSizeKbytes string `json:"controller_avg_read_io_size_kbytes"`
		TotalTransformedUsageBytes string `json:"total_transformed_usage_bytes"`
		AvgWriteIoLatencyUsecs string `json:"avg_write_io_latency_usecs"`
		NumReadIo string `json:"num_read_io"`
		WriteIoBandwidthKBps string `json:"write_io_bandwidth_kBps"`
		HypervisorReadIoBandwidthKBps string `json:"hypervisor_read_io_bandwidth_kBps"`
		RandomIoPpm string `json:"random_io_ppm"`
		TotalUntransformedUsageBytes string `json:"total_untransformed_usage_bytes"`
		HypervisorTotalIoTimeUsecs string `json:"hypervisor_total_io_time_usecs"`
		NumRandomIo string `json:"num_random_io"`
		ControllerAvgWriteIoSizeKbytes string `json:"controller_avg_write_io_size_kbytes"`
		ControllerAvgReadIoLatencyUsecs string `json:"controller_avg_read_io_latency_usecs"`
		NumWriteIo string `json:"num_write_io"`
		TotalIoSizeKbytes string `json:"total_io_size_kbytes"`
		IoBandwidthKBps string `json:"io_bandwidth_kBps"`
		ControllerTimespanUsecs string `json:"controller_timespan_usecs"`
		NumSeqIo string `json:"num_seq_io"`
		SeqIoPpm string `json:"seq_io_ppm"`
		WriteIoPpm string `json:"write_io_ppm"`
		ControllerAvgWriteIoLatencyUsecs string `json:"controller_avg_write_io_latency_usecs"`
	} `json:"stats"`
	UsageStats struct {
	} `json:"usageStats"`
	VdiskUUID interface{} `json:"vdiskUuid"`
}

func GetVDiskIDbyName(n *NTNXConnection, Name string) string {

	resp, _ := NutanixAPIGet(n, NutanixRestURL(n), `vdisks/?vdiskNames=`+Name)

	//remove "[" at begin and end "]" before Unmarshal
	r := resp[1 : len(resp)-1]

	var dl VDisk_json_REST

	json.Unmarshal(r, &dl)

	return dl.VdiskUUID

}


func CreateVDisk(n *NTNXConnection, d *VDisk) (TaskUUID,error) {

	var jsonStr = []byte(`{"containerId": "` + d.ContainerID + `", "name": "` + d.Name + `", "maxCapacityBytes": "` + d.MaxCapacityBytes + `"}`)
	var task TaskUUID

	resp, statusCode := NutanixAPIPost(n, NutanixRestURL(n), "vdisks", bytes.NewBuffer(jsonStr))
	
	if ( statusCode == 200 ) {
	   json.Unmarshal(resp, &task)	
       return task, nil
    } 
  
  log.Warn("vDisk "+d.Name+" could not created on container ID "+d.ContainerID)
  return task, fmt.Errorf("vDisk "+d.Name+" could not created on container ID "+d.ContainerID)
	
}

func CloneVDiskforVM(n *NTNXConnection, v *VM, VMDiskUUID string, ContainerID string) (TaskUUID,error) {

	var jsonStr = []byte(`{ "disks": [ { "vmDiskClone":  { "containerUuid": "` + ContainerID + `" "vmDiskUuid": "` + VMDiskUUID + `" } , "isCdrom" : "false"} ] }`)
	var task TaskUUID

	resp, statusCode := NutanixAPIPost(n, NutanixAHVurl(n), "vms/"+v.VmId+"/disks/", bytes.NewBuffer(jsonStr))
	
	if ( statusCode == 200 ) {
	   json.Unmarshal(resp, &task)	
       return task, nil
    } 

  log.Warn("vDisk could not cloned on container ID "+ContainerID+" from vDISK ID "+VMDiskUUID)
  return task, fmt.Errorf("vDisk could not cloned on container ID "+ContainerID+" from vDISK ID "+VMDiskUUID)

}
