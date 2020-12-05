package main

import (
	client "github.com/influxdata/influxdb1-client"
	"log"
	"time"
)

func writeSomeData() error {
	influxDBClient, err := createInfluxDBClient()
	if err != nil {
		log.Fatal("Error: ", err.Error())
		return err
	}

	// Point is one piece of data that is being written in the influxDB.
	// For example: In order to insert one row, you can create one point and that can be inserted.
	point := client.Point{
		Measurement: "test_measurement",
		Tags:        map[string]string {
			"tag1": "tag1value",
			"tag2": "tag2value",
		},
		Time:        time.Now(),
		Fields: map[string]interface{}{
			"field1": "value1",
			"field2": "value2",
		},
		Precision:   "ns", // ns means nanosecond precision. ms and s are other available precision.
		// Raw:         "",
	}

	// Influx db Go client doesn't has method to insert one single point i.e one row. InfluxDB is designed to insert
	// lots of data. Hence it has method to do insertion of bulk of points.
	// So create builk of points.
	bulkPoints := client.BatchPoints{
		Points:           []client.Point{point},
		Database:         influxDBName,
		//RetentionPolicy:  "",
		//Tags:             nil,
		//Time:             time.Time{},
		//Precision:        "",
		//WriteConsistency: "",
	}

	res, err := influxDBClient.Write(bulkPoints)
	if err != nil {
		log.Fatal("Error: %s", err.Error())
		return err
	}

	log.Println("Response: ", res)
	return nil
}

