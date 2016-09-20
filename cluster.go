package ntnxAPI

import (
	"bytes"

	log "github.com/Sirupsen/logrus"

	"encoding/json"
	"fmt"
)

//WhiteListJSON ...
type WhiteListJSON []string

// GetWhiteList ...
func GetWhiteList(n *NTNXConnection) ([]string, error) {

	var wl WhiteListJSON

	resp, statusCode := NutanixAPIGet(n, NutanixRestURL(n), "cluster/nfs_whitelist")

	if statusCode == 200 {
		json.Unmarshal(resp, &wl)
		return wl, nil
	}

	log.Warn("Could not retrieve nfs whitelist")
	return wl, fmt.Errorf("Could not retrieve nfs whitelist")
}

// AddWhiteList ...
func AddWhiteList(n *NTNXConnection, addr string) error {

	var jsonStr = []byte(`{"value": "` + addr + `"}`)
	var wl WhiteListJSON

	resp, statusCode := NutanixAPIPost(n, NutanixRestURL(n), "cluster/nfs_whitelist", bytes.NewBuffer(jsonStr))

	if statusCode == 200 {
		json.Unmarshal(resp, &wl)
		return nil
	}

	log.Warn("WhiteList: " + addr + " could not created on cluster")
	return fmt.Errorf("WhiteList: " + addr + " could not created on cluster")
}
