package main

import (
  "net/http"
  "fmt"
  "io/ioutil"
  "log"
  "encoding/json"
  "flag"
  "net"
)

var CHECKIP_URL = "http://whatismyip.herokuapp.com/"

type AwsAccessIdentifier struct {
  AccessKey string `json:"access_key"`
  SecretKey string `json:"secret_key"`
}

func ParseAccessIdentifierJSON(data []byte) (AwsAccessIdentifier, error) {
  var accessIdentifier AwsAccessIdentifier
  err := json.Unmarshal(data, &accessIdentifier)
  return accessIdentifier, err
}

func ReadLocalFile(location string) string {
  contents, err := ioutil.ReadFile(location)
  if err != nil {
    log.Fatal("ReadLocalFile Error:", contents, err)
  }
  return (string(contents))
}

func fetchWanIP() net.IP {
  ipAddress := ReadRemoteBody(CHECKIP_URL)
  return net.ParseIP(ipAddress)
}

func ReadRemoteBody(url string) string {
  resp, err := http.Get(url)
  defer resp.Body.Close()

  if err != nil {
    log.Fatal("Error:", url, err)
  }

  if resp.StatusCode == 200 {
    bodyBytes, err := ioutil.ReadAll(resp.Body)
    bodyString := string(bodyBytes)
    if err != nil {
      log.Fatal("Error unable to read webpage content", err)
    }
    return bodyString
  }
  log.Fatal("Error unable to read webpage content") 
  return ""
}

var aws_secrets *string = flag.String("secrets-file", "", "/path_to/.your_aws_secrets")

func main() {
  flag.Parse()
  fmt.Println(fetchWanIP())
  
}
