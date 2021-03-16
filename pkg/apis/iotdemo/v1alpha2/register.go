/*

 Copyright 2019 Advantech.
 Author: jin.xin@advaantech.com.cn.

*/
package v1alpha2

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"
	"iotdemo/pkg/apiserver/iotdemo"
	"iotdemo/pkg/apiserver/runtime"
	"iotdemo/pkg/utils/syslog"
	"net/http"
)

const (
	RespOK    = "ok"
)

func AddWebService() error {
	ws := runtime.NewWebService()

	//get status
	ws.Route(ws.GET("/status").To(iotdemo.GetStatus).
		Doc("Get the iotdemo server status").
		Returns(http.StatusOK, RespOK, iotdemo.OutputsResult{}))

	//get sso status
	ws.Route(ws.GET("/ssostatus").To(iotdemo.GetAuthStatus).
		Filter(runtime.AdvjwtAuthentication).
		Doc("Get sso status").
		Returns(http.StatusOK, RespOK, iotdemo.OutputsResult{}))

	LOG.Info("add rest router end")
	restful.DefaultContainer.Add(ws)

	config := restfulspec.Config{
		WebServices: restful.RegisteredWebServices(), // you control what services are visible
		APIPath:     "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}
	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(config))
	LOG.Info("Get the API using http://ip:port/apidocs.json")

	//path := os.Getenv("GOPATH") + "/src/iotdemo/docs/swagger"
	path := "/iotdemo/app/swagger"
	http.Handle("/apidocs/", http.StripPrefix("/apidocs/", http.FileServer(http.Dir(path))))
	http.ListenAndServe(":8080", nil)
	return nil
}

// refer to  https://github.com/emicklei/go-restful/blob/master/examples/restful-user-resource.go
func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "iotdemo",
			Description: "iot demo",
			Contact: &spec.ContactInfo{
				Name:  "advantech",
				Email: "jin.xin@advantech.com.cn",
				URL:   "http://www.advantech.com.cn",
			},
			License: &spec.License{
				Name: "MIT",
				URL:  "http://mit.org",
			},
			Version: "1.0.0",
		},
	}
	swo.Tags = []spec.Tag{spec.Tag{TagProps: spec.TagProps{
		Name:        "iotdemo",
		Description: "iot demo"}}}
}
