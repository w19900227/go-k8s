package services

// import (
// 	"fmt"
// 	"runtime"
// 	"time"
// )
import (
	bs "./baseService"
)

type ImageSearchService struct {
	bs.BaseService
}

// func (this *BaseService) init() {

	// this.baseurl = "http://192.168.6.13:8080/api/v1"
	// this.namespace = "/namespaces/default/"
	// this.hred_ip = "http://192.168.6.14:8090"
// }

func (this *ImageSearchService) GetTest() string {
	return this.GetBaseUrl()
}

func (this *ImageSearchService) GetService() string {
	return this.GetBaseUrl()
}

func (this *ImageSearchService) GetCluster() string {
	return this.GetBaseUrl()
}

func (this *ImageSearchService) GetContainer() string {
	return this.GetBaseUrl()
}

func (this *ImageSearchService) GetMachine() string {
	return this.GetBaseUrl()
}
