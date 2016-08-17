package ntnxAPI

import (
	"encoding/json"
	"fmt"
)

type Container_json_REST struct {
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

func GetContainer(n *NTNXConnection) []byte {

	return NutanixAPIGet(n, NutanixRestURL(n), "containers")

}

func GetContainerIDbyName(n *NTNXConnection, ContainerName string) (string,error) {

	resp := NutanixAPIGet(n, NutanixRestURL(n), "containers?filterCriteria=container_name%3D%3D"+ContainerName)

	var c Container_json_REST

	json.Unmarshal(resp, &c)

	s := c.Entities

	if len(s) == 0 {
		return "", fmt.Errorf("container not found")
	}

	if len(s) > 1 {
		// return error (container is not unique)

	}

	return s[0].ID, nil
}
