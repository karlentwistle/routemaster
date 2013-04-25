package main

import (
  "testing"
  "regexp"
  "strings"
)

var ACCESS_IDENTIFIERS = `{
  "access_key" : "022QF06E7MXBSH9DHM02",
  "secret_key" : "kWcrlUX5JEDGM/LtmEENI/aVmYvHNif5zB+d9+ct"
}`

var LOCAL_FILE = `routemaster_test.go`

//TODO: Make this test executable when offline
func TestRawRemoteUrl(t *testing.T) {
  output := ReadRemoteBody("http://example.iana.org/")
  if strings.Contains(output, "Example Domain") != true {
    t.Fatal("Webpage did not return", output)
  }
}

//TODO: Make this test executable when offline (might not be possible)
func TestFetchWanIP(t *testing.T) {
  ipRegex, _ := regexp.Compile(`\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b`)
  ip := fetchWanIP
  if ipRegex.MatchString(string(ip())) {
     t.Fatal("Invalid IP address", ip)
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

func TestReadLocalFile(t *testing.T) {
  output := ReadLocalFile(LOCAL_FILE)
  if len(output) < 0 {
    t.Fatalf("Unable to read contents of file")
  }
}
