package iotdemo

import (
	"github.com/emicklei/go-restful"
)

type OutputsResult struct {
	Status  string    `json:"status" description:"response status"`
}

func GetStatus(req *restful.Request, resp *restful.Response) {
	//InstInitialHelms()
	var result OutputsResult
	result.Status = "ok"
	_ = resp.WriteAsJson(&result)
	return
}

func GetAuthStatus(req *restful.Request, resp *restful.Response) {
	//InstInitialHelms()
	var result OutputsResult
	result.Status = "ok"
	_ = resp.WriteAsJson(&result)
	return
}