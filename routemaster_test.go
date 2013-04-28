package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var ACCESS_IDENTIFIERS = `{
  "access_key" : "022QF06E7MXBSH9DHM02",
  "secret_key" : "kWcrlUX5JEDGM/LtmEENI/aVmYvHNif5zB+d9+ct"
}`

var LOCAL_FILE = `routemaster_test.go`

type emptyHandler struct{}

func (h *emptyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Test-Header", "Testing")
	io.WriteString(w, "127.0.0.1")
}

func TestRawRemoteUrl(t *testing.T) {
	handler := &emptyHandler{}
	server := httptest.NewServer(handler)
	resp, err := getBody(server.URL)

	if err != nil {
		t.Fatal("Error:", err)
	}

	if strings.Contains(string(resp), "127.0.0.1") != true {
		t.Fatal("Webpage did not return", resp)
	}
}

func TestGetWanIP(t *testing.T) {
	handler := &emptyHandler{}
	server := httptest.NewServer(handler)
	ip, err := getWanIP(server.URL)

	if err != nil {
		t.Fatal("Error:", err)
	}

	if ip.String() != "127.0.0.1" {
		t.Fatal("Invalid IP address", ip)
	}
}

func TestParseAccessIdentifiers(t *testing.T) {
	accessIdentifier, error := parseAccessIdentifierJSON([]byte(ACCESS_IDENTIFIERS))

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
	output, err := readLocalFile(LOCAL_FILE)

	if err != nil {
		t.Fatal("Error:", err)
	}

	if len(output) < 0 {
		t.Fatalf("Unable to read contents of file")
	}
}

func TestFindZone(t *testing.T) {

}
