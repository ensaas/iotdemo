package iotdemo

import (
	"github.com/emicklei/go-restful"
	"net/http"
)

type OutputsResult struct {
	Status  int    `json:"status" description:"response status"`
	Error   string `json:"error,omitempty" description:"debug information"`
	Outputs string `json:"content,omitempty" description:"outputs string"`
}

func GetStatus(req *restful.Request, resp *restful.Response) {
	//InstInitialHelms()
	var result OutputsResult
	result.Outputs = "ok"
	result.Status = http.StatusOK
	_ = resp.WriteAsJson(&result)
	return
}

func GetAuthStatus(req *restful.Request, resp *restful.Response) {
	//InstInitialHelms()
	var result OutputsResult
	result.Outputs = "ok"
	result.Status = http.StatusOK
	_ = resp.WriteAsJson(&result)
	return
}