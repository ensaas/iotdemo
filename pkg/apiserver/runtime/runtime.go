/*

 Copyright 2019 Advantech.
 Author: jin.xin@advaantech.com.cn.

*/
package runtime

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"iotdemo/module/ssov3"
	"iotdemo/pkg/utils/syslog"
	"net/http"
	"strings"
)

const (
	ApiRootPath = "/v1"
	ApiTag = "iot"
)

// container holds all webservice of apiserver
var Container = restful.NewContainer()

type ContainerBuilder []func(c *restful.Container) error

const MimeMergePatchJson = "application/merge-patch+json"
const MimeJsonPatchJson = "application/json-patch+json"

func init() {
	restful.RegisterEntityAccessor(MimeMergePatchJson, restful.NewEntityAccessorJSON(restful.MIME_JSON))
	restful.RegisterEntityAccessor(MimeJsonPatchJson, restful.NewEntityAccessorJSON(restful.MIME_JSON))
}

func NewWebService() *restful.WebService {
	webservice := new(restful.WebService)
	webservice.Path(ApiRootPath + "/" + ApiTag).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	//webservice.Filter(jwtAuthentication)
	return webservice
}

func jwtAuthentication(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	tokenHeader := req.HeaderParameter("Authorization")
	//LOG.Info("token is:", tokenHeader)
	if tokenHeader == "" {
		resp.WriteErrorString(http.StatusForbidden, "Not Authorized")
		return
	}

	splitted := strings.Split(tokenHeader, " ")
	if len(splitted) != 2 {
		resp.WriteErrorString(http.StatusForbidden, "Not Authorized")
		return
	}

	chain.ProcessFilter(req, resp)
}

func AdvjwtAuthentication(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	var token string
	token = req.HeaderParameter("Authorization")
	if token == "" {
		cookie, err := req.Request.Cookie("EIToken")
		if err != nil {
			LOG.Error("failed to get EIToken.")
			resp.WriteErrorString(http.StatusUnauthorized, "Not Authorized")
			return
		}
		token = cookie.Value
		if token == "" {
			LOG.Error("token is empty.")
			resp.WriteErrorString(http.StatusForbidden, "Not Authorized")
			return
		}
	}
	info, err := ssov3.GetUsersMe(token)
	if err != nil {
		LOG.Error("get user info fail :", err)
		resp.WriteError(http.StatusUnauthorized, err)
		return
	}
	fmt.Printf("user name: %s\n", info.Username)
	//TODO permission check
	//
	chain.ProcessFilter(req, resp)
}