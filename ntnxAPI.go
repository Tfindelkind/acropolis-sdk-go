package ntnxAPI

import (
	"crypto/tls"
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	//"encoding/json"
	"bytes"
	//"fmt"
)

// Helper structs. They do not represent the full nutanix REST API

type NTNXConnection struct {
	NutanixHost string
	Username    string
	Password    string
	SEnc        string
	HttpClient  http.Client
}

type VM struct {
	MemoryMB string
	Name     string
	Vcpus    string
	VLAN     string
	VmId     string
}

type Task struct {
}

type Image struct {
	Name       string
	Annotation string
	ImageType  string
	UUID       string
	VMDiskID   string
}

type VDisk struct {
	ContainerName    string
	ContainerID      string
	Name             string
	MaxCapacityBytes string
	VdiskUuid        string
	IsCD             bool
}

type Network struct {
	Name   string
	UUID   string
	VlanID int
}

func EncodeCredentials(n *NTNXConnection) {
	n.SEnc = base64.StdEncoding.EncodeToString([]byte(n.Username + ":" + n.Password))
}

func NutanixAHVurl(n *NTNXConnection) string {

	return "https://" + n.NutanixHost + ":9440/api/nutanix/v0.8/"

}

func NutanixRestURL(n *NTNXConnection) string {

	return "https://" + n.NutanixHost + ":9440/PrismGateway/services/rest/v1/"

}

func CreateHttpClient(n *NTNXConnection) {

	// Ignore certificats which can not be validated (Nutanix CE edition)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	n.HttpClient = http.Client{Transport: tr}
}

func NutanixAPIGet(n *NTNXConnection, NutanixAPIurl string, NutanixURI string) []byte {

	var req *http.Request
	var err error

	req, err = http.NewRequest("GET", NutanixAPIurl+NutanixURI, nil)
	req.Header.Set("Authorization", "Basic "+n.SEnc)

	resp, err := n.HttpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)

	return bodyText

}

func NutanixAPIPost(n *NTNXConnection, NutanixAPIurl string, NutanixURI string, body *bytes.Buffer) []byte {

	var req *http.Request
	var err error

	req, err = http.NewRequest("POST", NutanixAPIurl+NutanixURI, body)
	req.Header.Set("Authorization", "Basic "+n.SEnc)

	resp, err := n.HttpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)

	return bodyText

}

func GetCluster(n *NTNXConnection) []byte {

	return NutanixAPIGet(n, NutanixRestURL(n), "cluster")

}
