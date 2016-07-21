package ntnxAPI

import (

	"encoding/json"
	"bytes"
	"fmt"
)

type VM_json_AHV struct {
	Config struct {
		MemoryMb int    `json:"memoryMb"`
		Name     string `json:"name"`
		NumVcpus int    `json:"numVcpus"`
		VMDisks  []struct {
			Addr struct {
				DeviceBus   string `json:"deviceBus"`
				DeviceIndex int    `json:"deviceIndex"`
			} `json:"addr"`
			ContainerID       int    `json:"containerId"`
			ContainerUUID     string `json:"containerUuid"`
			ID                string `json:"id"`
			IsCdrom           bool   `json:"isCdrom"`
			IsEmpty           bool   `json:"isEmpty"`
			IsSCSIPassthrough bool   `json:"isSCSIPassthrough"`
			SourceImage       string `json:"sourceImage"`
			SourceVMDiskUUID  string `json:"sourceVmDiskUuid"`
			VMDiskUUID        string `json:"vmDiskUuid"`
		} `json:"vmDisks"`
		VMNics []struct {
			MacAddress  string `json:"macAddress"`
			Model       string `json:"model"`
			NetworkUUID string `json:"networkUuid"`
		} `json:"vmNics"`
	} `json:"config"`
	HostUUID         string `json:"hostUuid"`
	LogicalTimestamp int    `json:"logicalTimestamp"`
	State            string `json:"state"`
	UUID             string `json:"uuid"`
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

type VMList_AHV struct {
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


func VMExist(n *NTNXConnection, v *VM) (bool) {
	
	resp := NutanixAPIGet(n,NutanixAHVurl(n),"vms")
	
	var vl VMList_AHV
	
	
	if err := json.Unmarshal(resp, &vl)	; err != nil {
		panic(err)		
	}
	
	s := vl.Entities
	
	for i:= 0; i < len(s); i++ {
		if s[i].Config.Name == v.Name {
			return true
		}
		
	}
	
	return false
	
} 

func GetVMIDbyName(n *NTNXConnection, Name string) (string, error) {
	//VM Names are not unique. Returns the last found 
	//raises an error if more than one found
	
	resp := NutanixAPIGet(n,NutanixAHVurl(n),"vms")
	
	var vl VMList_AHV
	
	if err := json.Unmarshal(resp, &vl)	; err != nil {
		panic(err)		
	}
		
	s := vl.Entities
				
	// TODO FilterCriteria seems not to work in 4.5 
	// Return error when > 1 found and not found
	var c int = 0
	
	for i:= 0; i < len(s); i++ {
		if s[i].Config.Name == Name {			
			c++
			// if the last one in the response is reached
			if (i == len(s)-1) {
				// and more than one is found	
				if  (c>1) {
				 return s[i].UUID, fmt.Errorf("NOT UNIQUE")
				} else { // exact one is found
				return s[i].UUID, nil
				}
			}
		}
	}	
	   		
    return "",fmt.Errorf("NOT FOUND")
    		
}

func CreateVM(n *NTNXConnection, v *VM) {
	
	var jsonStr = []byte(`{"memoryMb": "`+v.MemoryMB+`", "name": "`+v.Name+`", "numVcpus": "`+v.Vcpus+`"}`)	
	
	resp := NutanixAPIPost(n,NutanixAHVurl(n),"vms",bytes.NewBuffer(jsonStr))
	
	fmt.Println(resp);
	
}

func GetVMsbyContainer(n *NTNXConnection,Name string) []string {


	var VMListing []string
	
	containerId_used := GetContainerIDbyName(n,Name)
	
	
	fmt.Println(NutanixRestURL(n),"vms/")

    resp := NutanixAPIGet(n,NutanixRestURL(n),"vms/")
	
	var f interface{}
	
	json.Unmarshal(resp, &f)
	
	m := f.(map[string]interface{})
	
	e := m["entities"].([]interface{})
	
	for k := range e {
		t := e[k].(map[string]interface{})		
		 if t["containerIds"] != nil {
		  containerIds := t["containerIds"].([]interface{})
		   for _,element := range containerIds {
		   	if element.(string) == containerId_used {
				VMListing = append(VMListing,t["vmId"].(string))
			}
		   }
		 }	
     }
			
	return VMListing
}


func GetVMState(n *NTNXConnection,vm *VM) string {
	
	resp := NutanixAPIGet(n,NutanixAHVurl(n),"vms/"+vm.VmId)
	
	var vm_AHV VM_json_AHV
	
	json.Unmarshal(resp, &vm_AHV)	
	
	return vm_AHV.State	
	
}

func GetVMIP(n *NTNXConnection,vm *VM) (string, error) {
	
	resp := NutanixAPIGet(n,NutanixRestURL(n),"vms/"+vm.VmId)
	fmt.Println(NutanixRestURL(n),"vms/"+vm.VmId)
	
	var vm_REST VM_json_REST	
	
	if err := json.Unmarshal(resp, &vm_REST); err != nil {
		panic(err)		
	}

	if (len(vm_REST.IPAddresses)>0) {
		fmt.Println(vm_REST.IPAddresses[0])
		return vm_REST.IPAddresses[0], nil
		}	
	
	return "", fmt.Errorf("NO IP FOUND")
}

func CreateVDiskforVM (n *NTNXConnection,v *VM, d *VDisk) {
	
		
	var jsonStr = []byte(`{ "disks": [  { "vmDiskCreate":  { "sizeMb": "`+d.MaxCapacityBytes+`", "containerId": "`+d.ContainerID+`"}} ] }`)	
	
	fmt.Println(string(jsonStr)+"vms/"+v.VmId+"/disks/");
	
	resp := NutanixAPIPost(n,NutanixAHVurl(n),"vms/"+v.VmId+"/disks/",bytes.NewBuffer(jsonStr))
	
	fmt.Println(resp);

}

func CreateVNicforVM (n *NTNXConnection,v *VM, net *Network) {
	
	var jsonStr = []byte(`{ "specList": [ {"networkUuid": "`+net.UUID+`"} ] }`)
		
	resp := NutanixAPIPost(n,NutanixAHVurl(n),"vms/"+v.VmId+"/nics/",bytes.NewBuffer(jsonStr))
	
	fmt.Println(string(resp))
	
}


func StartVM(n *NTNXConnection,v *VM)  {
	
	var jsonStr = []byte(`{}`)
	
	
	resp := NutanixAPIPost(n,NutanixAHVurl(n),"vms/"+v.VmId+"/power_op/on",bytes.NewBuffer(jsonStr))
	
	fmt.Println(string(resp))
	
}

func StopVM(n *NTNXConnection,v *VM) {
	
	var jsonStr = []byte(`{}`)
	
	
	resp := NutanixAPIPost(n,NutanixAHVurl(n),"vms/"+v.VmId+"/power_op/off",bytes.NewBuffer(jsonStr))
	
	fmt.Println(string(resp))
}

func DeleteVM(n *NTNXConnection,v *VM) {
	
	var jsonStr = []byte(`{}`)
	
	resp := NutanixAPIPost(n,NutanixAHVurl(n),"vms/"+v.VmId,bytes.NewBuffer(jsonStr))
	
	fmt.Println(string(resp))
}

