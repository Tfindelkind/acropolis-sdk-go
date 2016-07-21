package ntnxAPI

import (

	"encoding/json"
	"fmt"
)

type NetworkList_REST struct {
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


func GetNetworkIDbyName(n *NTNXConnection,Name string) string {

	fmt.Println(NutanixRestURL(n)+"networks/?filterCriteria=name%3D%3D"+Name)

	resp := NutanixAPIGet(n,NutanixAHVurl(n),"networks/?filterCriteria=name%3D%3D"+Name)
	
	var netl NetworkList_REST
		
	json.Unmarshal(resp, &netl)
	
	// TODO check if field is empty or > 1
	return netl.Entities[0].UUID
}
