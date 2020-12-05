package main

import (
	"fmt"
	client "github.com/influxdata/influxdb1-client"
	"log"
)

func readData() error {
	query := client.Query{
		Command:         "select time,field1,field2,tag1,tag2 from test_measurement",
		Database:        influxDBName,
		//RetentionPolicy: "",
		//Chunked:         false,
		//ChunkSize:       0,
		//NodeID:          0,
	}
	influxDBClient, err := createInfluxDBClient()
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
		return err
	}
	resp, err := influxDBClient.Query(query)
	if resp.Err != nil {
		log.Fatal("Error: %s", resp.Err.Error())
		return resp.Err
	}

	result := resp.Results[0]
	// Above result variable consists of three fields: err, message and series.Series is list of rows results.
	if len(result.Series) == 0 {
		fmt.Println("No result found")
		return nil
	}

	series := result.Series[0]
	var rowsResult = make([][]interface{}, len(series.Values)) // rowsResult is list of string values of each and every rows.
	for i,d := range series.Values {
		rowsResult[i] = d
	}
	// Okay. We have stored the results, now it is time to display the row results.
	columns := series.Columns

	// Print the name of the columns.
	for _,c := range columns {
		fmt.Print(c + "   ")
	}
	fmt.Println("\n============================================")

	for _, eachRow := range rowsResult {
		for _, value := range eachRow {
			fmt.Print(fmt.Sprintf("%s", value) + "   ")
		}
		fmt.Println("")
	}
	return nil
}

/*
It prints something like this:

time   field1   field2   tag1   tag2
============================================
2020-12-05T02:28:15.658163229Z   value1   value2   tag1value   tag2value
2020-12-05T02:51:11.64192629Z   value1   value2   tag1value   tag2value
2020-12-05T02:52:10.528700583Z   value1   value2   tag1value   tag2value
2020-12-05T02:52:27.481367062Z   value1   value2   tag1value   tag2value
*/
