/*

 Copyright 2019 Advantech.
 Author: jin.xin@advaantech.com.cn.

*/
package main

import (
	"iotdemo/pkg/apis/iotdemo/v1alpha2"
	"iotdemo/pkg/utils/syslog"
)

func main() {
	LOG.Info("iotdemo start...")
	// start rest api server
	v1alpha2.AddWebService()
	LOG.Fatal("iotdemo terminated")
}
