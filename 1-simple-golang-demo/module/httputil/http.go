/*

 Copyright 2019 Advantech.
 Author: jin.xin@advaantech.com.cn.

*/
package httputil

import (
	"io/ioutil"
	"iotdemo/pkg/constants"
	"net/http"
	"strings"
)

func HttpPost(url string, headers map[string]string, jbody string) (string, int) {
	client := &http.Client{}
	//req_new := bytes.NewBuffer([]byte(jbody))
	req_new := strings.NewReader(jbody)
	request, _ := http.NewRequest("POST", url, req_new)
	if(headers != nil) {
		for header := range headers {
			request.Header.Add(header, headers[header])
		}
	}
	response, err := client.Do(request)
	//fmt.Printf("Post resp:%#v\n", response)
	if err != nil {
		return "", constants.INTERNAL_SERVER_ERROR
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		//fmt.Println(string(body))
		return string(body), response.StatusCode
	} else {
		body, _ := ioutil.ReadAll(response.Body)
		//LOG.Info("Post ErrMsg: ", string(body))
		return string(body), response.StatusCode
	}
}


func HttpGet(url string, headers map[string]string, jbody string) (string, int) {
	client := &http.Client{}
	//req_new := bytes.NewBuffer([]byte(jbody))
	req_new := strings.NewReader(jbody)
	request, _ := http.NewRequest("GET", url, req_new)
	if(headers != nil) {
		for header := range headers {
			request.Header.Add(header, headers[header])
		}
	}
	response, err := client.Do(request)
	//fmt.Printf("Get resp:%#v\n", response)
	if err != nil {
		return "", constants.INTERNAL_SERVER_ERROR
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		//fmt.Println(string(body))
		return string(body), response.StatusCode
	} else {
		body, _ := ioutil.ReadAll(response.Body)
		//LOG.Info("Get ErrMsg: ", string(body))
		return string(body), response.StatusCode
	}
}


func HttpPut(url string, headers map[string]string, jbody string) (string, int) {
	client := &http.Client{}
	//req_new := bytes.NewBuffer([]byte(jbody))
	req_new := strings.NewReader(jbody)
	request, _ := http.NewRequest("PUT", url, req_new)
	if(headers != nil) {
		for header := range headers {
			request.Header.Add(header, headers[header])
		}
	}
	response, err := client.Do(request)
	//fmt.Printf("Put resp:%#v\n", response)
	if err != nil {
		return "", constants.INTERNAL_SERVER_ERROR
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		//fmt.Println("http resp : "+string(body))
		return string(body), response.StatusCode
	} else {
		body, _ := ioutil.ReadAll(response.Body)
		//LOG.Info("Put ErrMsg: ", string(body))
		return string(body), response.StatusCode
	}
}

func HttpDelete(url string, headers map[string]string, jbody string) (string, int) {
	client := &http.Client{}
	//req_new := bytes.NewBuffer([]byte(jbody))
	req_new := strings.NewReader(jbody)
	request, _ := http.NewRequest("DELETE", url, req_new)
	if(headers != nil) {
		for header := range headers {
			request.Header.Add(header, headers[header])
		}
	}
	response, err := client.Do(request)
	//fmt.Printf("Delete resp:%#v\n", response)
	if err != nil {
		return "", constants.INTERNAL_SERVER_ERROR
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		//fmt.Println(string(body))
		return string(body), response.StatusCode
	} else {
		body, _ := ioutil.ReadAll(response.Body)
		//LOG.Info("Delete ErrMsg: ", string(body))
		return string(body), response.StatusCode
	}
}

