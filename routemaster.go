package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"route53"
)

var CHECKIP_URL = "http://whatismyip.herokuapp.com/"

type AwsAccessIdentifier struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}

func parseAccessIdentifierJSON(data []byte) (route53.AccessIdentifiers, error) {
	var accessIdentifier AwsAccessIdentifier
	err := json.Unmarshal(data, &accessIdentifier)
	return route53.AccessIdentifiers{
		AccessKey: accessIdentifier.AccessKey,
		SecretKey: accessIdentifier.SecretKey,
	}, err
}

func readLocalFile(location string) (c []byte, err error) {
	c, err = ioutil.ReadFile(location)
	if err != nil {
		log.Fatal("ReadLocalFile Error:", c, err)
	}
	return
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

//TODO add test
func findZone(zones route53.HostedZones) (hz route53.HostedZone) {
	for i := range zones.HostedZones {
		if zones.HostedZones[i].Name == *hosted_zone {
			return zones.HostedZones[i]
		}
	}
	return hz
}

//TODO add test
func findRecord(records route53.ResourceRecordSets) (rrs route53.ResourceRecordSet) {
	for i := range records.ResourceRecordSets {
		if records.ResourceRecordSets[i].Name == *subdomain+"."+*hosted_zone {
			return records.ResourceRecordSets[i]
		}
	}
	return rrs
}

func updateRecord(zone route53.HostedZone, aws route53.AccessIdentifiers, action string, name string, value string) {
	var create = route53.ChangeResourceRecordSetsRequest{
		ZoneID:  zone.HostedZoneId(),
		Comment: "",
		Changes: []route53.Change{
			{
				Action: action,
				Name:   name,
				Type:   "A",
				TTL:    300,
				Value:  value,
			},
		},
	}

	r, err := create.Create(aws)

	if err != nil {
		log.Fatal("Update record failed:", r, err)
	}
}

var aws_secrets *string = flag.String("secrets-file", "", "/path_to/.your_aws_secrets")
var hosted_zone *string = flag.String("hosted-zone", "", "[your hosted zone]")
var subdomain *string = flag.String("subdomain", "", "[your subdomain]")

func main() {
	flag.Parse()
	wanIP, err := getWanIP(CHECKIP_URL)

	if err != nil {
		log.Fatal("Failed to fetch current IP:", err)
	}

	fmt.Println("IP is " + wanIP.String())
	contents, _ := readLocalFile(*aws_secrets)
	aws, _ := parseAccessIdentifierJSON(contents)
	zone := findZone(aws.Zones())
	resourceRecordSets, err := zone.ResourceRecordSets(aws)

	if err != nil {
		log.Fatal("Resource Record Sets Invalid:", resourceRecordSets, err)
	}

	record := findRecord(resourceRecordSets)

	if record.Name == "" {
		updateRecord(zone, aws, "CREATE", *subdomain+"."+*hosted_zone, wanIP.String())
		log.Fatal("A record with name " + *subdomain + " was not found, created")
	}

	fmt.Println("IP was " + record.Value[0])

	if record.Value[0] == wanIP.String() {
		log.Fatal("Nothing to do")
	}

	fmt.Println("Updating IP with Route53")
	updateRecord(zone, aws, "DELETE", *subdomain+"."+*hosted_zone, record.Value[0])
	updateRecord(zone, aws, "CREATE", *subdomain+"."+*hosted_zone, wanIP.String())
	fmt.Println("Done")
}
