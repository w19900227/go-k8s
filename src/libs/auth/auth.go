package auth

import (
    "github.com/emicklei/go-restful"
    "github.com/garyburd/redigo/redis"
    "encoding/json"
    "io/ioutil"
    "encoding/base64"
    "time"
    "log"
    "strings"
    "crypto/md5"
    "fmt"
    // "io"
    // "net/http"
)

const (
    redis_path = "localhost:6379"
    // redis_path = "golangredis:6379"
)


type Auth struct {}

func (this *Auth) TokenRouter() *restful.WebService {
    service := new(restful.WebService)
    service.
        Path("/token").
        Consumes(restful.MIME_JSON, restful.MIME_JSON).
        Produces(restful.MIME_JSON, restful.MIME_JSON)

    service.Route(service.GET("/{token}").To(checkToken))
    service.Route(service.GET("/time").To(checkTimes))
 
    return service
}

func (this *Auth) UserRouter() *restful.WebService {
    service := new(restful.WebService)
    service.
        Path("/user").
        Consumes(restful.MIME_JSON, restful.MIME_JSON).
        Produces(restful.MIME_JSON, restful.MIME_JSON)

    var user User
    service.Route(service.GET("/{token}").To(user.get))
    service.Route(service.GET("").To(user.getList))
    service.Route(service.POST("").To(user.create))
    service.Route(service.POST("/list").To(user.createList))
    service.Route(service.POST("/login").To(user.login))
    service.Route(service.PUT("/{token}").To(user.update))
    service.Route(service.DELETE("/{token}").To(user.delete))
 
    return service
}

type User struct {
}
type resp struct {
    Status string `json:"status"`
    Erron  string `json:"erron"`
    Errmsg string`json:"errmsg"`
    Data   interface{} `json:"data"`
}

type User_format struct {
    Email string `json:"email"`
    Password  string `json:"password"`
    Start_time time.Time `json:"start_time"`
    Token string `json:"token"`
    Namespace string `json:"namespace,omitempty"`
}

func NewPool() *redis.Pool {
    return &redis.Pool{
                MaxIdle: 80,
                MaxActive: 12000, // max number of connections
                Dial: func() (redis.Conn, error) {
                        c, err := redis.Dial("tcp", redis_path)
                        if err != nil {
                                panic(err.Error())
                        }
                        return c, err
                },
        } 
}

func GetService(key string) ([]byte, error) {
    pool := NewPool()
    c := pool.Get()

    result, err := redis.Bytes(c.Do("GET", key))
    return result, err
}

func GetUser(token string) ([]byte, error) {
    key := "user_list:" + token

    result, err := GetService(key)
    return result, err
}

func (this *User) get(request *restful.Request, response *restful.Response) {
    var r resp

    token := request.PathParameter("token")
    key := "user_list:" + token

    result, err := GetService(key)

    if result == nil || err != nil {
        r.Status = "fail"
        r.Errmsg = "no user"
        response.WriteEntity(r)
        return
    }

    var user User_format
    json.Unmarshal(result, &user)

    r.Status = "ok"
    r.Data = user
    response.WriteEntity(r)
    return 
}

func (this *User) getList(request *restful.Request, response *restful.Response) {
    var r resp
    pool := NewPool()
    c := pool.Get()

    keys := "user_list:*"
    var user User_format
    var user_array []User_format

    var list_user []string
    result, err := redis.Values(c.Do("KEYS", keys))

    if result == nil || err != nil {
        r.Status = "fail"
        r.Errmsg = "no user"
        response.WriteEntity(r)
        return
    }

    redis.ScanSlice(result, &list_user)
    for _, value := range list_user {
        data, _ := GetService(value)
        json.Unmarshal(data, &user)
        user_array = append(user_array, user)
    }

    r.Status = "ok"
    r.Data = user_array
    response.WriteEntity(r)
    return 
}

func (this *User) create(request *restful.Request, response *restful.Response) {
    var r resp
    var user User_format

    body, err := ioutil.ReadAll(request.Request.Body)
    if err != nil {
        r.Status = "fail"
        r.Errmsg = "Request Body error"
        response.WriteEntity(r)
        return
    }

    json.Unmarshal(body, &user)
    user.Password = encryption(user.Password)
    user.Token = encryption(user.Email)
    user.Start_time = user.Start_time

    str := []byte(user.Email)
    bs := md5.Sum(str)
    user.Namespace = fmt.Sprintf("%x", bs)

    value, _ := json.Marshal(user)

    pool := NewPool()
    c := pool.Get()

    key := "user_list:" + user.Token

    result, err := c.Do("SET", key, value)

    if result == nil || err != nil {
        r.Status = "fail"
        r.Errmsg = "no create user"
        response.WriteEntity(r)
        return
    }

    r.Status = "ok"
    response.WriteEntity(r)
    return
}

func (this *User) createList(request *restful.Request, response *restful.Response) {
    var r resp
    var user []User_format

    body, err := ioutil.ReadAll(request.Request.Body)
    if err != nil {
        r.Status = "fail"
        r.Errmsg = "Request Body error"
        response.WriteEntity(r)
        return
    }

    json.Unmarshal(body, &user)
    
    for index := range user {
        user[index].Password = encryption(user[index].Password)
        user[index].Token = encryption(user[index].Email)
        user[index].Start_time = user[index].Start_time

        str := []byte(user[index].Email)
        bs := md5.Sum(str)
        user[index].Namespace = fmt.Sprintf("%x", bs)
        
        value, _ := json.Marshal(user[index])
        pool := NewPool()
        c := pool.Get()

        key := "user_list:" + user[index].Token

        result, err := c.Do("SET", key, value)

        if result == nil || err != nil {
            r.Status = "fail"
            r.Errmsg = "no create user"
            response.WriteEntity(r)
            return
        }
    }

    r.Status = "ok"
    response.WriteEntity(r)
    return
}

func (this *User) update(request *restful.Request, response *restful.Response) {
    var r resp
    var user User_format

    token := request.PathParameter("token")
    body, _ := ioutil.ReadAll(request.Request.Body)
    json.Unmarshal(body, &user)
    id, err := decryption(token)
    if err != nil {
        r.Status = "fail"
        r.Errmsg = "token Decryption error"
        response.WriteEntity(r)
        return
    }
    
    user.Email = id.(string)
    user.Password = encryption(user.Password)
    user.Token = token
    user.Start_time = user.Start_time

    str := []byte(id.(string))
    bs := md5.Sum(str)
    user.Namespace = fmt.Sprintf("%x", bs)
    
    value, _ := json.Marshal(user)

    pool := NewPool()
    c := pool.Get()
    key := "user_list:" + token

    result, err := c.Do("SET", key, value)

    if result == nil || err != nil {
        r.Status = "fail"
        r.Errmsg = "no user"
        response.WriteEntity(r)
        return
    }

    r.Status = "ok"
    response.WriteEntity(r)
    return
}

func (this *User) delete(request *restful.Request, response *restful.Response) {
    var r resp
    token := request.PathParameter("token")

    pool := NewPool()
    c := pool.Get()
    key := "user_list:" + token

    result, err := c.Do("DEL", key)

    if result == nil || err != nil {
        r.Status = "fail"
        r.Errmsg = "no user"
        response.WriteEntity(r)
        return
    }

    r.Status = "ok"
    response.WriteEntity(r)
    return 
}

func (this *User) login(request *restful.Request, response *restful.Response) {
    var r resp
    var data User_format
    body, _ := ioutil.ReadAll(request.Request.Body)
    json.Unmarshal(body, &data)

    pool := NewPool()
    c := pool.Get()

    key := "user_list:" + encryption(data.Email)
    result, err := redis.Bytes(c.Do("GET", key))

    if result == nil || err != nil {
        r.Status = "fail"
        r.Errmsg = "no user"
        response.WriteEntity(r)
        return
    }

    var user User_format
    json.Unmarshal(result, &user)

    pwd := encryption(data.Password)

    if user.Password != pwd {
        r.Status = "fail"
        r.Errmsg = "account or password is error"
        response.WriteEntity(r)
        return
    }

    resp := CheckToken(user.Token)

    if resp.Status != "ok" {
        response.WriteEntity(resp)
        return
    }

    r.Status = "ok"
    r.Data = user.Token
    response.WriteEntity(r)
    return 
}

func checkToken(request *restful.Request, response *restful.Response) {
    token := request.PathParameter("token")

    r := CheckToken(token)
    response.WriteEntity(r)
    return
}

func CheckToken(token string) (resp) {
    var r resp
    
    pool := NewPool()
    c := pool.Get()

    key := "user_list:" + token
    result, err := redis.Bytes(c.Do("GET", key))

    if result == nil || err != nil {
        r.Status = "fail"
        r.Errmsg = "no token"
        return r
    }

    var user User_format
    json.Unmarshal(result, &user)

    if user.Token != token {
        r.Status = "fail"
        r.Errmsg = "token auth issue"
        return r
    }

    layout := "2006-01-02 15:04:05"
    now_time := time.Now()
    now_time = now_time.Add(time.Hour * 8)
    start_time := user.Start_time
    last_time := user.Start_time.Add(time.Hour * 24)

    now_times := now_time.Format(layout)
    now_time_parse, err := time.Parse(layout, now_times)

    start_times := start_time.Format(layout)
    start_time_parse, err := time.Parse(layout, start_times)
    
    last_times := last_time.Format(layout)
    last_time_parse, err := time.Parse(layout, last_times)

    log.Printf("Token:%s, Now:%s, Start:%s, End:%s\n", token, now_time_parse, start_time_parse, last_time_parse)

    if start_time_parse.After(now_time_parse) {
        r.Status = "fail"
        r.Errmsg = "no start"
        return r
    }
    if last_time_parse.Before(now_time_parse) {
        r.Status = "fail"
        r.Errmsg = "expire"
        return r
    }

    r.Status = "ok"
    return r
}

func checkTimes(request *restful.Request, response *restful.Response) {
    var r resp

    now_time := time.Now()

    r.Status = "ok"
    r.Data = now_time
    response.WriteEntity(r)
    return 
}

func encryption(str string) string {
    data := []byte(str)
    ret := base64.StdEncoding.EncodeToString(data)
    return ret
}

func decryption(str string) (interface{}, error) {
    data, err := base64.StdEncoding.DecodeString(str)
    if err != nil {
        return nil, err
    }
    return string(data), nil
}

func set_namespace(str string) string {
    s := strings.NewReplacer("@", "", ".", "", "/", "")
    return s.Replace(str)
}


func Test(result interface{}) {
    js, _ := json.Marshal(result)
    fmt.Println(string(js))
}
