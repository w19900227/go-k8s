# 整合kubernetes API

主要提供前、後端資料整合傳遞(kubernetes)，`僅對於1.0版本進行測試正常`

### 安裝golang
    $ wget https://storage.googleapis.com/golang/go1.5.1.linux-amd64.tar.gz
    $ tar -C /usr/local -xzf go1.5.1.linux-amd64.tar.gz
    $ export PATH=$PATH:/usr/local/go/bin

### 環境部署
    $ git clone https://github.com/w19900227/go-k8s.git
    $ cd go-k8s
    $ export GOPATH=$PWD

### Testing
尚未開發

### get plug
	$ go get github.com/emicklei/go-restful
	$ go get github.com/fsouza/go-dockerclient
	$ go get github.com/garyburd/redigo/redis


### 參考資料
kubernetes API type format 
* https://github.com/kubernetes/kubernetes/blob/master/pkg/api/types.go