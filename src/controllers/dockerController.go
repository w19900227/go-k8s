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
        Status string `json:"status,inline"`
        Errno string `json:"errno,inline"`
        Errmsg string `json:"errmsg,inline"`
        Data interface{} `json:"data,inline"`
    }

    var docker services.DockerService
    images.Data = docker.SearchImage(body)
    // SearchImage
    images.Status = "ok"


    response.WriteEntity(images)
    return
}

func (this *DockerController) Images(request *restful.Request, response *restful.Response) {
    fmt.Println("[Get] Docker Images")

    var images struct {
        Status string `json:"status,inline"`
        Errno string `json:"errno,inline"`
        Errmsg string `json:"errmsg,inline"`
        Data interface{} `json:"data,inline"`
    }

    var d []interface{}
    d = append(d, map[string]string{"cael/wordpress":"Wordpress"})
    d = append(d, map[string]string{"chejenyeh/demo-order:tc":"order system"})
    d = append(d, map[string]string{"kubernetes/example-guestbook-php-redis:v2":"Guestbook - step 1"})
    d = append(d, map[string]string{"redis":"Guestbook - step 2"})
    d = append(d, map[string]string{"kubernetes/redis-slave:v2":"Guestbook - step 3"})
    images.Data = d
    images.Status = "ok"

    response.WriteEntity(images)
    return
}

