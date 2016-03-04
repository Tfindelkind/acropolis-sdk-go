package ntnxAPI

import (

	"net/http"	
	"crypto/tls"
	"io/ioutil"
	"log"
	"encoding/base64"
	"encoding/json"
	"bytes"
	"fmt"

)

type NTNXConnection struct {
	
	NutanixHost 	string
	Username		string
	Password		string
	SEnc 			string 
	HttpClient 		http.Client	
}

type VM struct {
	MemoryMB		string 
	Name			string 
	Vcpus			string 
	VLAN			string
	UUID			string
}

type Task struct {
}

type Image struct {
	Name			string
	Annotation		string
	ImageType		string
	UUID			string
	VMDiskID		string
	
}

type VDisk struct {
	ContainerName		string
	ContainerID			string
	Name				string
	MaxCapacityBytes	string
	VdiskUuid 			string
	IsCD				bool
}

type Network struct {
	Name             string 
	UUID             string 
	VlanID           int    
}

type VMList struct {
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

type ImageList struct {
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

type Container_json struct {
	Entities []struct {
		AdvertisedCapacity     interface{} `json:"advertisedCapacity"`
		ClusterUUID            string      `json:"clusterUuid"`
		CompressionDelayInSecs int         `json:"compressionDelayInSecs"`
		CompressionEnabled     bool        `json:"compressionEnabled"`
		ContainerUUID          string      `json:"containerUuid"`
		DownMigrateTimesInSecs struct {
			DAS_SATA int `json:"DAS-SATA"`
			SSD_PCIe int `json:"SSD-PCIe"`
			SSD_SATA int `json:"SSD-SATA"`
		} `json:"downMigrateTimesInSecs"`
		ErasureCode            string        `json:"erasureCode"`
		ErasureCodeDelaySecs   interface{}   `json:"erasureCodeDelaySecs"`
		FingerPrintOnWrite     string        `json:"fingerPrintOnWrite"`
		ID                     string        `json:"id"`
		IlmPolicy              interface{}   `json:"ilmPolicy"`
		MappedRemoteContainers struct{}      `json:"mappedRemoteContainers"`
		MarkedForRemoval       bool          `json:"markedForRemoval"`
		MaxCapacity            int           `json:"maxCapacity"`
		Name                   string        `json:"name"`
		NfsWhitelist           []interface{} `json:"nfsWhitelist"`
		NfsWhitelistInherited  bool          `json:"nfsWhitelistInherited"`
		OnDiskDedup            string        `json:"onDiskDedup"`
		OplogReplicationFactor int           `json:"oplogReplicationFactor"`
		RandomIoPreference     []string      `json:"randomIoPreference"`
		ReplicationFactor      int           `json:"replicationFactor"`
		SeqIoPreference        []string      `json:"seqIoPreference"`
		Stats                  struct {
			AvgIoLatencyUsecs                string `json:"avg_io_latency_usecs"`
			AvgReadIoLatencyUsecs            string `json:"avg_read_io_latency_usecs"`
			AvgWriteIoLatencyUsecs           string `json:"avg_write_io_latency_usecs"`
			ControllerAvgIoLatencyUsecs      string `json:"controller_avg_io_latency_usecs"`
			ControllerAvgReadIoLatencyUsecs  string `json:"controller_avg_read_io_latency_usecs"`
			ControllerAvgReadIoSizeKbytes    string `json:"controller_avg_read_io_size_kbytes"`
			ControllerAvgWriteIoLatencyUsecs string `json:"controller_avg_write_io_latency_usecs"`
			ControllerAvgWriteIoSizeKbytes   string `json:"controller_avg_write_io_size_kbytes"`
			ControllerIoBandwidthKBps        string `json:"controller_io_bandwidth_kBps"`
			ControllerNumIops                string `json:"controller_num_iops"`
			ControllerNumReadIops            string `json:"controller_num_read_iops"`
			ControllerNumWriteIops           string `json:"controller_num_write_iops"`
			ControllerRandomIoPpm            string `json:"controller_random_io_ppm"`
			ControllerReadIoBandwidthKBps    string `json:"controller_read_io_bandwidth_kBps"`
			ControllerWriteIoBandwidthKBps   string `json:"controller_write_io_bandwidth_kBps"`
			HypervisorAvgIoLatencyUsecs      string `json:"hypervisor_avg_io_latency_usecs"`
			HypervisorAvgReadIoLatencyUsecs  string `json:"hypervisor_avg_read_io_latency_usecs"`
			HypervisorAvgWriteIoLatencyUsecs string `json:"hypervisor_avg_write_io_latency_usecs"`
			HypervisorIoBandwidthKBps        string `json:"hypervisor_io_bandwidth_kBps"`
			HypervisorNumIops                string `json:"hypervisor_num_iops"`
			HypervisorNumReadIo              string `json:"hypervisor_num_read_io"`
			HypervisorNumReadIops            string `json:"hypervisor_num_read_iops"`
			HypervisorNumWriteIo             string `json:"hypervisor_num_write_io"`
			HypervisorNumWriteIops           string `json:"hypervisor_num_write_iops"`
			HypervisorReadIoBandwidthKBps    string `json:"hypervisor_read_io_bandwidth_kBps"`
			HypervisorWriteIoBandwidthKBps   string `json:"hypervisor_write_io_bandwidth_kBps"`
			IoBandwidthKBps                  string `json:"io_bandwidth_kBps"`
			NumIops                          string `json:"num_iops"`
			NumReadIops                      string `json:"num_read_iops"`
			NumWriteIops                     string `json:"num_write_iops"`
			RandomIoPpm                      string `json:"random_io_ppm"`
			ReadIoBandwidthKBps              string `json:"read_io_bandwidth_kBps"`
			ReadIoPpm                        string `json:"read_io_ppm"`
			SeqIoPpm                         string `json:"seq_io_ppm"`
			TotalTransformedUsageBytes       string `json:"total_transformed_usage_bytes"`
			TotalUntransformedUsageBytes     string `json:"total_untransformed_usage_bytes"`
			WriteIoBandwidthKBps             string `json:"write_io_bandwidth_kBps"`
			WriteIoPpm                       string `json:"write_io_ppm"`
		} `json:"stats"`
		StoragePoolID                 string `json:"storagePoolId"`
		TotalExplicitReservedCapacity int    `json:"totalExplicitReservedCapacity"`
		TotalImplicitReservedCapacity int    `json:"totalImplicitReservedCapacity"`
		UsageStats                    struct {
			DataReduction_compression_postReductionBytes       string `json:"data_reduction.compression.post_reduction_bytes"`
			DataReduction_compression_preReductionBytes        string `json:"data_reduction.compression.pre_reduction_bytes"`
			DataReduction_compression_savingRatioPpm           string `json:"data_reduction.compression.saving_ratio_ppm"`
			DataReduction_compression_userPostReductionBytes   string `json:"data_reduction.compression.user_post_reduction_bytes"`
			DataReduction_compression_userPreReductionBytes    string `json:"data_reduction.compression.user_pre_reduction_bytes"`
			DataReduction_compression_userSavedBytes           string `json:"data_reduction.compression.user_saved_bytes"`
			DataReduction_dedup_postReductionBytes             string `json:"data_reduction.dedup.post_reduction_bytes"`
			DataReduction_dedup_preReductionBytes              string `json:"data_reduction.dedup.pre_reduction_bytes"`
			DataReduction_dedup_savingRatioPpm                 string `json:"data_reduction.dedup.saving_ratio_ppm"`
			DataReduction_dedup_userPostReductionBytes         string `json:"data_reduction.dedup.user_post_reduction_bytes"`
			DataReduction_dedup_userPreReductionBytes          string `json:"data_reduction.dedup.user_pre_reduction_bytes"`
			DataReduction_dedup_userSavedBytes                 string `json:"data_reduction.dedup.user_saved_bytes"`
			DataReduction_erasureCoding_postReductionBytes     string `json:"data_reduction.erasure_coding.post_reduction_bytes"`
			DataReduction_erasureCoding_preReductionBytes      string `json:"data_reduction.erasure_coding.pre_reduction_bytes"`
			DataReduction_erasureCoding_savingRatioPpm         string `json:"data_reduction.erasure_coding.saving_ratio_ppm"`
			DataReduction_erasureCoding_userPostReductionBytes string `json:"data_reduction.erasure_coding.user_post_reduction_bytes"`
			DataReduction_erasureCoding_userPreReductionBytes  string `json:"data_reduction.erasure_coding.user_pre_reduction_bytes"`
			DataReduction_erasureCoding_userSavedBytes         string `json:"data_reduction.erasure_coding.user_saved_bytes"`
			DataReduction_postReductionBytes                   string `json:"data_reduction.post_reduction_bytes"`
			DataReduction_preReductionBytes                    string `json:"data_reduction.pre_reduction_bytes"`
			DataReduction_savedBytes                           string `json:"data_reduction.saved_bytes"`
			DataReduction_savingRatioPpm                       string `json:"data_reduction.saving_ratio_ppm"`
			DataReduction_userPostReductionBytes               string `json:"data_reduction.user_post_reduction_bytes"`
			DataReduction_userPreReductionBytes                string `json:"data_reduction.user_pre_reduction_bytes"`
			DataReduction_userSavedBytes                       string `json:"data_reduction.user_saved_bytes"`
			Storage_capacityBytes                              string `json:"storage.capacity_bytes"`
			Storage_diskPhysicalUsageBytes                     string `json:"storage.disk_physical_usage_bytes"`
			Storage_freeBytes                                  string `json:"storage.free_bytes"`
			Storage_logicalUsageBytes                          string `json:"storage.logical_usage_bytes"`
			Storage_reservedCapacityBytes                      string `json:"storage.reserved_capacity_bytes"`
			Storage_reservedFreeBytes                          string `json:"storage.reserved_free_bytes"`
			Storage_reservedUsageBytes                         string `json:"storage.reserved_usage_bytes"`
			Storage_unreservedCapacityBytes                    string `json:"storage.unreserved_capacity_bytes"`
			Storage_unreservedFreeBytes                        string `json:"storage.unreserved_free_bytes"`
			Storage_unreservedOwnUsageBytes                    string `json:"storage.unreserved_own_usage_bytes"`
			Storage_unreservedSharedUsageBytes                 string `json:"storage.unreserved_shared_usage_bytes"`
			Storage_unreservedUsageBytes                       string `json:"storage.unreserved_usage_bytes"`
			Storage_usageBytes                                 string `json:"storage.usage_bytes"`
			Storage_userCapacityBytes                          string `json:"storage.user_capacity_bytes"`
			Storage_userContainerOwnUsageBytes                 string `json:"storage.user_container_own_usage_bytes"`
			Storage_userDiskPhysicalUsageBytes                 string `json:"storage.user_disk_physical_usage_bytes"`
			Storage_userFreeBytes                              string `json:"storage.user_free_bytes"`
			Storage_userOtherContainersReservedCapacityBytes   string `json:"storage.user_other_containers_reserved_capacity_bytes"`
			Storage_userReservedCapacityBytes                  string `json:"storage.user_reserved_capacity_bytes"`
			Storage_userReservedFreeBytes                      string `json:"storage.user_reserved_free_bytes"`
			Storage_userReservedUsageBytes                     string `json:"storage.user_reserved_usage_bytes"`
			Storage_userStoragePoolCapacityBytes               string `json:"storage.user_storage_pool_capacity_bytes"`
			Storage_userUnreservedCapacityBytes                string `json:"storage.user_unreserved_capacity_bytes"`
			Storage_userUnreservedFreeBytes                    string `json:"storage.user_unreserved_free_bytes"`
			Storage_userUnreservedOwnUsageBytes                string `json:"storage.user_unreserved_own_usage_bytes"`
			Storage_userUnreservedSharedUsageBytes             string `json:"storage.user_unreserved_shared_usage_bytes"`
			Storage_userUnreservedUsageBytes                   string `json:"storage.user_unreserved_usage_bytes"`
			Storage_userUsageBytes                             string `json:"storage.user_usage_bytes"`
			StorageTier_das_sata_usageBytes                    string `json:"storage_tier.das-sata.usage_bytes"`
			StorageTier_ssd_usageBytes                         string `json:"storage_tier.ssd.usage_bytes"`
		} `json:"usageStats"`
		VstoreNameList []string `json:"vstoreNameList"`
	} `json:"entities"`
	Metadata struct {
		Count              int    `json:"count"`
		EndIndex           int    `json:"endIndex"`
		FilterCriteria     string `json:"filterCriteria"`
		GrandTotalEntities int    `json:"grandTotalEntities"`
		Page               int    `json:"page"`
		SortCriteria       string `json:"sortCriteria"`
		StartIndex         int    `json:"startIndex"`
		TotalEntities      int    `json:"totalEntities"`
	} `json:"metadata"`
}


type VDisk_json struct {
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

type NetworkList struct {
	Entities []struct {
		IPConfig struct {
			DhcpOptions  struct{}      `json:"dhcpOptions"`
			Pool         []interface{} `json:"pool"`
			PrefixLength int           `json:"prefixLength"`
		} `json:"ipConfig"`
		LogicalTimestamp int    `json:"logicalTimestamp"`
		Name             string `json:"name"`
		UUID             string `json:"uuid"`
		VlanID           int    `json:"vlanId"`
	} `json:"entities"`
	Metadata struct {
		GrandTotalEntities int `json:"grandTotalEntities"`
		TotalEntities      int `json:"totalEntities"`
	} `json:"metadata"`
}


func EncodeCredentials (n *NTNXConnection) {
   n.SEnc = base64.StdEncoding.EncodeToString([]byte(n.Username+":"+n.Password))  
}

func NutanixAHVurl(n *NTNXConnection) string {

	return  "https://"+n.NutanixHost+":9440/api/nutanix/v0.8/"
	
}

func NutanixRestURL(n *NTNXConnection) string {

	return  "https://"+n.NutanixHost+":9440/PrismGateway/services/rest/v1/"
	
}


func CreateHttpClient(n *NTNXConnection)  {
	
	// Ignore certificats which can not be validated (Nutanix CE edition)
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
	
	n.HttpClient = http.Client{Transport: tr}
}


func NutanixAPIGet (n *NTNXConnection,NutanixAPIurl string, NutanixURI string) []byte {
	
	
	var req *http.Request
	var err error
	
	
	
	req, err = http.NewRequest("GET", NutanixAPIurl+NutanixURI, nil)
    req.Header.Set("Authorization", "Basic "+n.SEnc)
    
	
	resp, err := n.HttpClient.Do(req)
    if err != nil{
        log.Fatal(err)
    }
    bodyText, err := ioutil.ReadAll(resp.Body)
   
    return bodyText 
	
}


func NutanixAPIPost (n *NTNXConnection,NutanixAPIurl string,NutanixURI string, body *bytes.Buffer) []byte {
	
	var req *http.Request
	var err error
	
	
	req, err = http.NewRequest("POST", NutanixAPIurl+NutanixURI,body)
    req.Header.Set("Authorization", "Basic "+n.SEnc)
	
	resp, err := n.HttpClient.Do(req)
    if err != nil{
        log.Fatal(err)
    }
    bodyText, err := ioutil.ReadAll(resp.Body)
    
    return bodyText
	
}

func GetCluster(n *NTNXConnection) []byte {
	
	return NutanixAPIGet(n,NutanixRestURL(n),"cluster");

}

func GetContainer(n *NTNXConnection) []byte {
	
	
	return NutanixAPIGet(n,NutanixRestURL(n),"containers");

}

func VMExist(n *NTNXConnection, v *VM) bool {
	

	
	resp := NutanixAPIGet(n,NutanixAHVurl(n),"vms")
	
	var vl VMList
	
	json.Unmarshal(resp, &vl)
	
	s := vl.Entities
	
	for i:= 0; i < len(s); i++ {
		if s[i].Config.Name == v.Name {
			return true
		}
		
	}
	
	return false
	
} 

func GetVMIDbyName(n *NTNXConnection, Name string) string {
	//VM Names are not unique. Returns the first found
	
	resp := NutanixAPIGet(n,NutanixAHVurl(n),"vms")
	
	var vl VMList
	
	json.Unmarshal(resp, &vl)	

	s := vl.Entities
				
	// TODO FilterCriteria seems not to work in 4.5 
	// Return error when > 1 found and not found
	for i:= 0; i < len(s); i++ {
		if s[i].Config.Name == Name {			
			return s[i].UUID
		}
	}	
    		
    return ""
    		
}

func CreateVM(n *NTNXConnection, v *VM) {
	
	var jsonStr = []byte(`{"memoryMb": "`+v.MemoryMB+`", "name": "`+v.Name+`", "numVcpus": "`+v.Vcpus+`"}`)	
	
	resp := NutanixAPIPost(n,NutanixAHVurl(n),"vms",bytes.NewBuffer(jsonStr))
	
	fmt.Println(resp);
	
}

func ImageExist(n *NTNXConnection, im *Image) bool {
	
	// Image names are not unique so using FilterCriteria could return > 1 value
	resp := NutanixAPIGet(n,NutanixAHVurl(n),"images")
	
	var iml ImageList
	
	json.Unmarshal(resp, &iml)
	
	s := iml.Entities
	
	for i:= 0; i < len(s); i++ {
		if s[i].Name == im.Name {
			im.UUID = s[i].UUID
			im.VMDiskID = s[i].VMDiskID
			return true
		}
		
	}
	
	return false
	
}

func GetContainerIDbyName(n *NTNXConnection,ContainerName string) string {
	
	resp := NutanixAPIGet(n,NutanixRestURL(n),"containers?filterCriteria=container_name%3D%3D"+ContainerName)
	
	var c Container_json
		
	json.Unmarshal(resp, &c)
	
	s := c.Entities
	
	if len(s) == 0 {
		fmt.Println("container not found")
	}
	
	if len(s) > 1 {
		// return error (container is not unique)
	}
	
	return s[0].ID
}

func CreateVDisk(n *NTNXConnection,d *VDisk) {	
	
	var jsonStr = []byte(`{"containerId": "`+d.ContainerID+`", "name": "`+d.Name+`", "maxCapacityBytes": "`+d.MaxCapacityBytes+`"}`)
		
	resp := NutanixAPIPost(n,NutanixRestURL(n),"vdisks",bytes.NewBuffer(jsonStr))
	
	fmt.Println(resp);
} 

func CreateVDiskforVM (n *NTNXConnection,v *VM, d *VDisk) {
	
		
	var jsonStr = []byte(`{ "disks": [  { "vmDiskCreate":  { "sizeMb": "`+d.MaxCapacityBytes+`", "containerId": "`+d.ContainerID+`" }} ] }`)	
	
	fmt.Println(string(jsonStr)+"vms/"+v.UUID+"/disks/");
	
	resp := NutanixAPIPost(n,NutanixAHVurl(n),"vms/"+v.UUID+"/disks/",bytes.NewBuffer(jsonStr))
	
	fmt.Println(resp);

}

func CloneCDforVM (n *NTNXConnection,v *VM, im *Image) {
		
	var jsonStr = []byte(`{ "disks": [ { "vmDiskClone":  { "vmDiskUuid": "`+im.VMDiskID+`" } , "isCdrom" : "true"} ] }`)
	
	fmt.Println(string(jsonStr))
		
	resp := NutanixAPIPost(n,NutanixAHVurl(n),"vms/"+v.UUID+"/disks/",bytes.NewBuffer(jsonStr))
	
	fmt.Println(resp);
	
}


func GetVDiskIDbyName(n *NTNXConnection,Name string) string {
	
	
	resp := NutanixAPIGet(n,NutanixRestURL(n),`vdisks/?vdiskNames=`+Name)
	
	//remove "[" at begin and end "]" before Unmarshal	
	r := resp[1:len(resp)-1] 
	
	fmt.Println(string(r)) 
	
	var dl VDisk_json
	
	json.Unmarshal(r, &dl)
	
	return dl.VdiskUUID
		
}

func GetNetworkIDbyName(n *NTNXConnection,Name string) string {

	fmt.Println(NutanixRestURL(n)+"networks/?filterCriteria=name%3D%3D"+Name)

	resp := NutanixAPIGet(n,NutanixAHVurl(n),"networks/?filterCriteria=name%3D%3D"+Name)
	
	var netl NetworkList
		
	json.Unmarshal(resp, &netl)
	
	// TODO check if field is empty or > 1
	return netl.Entities[0].UUID
}

func CreateVNicforVM (n *NTNXConnection,v *VM, net *Network) {
	
	var jsonStr = []byte(`{ "specList": [ {"networkUuid": "`+net.UUID+`"} ] }`)
		
	resp := NutanixAPIPost(n,NutanixAHVurl(n),"vms/"+v.UUID+"/nics/",bytes.NewBuffer(jsonStr))
	
	fmt.Println(resp);
}

//func Poll_task(n *NTNX,t task


