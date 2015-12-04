package services

import (
	// "fmt"
	"encoding/json"

// 	"runtime"
// 	"time"
)

import (
	bs "services/baseService"
    "github.com/fsouza/go-dockerclient"
	// "models/k8sModel"
	// "models/herdModel"
	// "strings"
	// "strconv"
	// herd_format "models/herdModel/format"

)

type DockerService struct {
	bs.BaseService
	
}

type image_info struct {
    Name string `json:"name"`
    Description string `json:"description"`
    Stars int `json:"stars"`
    Offical bool `json:"offical"`
    Automated bool `json:"automated"`
}
type docker_data_format struct {
	Image string `"json:image"`
}
func (this *DockerService) SearchImage(data []byte) []image_info {
	var _data_format docker_data_format
	json.Unmarshal(data, &_data_format)

	endpoint := "unix:///var/run/docker.sock"
    client, _ := docker.NewClient(endpoint)
    imgs, _ := client.SearchImages(_data_format.Image)
    
    var images []image_info
    for _, img := range imgs {
        var image_info image_info
        image_info.Name = img.Name
        image_info.Description = img.Description
        image_info.Stars = img.StarCount
        image_info.Offical = img.IsOfficial
        image_info.Automated = img.IsAutomated
        images = append(images, image_info)
    }
    return images
}
