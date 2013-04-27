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

func parseAccessIdentifierJSON(data []byte) (AwsAccessIdentifier, error) {
  var accessIdentifier AwsAccessIdentifier
  err := json.Unmarshal(data, &accessIdentifier)
  return accessIdentifier, err
}

func readLocalFile(location string) string {
  contents, err := ioutil.ReadFile(location)
  if err != nil {
    log.Fatal("ReadLocalFile Error:", contents, err)
  }
  return (string(contents))
}

func getWanIP(url string) (net.IP, error) {
  resp, err := getBody(url)
  if err == nil {
    return net.ParseIP(string(resp)), nil  
  }
  return nil, err
}

func getBody(url string) ([]byte, error) {
  client := &http.Client{}
  req, err := http.NewRequest("GET", url, nil)
  if err == nil {
    resp, err := client.Do(req)
    defer resp.Body.Close()
    if err == nil {
      return ioutil.ReadAll(resp.Body)
    }
  }
  return nil, err
}

var aws_secrets *string = flag.String("secrets-file", "", "/path_to/.your_aws_secrets")

func main() {
  flag.Parse()
  fmt.Println(getWanIP(CHECKIP_URL))
  
}
