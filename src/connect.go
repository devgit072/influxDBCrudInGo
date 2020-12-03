package main

import (
	"fmt"
	"log"
	"net/url"

	client "github.com/influxdata/influxdb1-client"
)

const (
	influxDBHost = "localhost"
	influxDBPort = 8086
	influxDBUserName = ""
	influxDBPassword = ""
)

const influxDBName = "test_db"

func createInfluxDBClient() (*client.Client, error) {
	host, err := url.Parse(fmt.Sprintf("http://%s:%d", influxDBHost, influxDBPort))
	if err != nil {
		log.Fatal("Error: %s", err.Error())
		return nil, err
	}

	conf := client.Config{
		URL: *host,
		Username: influxDBUserName,
		Password: influxDBPassword,
	}

	influxDBClient, err := client.NewClient(conf)
	if err != nil {
		log.Fatal("Error: %s", err.Error())
		return nil, err
	}

	if _,_, err := influxDBClient.Ping(); err != nil {
		log.Fatal("Error while pinging. Error:", err.Error())
	}
	return influxDBClient, nil
}