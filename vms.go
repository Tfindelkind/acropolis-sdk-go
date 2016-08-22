package ntnxAPI

import (

	log "github.com/Sirupsen/logrus"

	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

type VM_json_AHV struct {
	UUID string `json:"uuid"`
	LogicalTimestamp int `json:"logicalTimestamp"`
	Config struct {
		Name string `json:"name"`
		Description string `json:"description"`
		NumVcpus int `json:"numVcpus"`
		NumCoresPerVcpu int `json:"numCoresPerVcpu"`
		MemoryMb int `json:"memoryMb"`
		VMDisks []struct {
			Addr struct {
				DeviceBus string `json:"deviceBus"`
				DeviceIndex int `json:"deviceIndex"`
			} `json:"addr"`
			IsCdrom bool `json:"isCdrom"`
			IsEmpty bool `json:"isEmpty"`
			SourceImage string `json:"sourceImage"`
			IsSCSIPassthrough bool `json:"isSCSIPassthrough"`
			ID string `json:"id"`
			VMDiskUUID string `json:"vmDiskUuid,omitempty"`
			SourceVMDiskUUID string `json:"sourceVmDiskUuid,omitempty"`
			ContainerID int `json:"containerId,omitempty"`
			ContainerUUID string `json:"containerUuid,omitempty"`
		} `json:"vmDisks"`
		VMNics []struct {
			MacAddress string `json:"macAddress"`
			NetworkUUID string `json:"networkUuid"`
			Model string `json:"model"`
		} `json:"vmNics"`
	} `json:"config"`
	State string `json:"state"`
}

type VM_json_REST struct {
	AcropolisVM                   bool        `json:"acropolisVm"`
	ClusterUUID                   string      `json:"clusterUuid"`
	ConsistencyGroupName          interface{} `json:"consistencyGroupName"`
	ContainerIds                  []string    `json:"containerIds"`
	ContainerUuids                []string    `json:"containerUuids"`
	ControllerVM                  bool        `json:"controllerVm"`
	CPUReservedInHz               interface{} `json:"cpuReservedInHz"`
	DiskCapacityInBytes           int         `json:"diskCapacityInBytes"`
	Displayable                   bool        `json:"displayable"`
	FingerPrintOnWrite            string      `json:"fingerPrintOnWrite"`
	GuestOperatingSystem          interface{} `json:"guestOperatingSystem"`
	HostID                        string      `json:"hostId"`
	HostName                      string      `json:"hostName"`
	HostUUID                      string      `json:"hostUuid"`
	HypervisorType                string      `json:"hypervisorType"`
	IPAddresses                   []string    `json:"ipAddresses"`
	MemoryCapacityInBytes         int         `json:"memoryCapacityInBytes"`
	MemoryReservedCapacityInBytes int         `json:"memoryReservedCapacityInBytes"`
	NumNetworkAdapters            int         `json:"numNetworkAdapters"`
	NumVCpus                      int         `json:"numVCpus"`
	NutanixGuestTools             struct {
		Applications struct {
			FileLevelRestore bool `json:"file_level_restore"`
			VssSnapshot      bool `json:"vss_snapshot"`
		} `json:"applications"`
		CommunicationLinkActive bool   `json:"communicationLinkActive"`
		Enabled                 bool   `json:"enabled"`
		ToRemove                bool   `json:"toRemove"`
		ToolsMounted            bool   `json:"toolsMounted"`
		VMID                    string `json:"vmId"`
		VMName                  string `json:"vmName"`
		VMUUID                  string `json:"vmUuid"`
	} `json:"nutanixGuestTools"`
	NutanixVirtualDiskIds   []string    `json:"nutanixVirtualDiskIds"`
	NutanixVirtualDiskUuids []string    `json:"nutanixVirtualDiskUuids"`
	NutanixVirtualDisks     []string    `json:"nutanixVirtualDisks"`
	OnDiskDedup             string      `json:"onDiskDedup"`
	PowerState              string      `json:"powerState"`
	ProtectionDomainName    interface{} `json:"protectionDomainName"`
	RunningOnNdfs           bool        `json:"runningOnNdfs"`
	Stats                   struct {
		AvgIoLatencyUsecs                                     string `json:"avg_io_latency_usecs"`
		AvgReadIoLatencyUsecs                                 string `json:"avg_read_io_latency_usecs"`
		AvgWriteIoLatencyUsecs                                string `json:"avg_write_io_latency_usecs"`
		Controller_storageTier_cloud_usageBytes               string `json:"controller.storage_tier.cloud.usage_bytes"`
		Controller_storageTier_das_sata_configuredPinnedBytes string `json:"controller.storage_tier.das-sata.configured_pinned_bytes"`
		Controller_storageTier_das_sata_usageBytes            string `json:"controller.storage_tier.das-sata.usage_bytes"`
		Controller_storageTier_ssd_configuredPinnedBytes      string `json:"controller.storage_tier.ssd.configured_pinned_bytes"`
		Controller_storageTier_ssd_usageBytes                 string `json:"controller.storage_tier.ssd.usage_bytes"`
		ControllerAvgIoLatencyUsecs                           string `json:"controller_avg_io_latency_usecs"`
		ControllerAvgReadIoLatencyUsecs                       string `json:"controller_avg_read_io_latency_usecs"`
		ControllerAvgReadIoSizeKbytes                         string `json:"controller_avg_read_io_size_kbytes"`
		ControllerAvgWriteIoLatencyUsecs                      string `json:"controller_avg_write_io_latency_usecs"`
		ControllerAvgWriteIoSizeKbytes                        string `json:"controller_avg_write_io_size_kbytes"`
		ControllerIoBandwidthKBps                             string `json:"controller_io_bandwidth_kBps"`
		ControllerNumIo                                       string `json:"controller_num_io"`
		ControllerNumIops                                     string `json:"controller_num_iops"`
		ControllerNumRandomIo                                 string `json:"controller_num_random_io"`
		ControllerNumReadIo                                   string `json:"controller_num_read_io"`
		ControllerNumReadIops                                 string `json:"controller_num_read_iops"`
		ControllerNumSeqIo                                    string `json:"controller_num_seq_io"`
		ControllerNumWriteIo                                  string `json:"controller_num_write_io"`
		ControllerNumWriteIops                                string `json:"controller_num_write_iops"`
		ControllerRandomIoPpm                                 string `json:"controller_random_io_ppm"`
		ControllerReadIoBandwidthKBps                         string `json:"controller_read_io_bandwidth_kBps"`
		ControllerReadIoPpm                                   string `json:"controller_read_io_ppm"`
		ControllerSeqIoPpm                                    string `json:"controller_seq_io_ppm"`
		ControllerTimespanUsecs                               string `json:"controller_timespan_usecs"`
		ControllerTotalIoSizeKbytes                           string `json:"controller_total_io_size_kbytes"`
		ControllerTotalIoTimeUsecs                            string `json:"controller_total_io_time_usecs"`
		ControllerTotalReadIoSizeKbytes                       string `json:"controller_total_read_io_size_kbytes"`
		ControllerTotalReadIoTimeUsecs                        string `json:"controller_total_read_io_time_usecs"`
		ControllerTotalTransformedUsageBytes                  string `json:"controller_total_transformed_usage_bytes"`
		ControllerUserBytes                                   string `json:"controller_user_bytes"`
		ControllerWriteIoBandwidthKBps                        string `json:"controller_write_io_bandwidth_kBps"`
		ControllerWriteIoPpm                                  string `json:"controller_write_io_ppm"`
		Hypervisor_cpuReadyTimePpm                            string `json:"hypervisor.cpu_ready_time_ppm"`
		HypervisorAvgIoLatencyUsecs                           string `json:"hypervisor_avg_io_latency_usecs"`
		HypervisorAvgReadIoLatencyUsecs                       string `json:"hypervisor_avg_read_io_latency_usecs"`
		HypervisorAvgWriteIoLatencyUsecs                      string `json:"hypervisor_avg_write_io_latency_usecs"`
		HypervisorCPUUsagePpm                                 string `json:"hypervisor_cpu_usage_ppm"`
		HypervisorIoBandwidthKBps                             string `json:"hypervisor_io_bandwidth_kBps"`
		HypervisorMemoryAssignedBytes                         string `json:"hypervisor_memory_assigned_bytes"`
		HypervisorMemoryUsagePpm                              string `json:"hypervisor_memory_usage_ppm"`
		HypervisorNumIo                                       string `json:"hypervisor_num_io"`
		HypervisorNumIops                                     string `json:"hypervisor_num_iops"`
		HypervisorNumReadIo                                   string `json:"hypervisor_num_read_io"`
		HypervisorNumReadIops                                 string `json:"hypervisor_num_read_iops"`
		HypervisorNumReceivedBytes                            string `json:"hypervisor_num_received_bytes"`
		HypervisorNumTransmittedBytes                         string `json:"hypervisor_num_transmitted_bytes"`
		HypervisorNumWriteIo                                  string `json:"hypervisor_num_write_io"`
		HypervisorNumWriteIops                                string `json:"hypervisor_num_write_iops"`
		HypervisorReadIoBandwidthKBps                         string `json:"hypervisor_read_io_bandwidth_kBps"`
		HypervisorTimespanUsecs                               string `json:"hypervisor_timespan_usecs"`
		HypervisorTotalIoSizeKbytes                           string `json:"hypervisor_total_io_size_kbytes"`
		HypervisorTotalIoTimeUsecs                            string `json:"hypervisor_total_io_time_usecs"`
		HypervisorTotalReadIoSizeKbytes                       string `json:"hypervisor_total_read_io_size_kbytes"`
		HypervisorTotalReadIoTimeUsecs                        string `json:"hypervisor_total_read_io_time_usecs"`
		HypervisorWriteIoBandwidthKBps                        string `json:"hypervisor_write_io_bandwidth_kBps"`
		IoBandwidthKBps                                       string `json:"io_bandwidth_kBps"`
		NumIo                                                 string `json:"num_io"`
		NumIops                                               string `json:"num_iops"`
		NumRandomIo                                           string `json:"num_random_io"`
		NumReadIo                                             string `json:"num_read_io"`
		NumReadIops                                           string `json:"num_read_iops"`
		NumSeqIo                                              string `json:"num_seq_io"`
		NumWriteIo                                            string `json:"num_write_io"`
		NumWriteIops                                          string `json:"num_write_iops"`
		RandomIoPpm                                           string `json:"random_io_ppm"`
		ReadIoBandwidthKBps                                   string `json:"read_io_bandwidth_kBps"`
		ReadIoPpm                                             string `json:"read_io_ppm"`
		SeqIoPpm                                              string `json:"seq_io_ppm"`
		TimespanUsecs                                         string `json:"timespan_usecs"`
		TotalIoSizeKbytes                                     string `json:"total_io_size_kbytes"`
		TotalIoTimeUsecs                                      string `json:"total_io_time_usecs"`
		TotalReadIoSizeKbytes                                 string `json:"total_read_io_size_kbytes"`
		TotalReadIoTimeUsecs                                  string `json:"total_read_io_time_usecs"`
		TotalTransformedUsageBytes                            string `json:"total_transformed_usage_bytes"`
		TotalUntransformedUsageBytes                          string `json:"total_untransformed_usage_bytes"`
		WriteIoBandwidthKBps                                  string `json:"write_io_bandwidth_kBps"`
		WriteIoPpm                                            string `json:"write_io_ppm"`
	} `json:"stats"`
	UsageStats      struct{} `json:"usageStats"`
	UUID            string   `json:"uuid"`
	VdiskFilePaths  []string `json:"vdiskFilePaths"`
	VdiskNames      []string `json:"vdiskNames"`
	VirtualNicIds   []string `json:"virtualNicIds"`
	VirtualNicUuids []string `json:"virtualNicUuids"`
	VMID            string   `json:"vmId"`
	VMName          string   `json:"vmName"`
}

type VMList_json_AHV struct {
	Entities []struct {
		Config struct {
			MemoryMb int    `json:"memoryMb"`
			Name     string `json:"name"`
			NumVcpus int    `json:"numVcpus"`
		} `json:"config"`
		LogicalTimestamp int    `json:"logicalTimestamp"`
		State            string `json:"state"`
		UUID             string `json:"uuid"`
	} `json:"entities"`
	Metadata struct {
		GrandTotalEntities int `json:"grandTotalEntities"`
		TotalEntities      int `json:"totalEntities"`
	} `json:"metadata"`
}

type VMList_json_REST struct {
	Metadata struct {
		GrandTotalEntities int `json:"grandTotalEntities"`
		TotalEntities int `json:"totalEntities"`
		FilterCriteria string `json:"filterCriteria"`
		SortCriteria string `json:"sortCriteria"`
		Page int `json:"page"`
		Count int `json:"count"`
		StartIndex int `json:"startIndex"`
		EndIndex int `json:"endIndex"`
	} `json:"metadata"`
	Entities []struct {
		VMID string `json:"vmId"`
		UUID string `json:"uuid"`
		PowerState string `json:"powerState"`
		VMName string `json:"vmName"`
		GuestOperatingSystem interface{} `json:"guestOperatingSystem"`
		IPAddresses []string `json:"ipAddresses"`
		HypervisorType string `json:"hypervisorType"`
		HostName string `json:"hostName"`
		HostID string `json:"hostId"`
		HostUUID string `json:"hostUuid"`
		ContainerIds []string `json:"containerIds"`
		ContainerUuids []string `json:"containerUuids"`
		NutanixVirtualDisks []string `json:"nutanixVirtualDisks"`
		NutanixVirtualDiskIds []string `json:"nutanixVirtualDiskIds"`
		NutanixVirtualDiskUuids []string `json:"nutanixVirtualDiskUuids"`
		VirtualNicIds []string `json:"virtualNicIds"`
		VirtualNicUuids []string `json:"virtualNicUuids"`
		ClusterUUID string `json:"clusterUuid"`
		MemoryCapacityInBytes int64 `json:"memoryCapacityInBytes"`
		MemoryReservedCapacityInBytes int64 `json:"memoryReservedCapacityInBytes"`
		NumVCpus int `json:"numVCpus"`
		CPUReservedInHz interface{} `json:"cpuReservedInHz"`
		NumNetworkAdapters int `json:"numNetworkAdapters"`
		VdiskNames []string `json:"vdiskNames"`
		VdiskFilePaths []string `json:"vdiskFilePaths"`
		DiskCapacityInBytes int64 `json:"diskCapacityInBytes"`
		ProtectionDomainName interface{} `json:"protectionDomainName"`
		ConsistencyGroupName interface{} `json:"consistencyGroupName"`
		FingerPrintOnWrite string `json:"fingerPrintOnWrite"`
		OnDiskDedup string `json:"onDiskDedup"`
		Stats struct {
			HypervisorAvgIoLatencyUsecs string `json:"hypervisor_avg_io_latency_usecs"`
			NumReadIops string `json:"num_read_iops"`
			HypervisorWriteIoBandwidthKBps string `json:"hypervisor_write_io_bandwidth_kBps"`
			TimespanUsecs string `json:"timespan_usecs"`
			ControllerNumReadIops string `json:"controller_num_read_iops"`
			ControllerStorageTierSsdUsageBytes string `json:"controller.storage_tier.ssd.usage_bytes"`
			ReadIoPpm string `json:"read_io_ppm"`
			ControllerNumIops string `json:"controller_num_iops"`
			HypervisorMemoryAssignedBytes string `json:"hypervisor_memory_assigned_bytes"`
			TotalReadIoTimeUsecs string `json:"total_read_io_time_usecs"`
			ControllerTotalReadIoTimeUsecs string `json:"controller_total_read_io_time_usecs"`
			ControllerStorageTierSsdConfiguredPinnedBytes string `json:"controller.storage_tier.ssd.configured_pinned_bytes"`
			HypervisorNumIo string `json:"hypervisor_num_io"`
			ControllerTotalTransformedUsageBytes string `json:"controller_total_transformed_usage_bytes"`
			HypervisorCPUUsagePpm string `json:"hypervisor_cpu_usage_ppm"`
			ControllerNumWriteIo string `json:"controller_num_write_io"`
			AvgReadIoLatencyUsecs string `json:"avg_read_io_latency_usecs"`
			ControllerTotalIoTimeUsecs string `json:"controller_total_io_time_usecs"`
			ControllerTotalReadIoSizeKbytes string `json:"controller_total_read_io_size_kbytes"`
			ControllerNumSeqIo string `json:"controller_num_seq_io"`
			ControllerReadIoPpm string `json:"controller_read_io_ppm"`
			ControllerTotalIoSizeKbytes string `json:"controller_total_io_size_kbytes"`
			HypervisorCPUReadyTimePpm string `json:"hypervisor.cpu_ready_time_ppm"`
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
			HypervisorNumReceivedBytes string `json:"hypervisor_num_received_bytes"`
			HypervisorTimespanUsecs string `json:"hypervisor_timespan_usecs"`
			HypervisorNumWriteIops string `json:"hypervisor_num_write_iops"`
			TotalReadIoSizeKbytes string `json:"total_read_io_size_kbytes"`
			HypervisorTotalIoSizeKbytes string `json:"hypervisor_total_io_size_kbytes"`
			AvgIoLatencyUsecs string `json:"avg_io_latency_usecs"`
			HypervisorNumReadIops string `json:"hypervisor_num_read_iops"`
			ControllerWriteIoBandwidthKBps string `json:"controller_write_io_bandwidth_kBps"`
			ControllerWriteIoPpm string `json:"controller_write_io_ppm"`
			ControllerUserBytes string `json:"controller_user_bytes"`
			HypervisorAvgWriteIoLatencyUsecs string `json:"hypervisor_avg_write_io_latency_usecs"`
			HypervisorNumTransmittedBytes string `json:"hypervisor_num_transmitted_bytes"`
			HypervisorTotalReadIoSizeKbytes string `json:"hypervisor_total_read_io_size_kbytes"`
			ReadIoBandwidthKBps string `json:"read_io_bandwidth_kBps"`
			HypervisorMemoryUsagePpm string `json:"hypervisor_memory_usage_ppm"`
			HypervisorNumIops string `json:"hypervisor_num_iops"`
			HypervisorIoBandwidthKBps string `json:"hypervisor_io_bandwidth_kBps"`
			ControllerNumWriteIops string `json:"controller_num_write_iops"`
			TotalIoTimeUsecs string `json:"total_io_time_usecs"`
			ControllerRandomIoPpm string `json:"controller_random_io_ppm"`
			ControllerStorageTierDasSataUsageBytes string `json:"controller.storage_tier.das-sata.usage_bytes"`
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
			ControllerStorageTierDasSataConfiguredPinnedBytes string `json:"controller.storage_tier.das-sata.configured_pinned_bytes"`
			NumWriteIo string `json:"num_write_io"`
			TotalIoSizeKbytes string `json:"total_io_size_kbytes"`
			ControllerStorageTierCloudUsageBytes string `json:"controller.storage_tier.cloud.usage_bytes"`
			IoBandwidthKBps string `json:"io_bandwidth_kBps"`
			ControllerTimespanUsecs string `json:"controller_timespan_usecs"`
			NumSeqIo string `json:"num_seq_io"`
			SeqIoPpm string `json:"seq_io_ppm"`
			WriteIoPpm string `json:"write_io_ppm"`
			ControllerAvgWriteIoLatencyUsecs string `json:"controller_avg_write_io_latency_usecs"`
		} `json:"stats"`
		UsageStats struct {
		} `json:"usageStats"`
		NutanixGuestTools struct {
			VMID string `json:"vmId"`
			VMUUID string `json:"vmUuid"`
			Enabled bool `json:"enabled"`
			Applications struct {
				VssSnapshot bool `json:"vss_snapshot"`
				FileLevelRestore bool `json:"file_level_restore"`
			} `json:"applications"`
			ToolsMounted bool `json:"toolsMounted"`
			VMName string `json:"vmName"`
			CommunicationLinkActive bool `json:"communicationLinkActive"`
			ToRemove bool `json:"toRemove"`
		} `json:"nutanixGuestTools"`
		ControllerVM bool `json:"controllerVm"`
		Displayable bool `json:"displayable"`
		AcropolisVM bool `json:"acropolisVm"`
		RunningOnNdfs bool `json:"runningOnNdfs"`
		NonNdfsDetails string `json:"nonNdfsDetails,omitempty"`
	} `json:"entities"`
}

// This is an implementation of the same function based on interface with arbitrary data
// see https://blog.golang.org/json-and-go
// I avoid to use it because it may raise exceptions
/* func VMExist(n *NTNXConnection, v *VM) (bool) {

	resp := NutanixAPIGet(n,NutanixAHVurl(n),"vms")

	var f interface{}

	if err := json.Unmarshal(resp, &f)	; err != nil {
		panic(err)
	}

	m := f.(map[string]interface{})

	e := m["entities"].([]interface{})

	for k := range e {
	  t := e[k].(map[string]interface{})
	  c := t["config"].(map[string]interface{})
		if c["name"] == v.Name {
			return true
		}
	}

	return false

} */

func VMExist(n *NTNXConnection, vmName string) (bool, error) {

	resp, _ := NutanixAPIGet(n, NutanixAHVurl(n), "vms")

	var vl VMList_json_AHV

	if err := json.Unmarshal(resp, &vl); err != nil {
		panic(err)
	}

	s := vl.Entities
	
	var c int = 0

	for i := 0; i < len(s); i++ {
		if s[i].Config.Name == vmName {
			c++
		}
	}
	
	if c == 1 { 
		return true, nil 
	} else if c > 1 { 
		log.Warn("VM name "+vmName+" is not unique")
		return true, fmt.Errorf("VM name "+vmName+" is not unique") 
	} else { 
		return false , nil
	}
}

func GetVMIDbyName(n *NTNXConnection, Name string) (string, error) {
	//VM Names are not unique. Returns the last found
	//raises an error if more than one found

	resp, _ := NutanixAPIGet(n, NutanixAHVurl(n), "vms")

	var vl VMList_json_AHV

	json.Unmarshal(resp, &vl)

	s := vl.Entities

	// Return error when > 1 found and VM not found
	var c int = 0
	var last int 

	for i := 0; i < len(s); i++ {		
		if s[i].Config.Name == Name {
			c++
			last = i
			// if  more than one is found
				if c > 1 {
					log.Warn("VM name "+Name+" is not unique")
					return s[last].UUID, fmt.Errorf("VM name "+Name+" is not unique")
				} 
		}
	}
	
	if c == 1 {
	 return s[last].UUID, nil
	} else
	{
	  log.Warn("VM name "+Name+" not found")		
	  return "", fmt.Errorf("VM name "+Name+" not found") 
    } 
}

func GetVMbyName(n *NTNXConnection, v *VM_json_AHV) (VM_json_AHV, error) {
	
	var vm_AHV VM_json_AHV
	var err error
	
	v.UUID, err = GetVMIDbyName(n, v.Config.Name)
	
	if err != nil {
     log.Fatal(err)
     return vm_AHV, err
    }	
	
	resp, _ := NutanixAPIGet(n, NutanixAHVurl(n), "vms/"+v.UUID)

	json.Unmarshal(resp, &vm_AHV)
	
	return vm_AHV, err
}

func GetVMIDbyTask(n *NTNXConnection, t *Task_json_REST) string {
	
	return t.EntityList[0].UUID
		
}

func CreateVM(n *NTNXConnection, v *VM_json_AHV) (TaskUUID,error) {

	var jsonStr = []byte(`{"memoryMb": "` + strconv.Itoa(v.Config.MemoryMb) + `", "name": "` + v.Config.Name + `", "numVcpus": "` + strconv.Itoa(v.Config.NumVcpus) + `"}`)
	var task TaskUUID

	resp, statusCode := NutanixAPIPost(n, NutanixAHVurl(n), "vms", bytes.NewBuffer(jsonStr))
	
	if ( statusCode == 200 ) {
	   json.Unmarshal(resp, &task)	
       return task, nil
    } 
  
	log.Warn("VM "+v.Config.Name+" could not created")
	return task, fmt.Errorf("VM "+v.Config.Name+" could not created")
	
}

func CreateVM_AHV(n *NTNXConnection, v *VM_json_AHV) (TaskUUID,error) {

	var jsonStr = []byte(`{"name": "` + v.Config.Name + `", "description": "` + v.Config.Description + `", "memoryMb": "` + strconv.Itoa(v.Config.MemoryMb) + `", "numVcpus": "` + strconv.Itoa(v.Config.NumVcpus) + `", "numCoresPerVcpu": "` + strconv.Itoa(v.Config.NumCoresPerVcpu)  + `"}`)
	var task TaskUUID
	
	resp, statusCode := NutanixAPIPost(n, NutanixAHVurl(n), "vms", bytes.NewBuffer(jsonStr))	
	
	if ( statusCode == 200 ) {
	   json.Unmarshal(resp, &task)	
       return task, nil
    } 
  
	log.Warn("VM "+v.Config.Name+" could not created")
	return task, fmt.Errorf("VM "+v.Config.Name+" could not created")
	
}

// returns an array of VmIds of VMs reside on a given container
// nil if nothing could found
func GetVMsbyContainer(n *NTNXConnection, container_name string) ([]string, error) {
	
	var  VMListing []string
	
	containerId, err := GetContainerIDbyName(n, container_name)
	if err != nil {
	 return nil , err }

	resp, _ := NutanixAPIGet(n, NutanixRestURL(n), "vms")

	var vl VMList_json_REST

	if err := json.Unmarshal(resp, &vl); err != nil {
		panic(err)
	}

	s := vl.Entities

    for _,ent := range s { 
      for _,con := range ent.ContainerIds {
		if con == containerId {
		  VMListing = append(VMListing,ent.VMID)
		}
	  }
	}	
    
   return VMListing, nil
}

func GetVMState(n *NTNXConnection, vm *VM_json_AHV) string {

	resp, _ := NutanixAPIGet(n, NutanixAHVurl(n), "vms/"+vm.UUID)

	var vm_AHV VM_json_AHV

	json.Unmarshal(resp, &vm_AHV)

	return vm_AHV.State

}

func GetVMIP(n *NTNXConnection, vm *VM_json_AHV) (string, error) {

	resp, _ := NutanixAPIGet(n, NutanixRestURL(n), "vms/"+vm.UUID)

	var vm_REST VM_json_REST

	if err := json.Unmarshal(resp, &vm_REST); err != nil {
		panic(err)
	}

	if len(vm_REST.IPAddresses) > 0 {
		return vm_REST.IPAddresses[0], nil
	}

	log.Warn("No IP found for VM "+vm.UUID)
	return "", fmt.Errorf("No IP found for VM "+vm.UUID)
}

func CreateVDiskforVM(n *NTNXConnection, v *VM_json_AHV, d *VDisk_json_REST) (TaskUUID,error) {

	var jsonStr = []byte(`{ "disks": [  { "vmDiskCreate":  { "sizeMb": "` + strconv.Itoa(d.MaxCapacityBytes) + `", "containerId": "` + d.ContainerID + `"}} ] }`)
	var task TaskUUID

	resp, statusCode := NutanixAPIPost(n, NutanixAHVurl(n), "vms/"+v.UUID+"/disks/", bytes.NewBuffer(jsonStr))
	
	if ( statusCode == 200 ) {
	   json.Unmarshal(resp, &task)	
       return task, nil
    } 
  
  log.Warn("vDisk "+d.Name+" could not created on container ID "+d.ContainerID+" for VM "+v.UUID)
  return task, fmt.Errorf("vDisk "+d.Name+" could not created on container ID "+d.ContainerID+" for VM "+v.UUID)

}

func CreateVNicforVM(n *NTNXConnection, v *VM_json_AHV, net *Network_REST) (TaskUUID,error) {

	var jsonStr = []byte(`{ "specList": [ {"networkUuid": "` + net.UUID + `"} ] }`)
	var task TaskUUID

	resp, statusCode := NutanixAPIPost(n, NutanixAHVurl(n), "vms/"+v.UUID+"/nics/", bytes.NewBuffer(jsonStr))
	
	if ( statusCode == 200 ) {
	   json.Unmarshal(resp, &task)	
       return task, nil
    } 

  log.Warn("vNic could not created for VM "+v.UUID)
  return task, fmt.Errorf("vNic could not created for VM "+v.UUID)

}

func CreateVNicforVMwithMAC(n *NTNXConnection, v *VM_json_AHV, net *Network_REST, macAddress string) (TaskUUID,error) {

	var jsonStr = []byte(`{ "specList": [ {"networkUuid":"` + net.UUID + `" , "macAddress":"` + macAddress+ `" } ] }`)
	var task TaskUUID

	resp, statusCode := NutanixAPIPost(n, NutanixAHVurl(n), "vms/"+v.UUID+"/nics/", bytes.NewBuffer(jsonStr))
	
	if ( statusCode == 200 ) {
	   json.Unmarshal(resp, &task)	
       return task, nil
    } 

  log.Warn("vNic could not created for VM "+v.UUID)
  return task, fmt.Errorf("vNic could not created for VM "+v.UUID)

}

func StartVM(n *NTNXConnection, v *VM_json_AHV) (TaskUUID,error) {

	var jsonStr = []byte(`{}`)
	var task TaskUUID

	resp, statusCode := NutanixAPIPost(n, NutanixAHVurl(n), "vms/"+v.UUID+"/power_op/on", bytes.NewBuffer(jsonStr))
	
	if ( statusCode == 200 ) {
	   json.Unmarshal(resp, &task)	
       return task, nil
    } 

  log.Warn("VM "+v.UUID+" could not started")
  return task, fmt.Errorf("VM "+v.UUID+" could not started")

}

func StopVM(n *NTNXConnection, v *VM_json_AHV)  (TaskUUID,error) {

	var jsonStr = []byte(`{}`)
	var task TaskUUID

	resp, statusCode := NutanixAPIPost(n, NutanixAHVurl(n), "vms/"+v.UUID+"/power_op/off", bytes.NewBuffer(jsonStr))
	
	if ( statusCode == 200 ) {
	   json.Unmarshal(resp, &task)	
       return task, nil
    } 

  log.Warn("VM "+v.UUID+" could not stopped")
  return task, fmt.Errorf("VM "+v.UUID+" could not stopped")

}

func DeleteVM(n *NTNXConnection, v *VM_json_AHV) (TaskUUID,error) {

	var jsonStr = []byte(`{}`)
	var task TaskUUID

	resp, statusCode := NutanixAPIPost(n, NutanixAHVurl(n), "vms/"+v.UUID, bytes.NewBuffer(jsonStr))
	
	if ( statusCode == 200 ) {
	   json.Unmarshal(resp, &task)	
       return task, nil
    } 

  log.Warn("VM "+v.UUID+" could not stopped")
  return task, fmt.Errorf("VM "+v.UUID+" could not stopped")

}
