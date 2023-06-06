package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Backup struct {
	Version          string        `json:"version"`
	NotificationList []interface{} `json:"notificationList"`
	MonitorList      []MonitorList `json:"monitorList"`
}

type MonitorList struct {
	ID                       int64              `json:"id"`
	Name                     string             `json:"name"`
	Description              interface{}        `json:"description"`
	URL                      string             `json:"url"`
	Method                   string             `json:"method"`
	Hostname                 interface{}        `json:"hostname"`
	Port                     interface{}        `json:"port"`
	Maxretries               int64              `json:"maxretries"`
	Weight                   int64              `json:"weight"`
	Active                   int64              `json:"active"`
	Type                     string             `json:"type"`
	Interval                 int64              `json:"interval"`
	RetryInterval            int64              `json:"retryInterval"`
	ResendInterval           int64              `json:"resendInterval"`
	Keyword                  interface{}        `json:"keyword"`
	ExpiryNotification       bool               `json:"expiryNotification"`
	IgnoreTLS                bool               `json:"ignoreTls"`
	UpsideDown               bool               `json:"upsideDown"`
	PacketSize               int64              `json:"packetSize"`
	Maxredirects             int64              `json:"maxredirects"`
	AcceptedStatuscodes      []string           `json:"accepted_statuscodes"`
	DNSResolveType           string             `json:"dns_resolve_type"`
	DNSResolveServer         string             `json:"dns_resolve_server"`
	DNSLastResult            interface{}        `json:"dns_last_result"`
	DockerContainer          string             `json:"docker_container"`
	DockerHost               interface{}        `json:"docker_host"`
	ProxyID                  interface{}        `json:"proxyId"`
	NotificationIDList       NotificationIDList `json:"notificationIDList"`
	Tags                     []interface{}      `json:"tags"`
	Maintenance              bool               `json:"maintenance"`
	MqttTopic                string             `json:"mqttTopic"`
	MqttSuccessMessage       string             `json:"mqttSuccessMessage"`
	DatabaseQuery            interface{}        `json:"databaseQuery"`
	AuthMethod               interface{}        `json:"authMethod"`
	GrpcURL                  interface{}        `json:"grpcUrl"`
	GrpcProtobuf             interface{}        `json:"grpcProtobuf"`
	GrpcMethod               interface{}        `json:"grpcMethod"`
	GrpcServiceName          interface{}        `json:"grpcServiceName"`
	GrpcEnableTLS            bool               `json:"grpcEnableTls"`
	RadiusCalledStationID    interface{}        `json:"radiusCalledStationId"`
	RadiusCallingStationID   interface{}        `json:"radiusCallingStationId"`
	Game                     interface{}        `json:"game"`
	HTTPBodyEncoding         string             `json:"httpBodyEncoding"`
	Headers                  interface{}        `json:"headers"`
	Body                     interface{}        `json:"body"`
	GrpcBody                 interface{}        `json:"grpcBody"`
	GrpcMetadata             interface{}        `json:"grpcMetadata"`
	BasicAuthUser            interface{}        `json:"basic_auth_user"`
	BasicAuthPass            interface{}        `json:"basic_auth_pass"`
	PushToken                interface{}        `json:"pushToken"`
	DatabaseConnectionString interface{}        `json:"databaseConnectionString"`
	RadiusUsername           interface{}        `json:"radiusUsername"`
	RadiusPassword           interface{}        `json:"radiusPassword"`
	RadiusSecret             interface{}        `json:"radiusSecret"`
	MqttUsername             string             `json:"mqttUsername"`
	MqttPassword             string             `json:"mqttPassword"`
	AuthWorkstation          interface{}        `json:"authWorkstation"`
	AuthDomain               interface{}        `json:"authDomain"`
	TLSCA                    interface{}        `json:"tlsCa"`
	TLSCERT                  interface{}        `json:"tlsCert"`
	TLSKey                   interface{}        `json:"tlsKey"`
	IncludeSensitiveData     bool               `json:"includeSensitiveData"`
}

type NotificationIDList struct {
}

func main() {
	f, err := os.Open("hosts.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	backup := Backup{
		Version:          "1.21.3",
		NotificationList: []interface{}{},
	}
	// Loop over each line in file and print output
	scan := bufio.NewScanner(f)

	scan.Split(bufio.ScanLines)
	i := 1
	for scan.Scan() {
		name := scan.Text()
		name = strings.Split(name, ".")[0]
		// Capitalize first letter
		name = strings.Title(name)
		monitorList := MonitorList{
			ID:                       int64(i),
			Name:                     name,
			Description:              nil,
			URL:                      fmt.Sprintf("https://%s", scan.Text()),
			Method:                   "GET",
			Hostname:                 nil,
			Port:                     nil,
			Maxretries:               1,
			Weight:                   2000,
			Active:                   1,
			Type:                     "http",
			Interval:                 60,
			RetryInterval:            60,
			ResendInterval:           0,
			Keyword:                  nil,
			ExpiryNotification:       false,
			IgnoreTLS:                false,
			UpsideDown:               false,
			PacketSize:               56,
			Maxredirects:             10,
			AcceptedStatuscodes:      []string{"200"},
			DNSResolveType:           "A",
			DNSResolveServer:         "1.1.1.1",
			DNSLastResult:            nil,
			DockerContainer:          "",
			DockerHost:               nil,
			ProxyID:                  nil,
			NotificationIDList:       NotificationIDList{},
			Tags:                     []interface{}{},
			Maintenance:              false,
			MqttTopic:                "",
			MqttSuccessMessage:       "",
			DatabaseQuery:            nil,
			AuthMethod:               nil,
			GrpcURL:                  nil,
			GrpcProtobuf:             nil,
			GrpcMethod:               nil,
			GrpcServiceName:          nil,
			GrpcEnableTLS:            false,
			RadiusCalledStationID:    nil,
			RadiusCallingStationID:   nil,
			Game:                     nil,
			HTTPBodyEncoding:         "json",
			Headers:                  nil,
			Body:                     nil,
			GrpcBody:                 nil,
			GrpcMetadata:             nil,
			BasicAuthUser:            nil,
			BasicAuthPass:            nil,
			PushToken:                nil,
			DatabaseConnectionString: nil,
			RadiusUsername:           nil,
			RadiusPassword:           nil,
			RadiusSecret:             nil,
			MqttUsername:             "",
			MqttPassword:             "",
			AuthWorkstation:          nil,
			AuthDomain:               nil,
			TLSCA:                    nil,
			TLSCERT:                  nil,
			TLSKey:                   nil,
			IncludeSensitiveData:     false,
		}
		backup.MonitorList = append(backup.MonitorList, monitorList)
		i++
	}
	// Convert struct to json
	json, err := json.Marshal(backup)
	if err != nil {
		log.Fatal(err)
	}
	// Write json to file
	err = ioutil.WriteFile("backup.json", json, 0644)

}
