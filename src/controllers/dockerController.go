package controllers

import (
    "github.com/emicklei/go-restful"
    "services"
    "io/ioutil"
)
import (
    "fmt"
    // "encoding/json"
    // "github.com/fsouza/go-dockerclient"
)

type DockerController struct {

}

func (this *DockerController) SearchImages(request *restful.Request, response *restful.Response) {
    fmt.Println("[Post] Docker search")

    body, err := ioutil.ReadAll(request.Request.Body)

    if err != nil {
        response.WriteEntity("fail")
        return 
    }    

    var images struct {
        Status string `json:"status"`
        Errno string `json:"errno"`
        Errmsg string `json:"errmsg"`
        Data interface{} `json:"data"`
    }

    var docker services.DockerService
    images.Data = docker.SearchImage(body)
    // SearchImage
    images.Status = "ok"


    response.WriteEntity(images)
    return
}

