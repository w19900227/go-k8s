package request

import (
    "net/http"
    "io/ioutil"
    "encoding/json"
    "strings"
)
import (
    "fmt"
)

var types string = "application/json"

type Request struct {

}

func (this *Request) ClearRequest() {
}

func (this *Request) Get(url string) []byte {
    response, _ := http.Get(url)
    datas, _ := ioutil.ReadAll(response.Body)
    return datas
}

func (this *Request) Post(url string, data interface{}) []byte {
    client := &http.Client{}
    
    buf, _ := json.Marshal(data)
    body := strings.NewReader(string(buf))

    response, _ := http.NewRequest("POST", url, body)
    response.Header.Set("Content-Type", "application/json")
    re, _ := client.Do(response)
    defer re.Body.Close()
    datas, _ := ioutil.ReadAll(re.Body)
    return datas    
}

func (this *Request) Put(url string, data interface{}) []byte {
    client := &http.Client{}

    buf, _ := json.Marshal(data)
    body := strings.NewReader(string(buf))

    response, _ := http.NewRequest("PUT", url, body)
    response.Header.Set("Content-Type", "application/json")
    re, _ := client.Do(response)
    defer re.Body.Close()
    datas, _ := ioutil.ReadAll(re.Body)
    return datas    
}


// func (this *Request) Delete(url string, data interface{}) []byte {
func (this *Request) Delete(url string, data interface{}) []byte {
    if data != nil {
        return this.DeleteData(url, data)
    }
    client := &http.Client{}

    response, _ := http.NewRequest("DELETE", url, nil)
    response.Header.Set("Content-Type", "application/json")
    re, _ := client.Do(response)
    defer re.Body.Close()
    datas, _ := ioutil.ReadAll(re.Body)
    return datas  
}
func (this *Request) DeleteData(url string, data interface{}) []byte {
    client := &http.Client{}

    buf, _ := json.Marshal(data)
    body := strings.NewReader(string(buf))

    response, _ := http.NewRequest("DELETE", url, body)
    response.Header.Set("Content-Type", "application/json")
    re, _ := client.Do(response)
    defer re.Body.Close()
    datas, _ := ioutil.ReadAll(re.Body)
    return datas    
}
func tt() {
    fmt.Println("req test")
}


