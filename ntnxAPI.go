package ntnxAPI

import (

	log "github.com/Sirupsen/logrus"

	"crypto/tls"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"bytes"
	"os"
	"strconv"	
)

type NTNXConnection struct {
	NutanixHost string
	Username    string
	Password    string
	SEnc        string
	HttpClient  http.Client
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

func DebugRequest(req *http.Request) {
	
	log.Debug(req.Method)	
	log.Debug(req.URL)
	log.Debug(req.Header)
}

func DebugResponse(resp *http.Response,bodyText []byte) {
			
	log.Debug(resp.StatusCode,string(bodyText))
}

	

func CreateHttpClient(n *NTNXConnection) {

	// Ignore certificats which can not be validated (Nutanix CE edition)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	n.HttpClient = http.Client{Transport: tr}
}

func NutanixAPIGet(n *NTNXConnection, NutanixAPIurl string, NutanixURI string) ([]byte,int) {

	var req *http.Request
	var err error

	req, err = http.NewRequest("GET", NutanixAPIurl+NutanixURI, nil)
	req.Header.Set("Authorization", "Basic "+n.SEnc)
	
	DebugRequest(req)

	resp, err := n.HttpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	DebugResponse(resp,bodyText)

	return bodyText, resp.StatusCode

}

func NutanixAPIPost(n *NTNXConnection, NutanixAPIurl string, NutanixURI string, body *bytes.Buffer) ([]byte,int) {

	var req *http.Request
	var err error

	req, err = http.NewRequest("POST", NutanixAPIurl+NutanixURI, body)
	req.Header.Set("Authorization", "Basic "+n.SEnc)

	DebugRequest(req)
	
	resp, err := n.HttpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
		
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	DebugResponse(resp,bodyText)

	return bodyText, resp.StatusCode

}

func NutanixAPIDelete(n *NTNXConnection, NutanixAPIurl string, NutanixURI string) ([]byte,int) {

	var req *http.Request
	var err error

	req, err = http.NewRequest("DELETE", NutanixAPIurl+NutanixURI, nil)
	req.Header.Set("Authorization", "Basic "+n.SEnc)
	
	DebugRequest(req)

	resp, err := n.HttpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	DebugResponse(resp,bodyText)

	return bodyText, resp.StatusCode

}


func PutFileToImage(n *NTNXConnection, NutanixAPIurl string, NutanixURI string, filename string, containerName string) ([]byte,int) {

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

	log.Debug("The file "+filename+" is "+strconv.FormatInt(fStat.Size(),10)+" bytes long")
		
	req, err = http.NewRequest("PUT", NutanixAPIurl+NutanixURI, f)
	
	containerId, _ := GetContainerUUIDbyName(n,containerName)
	
	req.ContentLength = fStat.Size()
	req.Header.Set("Authorization", "Basic "+n.SEnc)
	req.Header.Set("X-Nutanix-Destination-Container",containerId)
	req.Header.Set("Content-Type","application/octet-stream;charset=UTF-8")
	req.Header.Set("Content-Length",strconv.FormatInt(fStat.Size(),10))
	req.Header.Set("Accept-Language","de-DE,de;q=0.8,en-US;q=0.6,en;q=0.4")
	req.Header.Set("Connection","keep-alive")
	
	DebugRequest(req)
	
	log.Info("Uploading file "+filename+" ...")

	resp, err := n.HttpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	DebugResponse(resp,bodyText)

	return bodyText, resp.StatusCode

}

func GetCluster(n *NTNXConnection) []byte {

	resp , _ := NutanixAPIGet(n, NutanixRestURL(n), "cluster")
	
	return resp

}
