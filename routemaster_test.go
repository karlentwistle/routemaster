package main

import (
  "testing"
  "strings"
)

var ACCESS_IDENTIFIERS = `{
  "access_key" : "022QF06E7MXBSH9DHM02",
  "secret_key" : "kWcrlUX5JEDGM/LtmEENI/aVmYvHNif5zB+d9+ct"
}`

var LOCAL_FILE = `routemaster_test.go`

//TODO: Make this test executable when offline
func TestRawRemoteUrl(t *testing.T) {
  output := ReadRemoteBody("http://checkip.dyndns.org")
  if strings.Contains(output, "Current IP Address:") != true {
    t.Fatal("Webpage did not return", output)
  }
}

func TestParseStringForIP(t *testing.T) {
  input  := "127.0.0.1"
  output := ParseStringForIP(input)
  if output != "127.0.0.1" {
    t.Fatal("Invalid IP: returned", output) 
  }

  input  = "Current IP Address: 176.251.76.232"
  output = ParseStringForIP(input)
  if output != "176.251.76.232" {
    t.Fatal("Invalid IP: returned", output) 
  }

  input  = "<html><head><title>Current IP Check</title></head><body>176.251.76.232</body></html>"
  output = ParseStringForIP(input)
  if output != "176.251.76.232" {
    t.Fatal("Invalid IP: returned", output) 
  }

  // TODO: Implement a real error framework so we can throw a failure and test it
  input  = ""
  output = ParseStringForIP(input)
  if output != "" {
    t.Fatal("Invalid IP: returned", output) 
  }
}

func TestParseAccessIdentifiers(t *testing.T) {
  accessIdentifier, error := ParseAccessIdentifierJSON([]byte(ACCESS_IDENTIFIERS))

  if error != nil {
    t.Fatalf("error parsing JSON", accessIdentifier, error)
  }

  if accessIdentifier.AccessKey != "022QF06E7MXBSH9DHM02" {
    t.Fatalf("incorrect JSON parsing element mismatch", accessIdentifier)  
  }

  if accessIdentifier.SecretKey != "kWcrlUX5JEDGM/LtmEENI/aVmYvHNif5zB+d9+ct" {
    t.Fatalf("incorrect JSON parsing element mismatch", accessIdentifier)  
  }

}

// Testing if the contents of a file contains something
// Using this file as the test case... not sure if this is a good idea
func TestReadLocalFile(t *testing.T) {
  output := ReadLocalFile(LOCAL_FILE)
  if len(output) < 0 {
    t.Fatalf("Unable to read contents of file")
  }
}
