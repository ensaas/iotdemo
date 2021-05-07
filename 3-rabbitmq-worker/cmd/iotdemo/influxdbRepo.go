/*

 Copyright 2021 Advantech
 Author: chienhsiang.chen@advantech.com.tw

*/
package main

import (
	"log"
	"time"

	"github.com/influxdata/influxdb/client/v2"
)

type InfluxdbRepo struct {
	Address  string
	Username string
	Password string
	Database string
	conn     client.Client
}

func (p *InfluxdbRepo) Connect() {
	var err error
	p.conn, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     p.Address,
		Username: p.Username,
		Password: p.Password,
	})

	failOnError(err, "Connect to Influx DB failed.")
}

func (p *InfluxdbRepo) Insert(data Data) {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  p.Database,
		Precision: "ms",
	})

	failOnError(err, "Create Batch Point failed.")

	tags := map[string]string{"id": data.Id}

	fields := map[string]interface{}{
		"value": data.Value,
	}

	pt, err := client.NewPoint("data", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	if err := p.conn.Write(bp); err != nil {
		log.Fatal(err)
	}
}
