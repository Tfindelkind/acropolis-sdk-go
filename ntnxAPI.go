package ntnxAPI

import (
	log "github.com/Sirupsen/logrus"

	"bytes"
	"crypto/tls"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

const appVersion = "0.9 beta"

// NTNXConnection ...
type NTNXConnection struct {
	NutanixHost string
	Username    string
	Password    string
	SEnc        string
	HTTPClient  http.Client
}

// EncodeCredentials ...
func EncodeCredentials(n *NTNXConnection) {
	n.SEnc = base64.StdEncoding.EncodeToString([]byte(n.Username + ":" + n.Password))
}

// NutanixAHVurl ...
func NutanixAHVurl(n *NTNXConnection) string {

	return "https://" + n.NutanixHost + ":9440/api/nutanix/v0.8/"

}

// NutanixRestURL ...
func NutanixRestURL(n *NTNXConnection) string {

	return "https://" + n.NutanixHost + ":9440/PrismGateway/services/rest/v1/"

}

// DebugRequest ...
func DebugRequest(req *http.Request) {

	log.Debug(req.Method)
	log.Debug(req.URL)
	log.Debug(req.Header)
}

// DebugResponse ...
func DebugResponse(resp *http.Response, bodyText []byte) {

	log.Debug(resp.StatusCode, string(bodyText))
}

// CreateHTTPClient ...
func CreateHTTPClient(n *NTNXConnection) {

	// Ignore certificats which can not be validated (Nutanix CE edition)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	n.HTTPClient = http.Client{Transport: tr}
}

// NutanixAPIGet ...
func NutanixAPIGet(n *NTNXConnection, NutanixAPIurl string, NutanixURI string) ([]byte, int) {

	var req *http.Request
	var err error

	req, err = http.NewRequest("GET", NutanixAPIurl+NutanixURI, nil)
	req.Header.Set("Authorization", "Basic "+n.SEnc)

	DebugRequest(req)

	resp, err := n.HTTPClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	DebugResponse(resp, bodyText)

	return bodyText, resp.StatusCode

}

// NutanixAPIPost ...
func NutanixAPIPost(n *NTNXConnection, NutanixAPIurl string, NutanixURI string, body *bytes.Buffer) ([]byte, int) {

	var req *http.Request
	var err error

	req, err = http.NewRequest("POST", NutanixAPIurl+NutanixURI, body)
	req.Header.Set("Authorization", "Basic "+n.SEnc)

	DebugRequest(req)

	resp, err := n.HTTPClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	DebugResponse(resp, bodyText)

	return bodyText, resp.StatusCode

}

// NutanixAPIDelete ...
func NutanixAPIDelete(n *NTNXConnection, NutanixAPIurl string, NutanixURI string) ([]byte, int) {

	var req *http.Request
	var err error

	req, err = http.NewRequest("DELETE", NutanixAPIurl+NutanixURI, nil)
	req.Header.Set("Authorization", "Basic "+n.SEnc)

	DebugRequest(req)

	resp, err := n.HTTPClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	DebugResponse(resp, bodyText)

	return bodyText, resp.StatusCode

}

// NutanixCheckCredentials ...
func NutanixCheckCredentials(n *NTNXConnection) {

	_, statusCode := NutanixAPIGet(n, NutanixRestURL(n), "cluster")

	if statusCode == 401 {
		log.Fatal("Username or password not valid for host: " + n.NutanixHost)
		os.Exit(1)
	}

	if statusCode != 200 {
		log.Fatal("Connection to host: " + n.NutanixHost + " not possible")
		os.Exit(1)
	}
}

// PutFileToImage ...
func PutFileToImage(n *NTNXConnection, NutanixAPIurl string, NutanixURI string, filename string, containerName string) ([]byte, int) {

	var req *http.Request

	// Open file which will be send via PUT
	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// file Stat().Size is needed to set Content-Length
	fStat, err := f.Stat()
	if err != nil {
		// Could not obtain stat, handle error
		log.Fatal(err)
	}

	log.Debug("The file " + filename + " is " + strconv.FormatInt(fStat.Size(), 10) + " bytes long")

	req, err = http.NewRequest("PUT", NutanixAPIurl+NutanixURI, f)

	containerID, _ := GetContainerUUIDbyName(n, containerName)

	req.ContentLength = fStat.Size()
	req.Header.Set("Authorization", "Basic "+n.SEnc)
	req.Header.Set("X-Nutanix-Destination-Container", containerID)
	req.Header.Set("Content-Type", "application/octet-stream;charset=UTF-8")
	req.Header.Set("Content-Length", strconv.FormatInt(fStat.Size(), 10))
	req.Header.Set("Accept-Language", "de-DE,de;q=0.8,en-US;q=0.6,en;q=0.4")
	req.Header.Set("Connection", "keep-alive")

	DebugRequest(req)

	log.Info("Uploading file " + filename + " ...")

	resp, err := n.HTTPClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	DebugResponse(resp, bodyText)

	return bodyText, resp.StatusCode

}

// GetCluster ...
func GetCluster(n *NTNXConnection) []byte {

	resp, _ := NutanixAPIGet(n, NutanixRestURL(n), "cluster")

	return resp

}
