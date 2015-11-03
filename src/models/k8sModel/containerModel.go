package k8sModel

import (
	bm "../baseModel"
	"encoding/json"
	"libs/request"
	"./format"
)

type ContainerModel struct {
	bm.BaseModel
	req request.Request
}

// func (this *ContainerModel) init() {

	// this.baseurl = "http://192.168.6.13:8080/api/v1"
	// this.namespace = "/namespaces/default/"
	// this.hred_ip = "http://192.168.6.14:8090"
// }

// func (this *ContainerModel) GetContainer() (format.Pod) {}

//curl -H "Content-Type: application/json" -X GET  http://192.168.12.8:8080/api/v1/namespaces/default/pods
func (this *ContainerModel) GetContainerList() (format.PodList) {
	// body := this.req.Get("http://192.168.15.77:8080/api/v1/namespaces/default/replicationcontrollers")
	body := this.getPagInfo_data()//在伺服器run必須移除此行

    var data format.PodList
	json.Unmarshal(body, &data)
	// fmt.Println(err)
	// fmt.Println(b)
	return data
}

func (this *ContainerModel) getPagInfo_data() []byte {
	body := []byte(`{"kind": "PodList","apiVersion": "v1","metadata": {"selfLink": "/api/v1/namespaces/default/pods","resourceVersion": "1367457"},"items": [{"metadata": {"name": "app-5480l","generateName": "app-","namespace": "default","selfLink": "/api/v1/namespaces/default/pods/app-5480l","uid": "6ca02041-56a6-11e5-9b89-fa163e78675c","resourceVersion": "943738","creationTimestamp": "2015-09-09T03:54:09Z","labels": {"name": "app"},"annotations": {"kubernetes.io/created-by": "{\"kind\":\"SerializedReference\",\"apiVersion\":\"v1\",\"reference\":{\"kind\":\"ReplicationController\",\"namespace\":\"default\",\"name\":\"app\",\"uid\":\"4af54540-5603-11e5-9b89fa163e78675c\",\"apiVersion\":\"v1\",\"resourceVersion\":\"657302\"}}"}},"spec": {"containers": [{"name": "app","image": "192.168.12.7:5000/app","ports": [{"containerPort": 8000,"protocol": "TCP"}],"resources": {},"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"}],"restartPolicy": "Always","dnsPolicy": "ClusterFirst","nodeName": "192.168.12.9"},"status": {"phase": "Running","conditions": [{"type": "Ready","status": "True"}],"hostIP": "192.168.12.9","podIP": "10.1.83.34","startTime": "2015-09-09T03:54:12Z","containerStatuses": [{"name": "app","state": {"running": {"startedAt": "2015-09-09T03:54:11Z"}},"lastState": {},"ready": true,"restartCount": 0,"image": "192.168.12.7:5000/app","imageID": "docker://b69a2ec1517d2bde1eabc8300246ba5cc38116af835c6383b1283013f1883280","containerID": "docker://2aa729385841fb3f98574a76e2106f599851dbd758161b26aea6aba08b6ab6f0"}]}},{"metadata": {"name": "app-o06sd","generateName": "app-","namespace": "default","selfLink": "/api/v1/namespaces/default/pods/app-o06sd","uid": "4b078d68-5603-11e5-9b89-fa163e78675c","resourceVersion": "645685","creationTimestamp": "2015-09-08T08:26:24Z","labels": {"name": "app"},"annotations": {"kubernetes.io/created-by": "{\"kind\":\"SerializedReference\",\"apiVersion\":\"v1\",\"reference\":{\"kind\":\"ReplicationController\",\"namespace\":\"default\",\"name\":\"app\",\"uid\":\"4af54540-5603-11e5-9b89fa163e78675c\",\"apiVersion\":\"v1\",\"resourceVersion\":\"645626\"}}"}},"spec": {"containers": [{"name": "app","image": "192.168.12.7:5000/app","ports": [{"containerPort": 8000,"protocol": "TCP"}],"resources": {},"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"}],"restartPolicy": "Always","dnsPolicy": "ClusterFirst","nodeName": "192.168.12.9"},"status": {"phase": "Running","conditions": [{"type": "Ready","status": "True"}],"hostIP": "192.168.12.9","podIP": "10.1.83.32","startTime": "2015-09-08T08:26:26Z","containerStatuses": [{"name": "app","state": {"running": {"startedAt": "2015-09-08T08:26:26Z"}},"lastState": {},"ready": true,"restartCount": 0,"image": "192.168.12.7:5000/app","imageID": "docker://b69a2ec1517d2bde1eabc8300246ba5cc38116af835c6383b1283013f1883280","containerID": "docker://c117bd35178da6b63ef6f21d30be45507b2ea8a1417721650cd14351ed595479"}]}},{"metadata": {"name": "app-ytus0","generateName": "app-","namespace": "default","selfLink": "/api/v1/namespaces/default/pods/app-ytus0","uid": "6cd0ad4a-56a6-11e5-9b89-fa163e78675c","resourceVersion": "943767","creationTimestamp": "2015-09-09T03:54:09Z","labels": {"name": "app"},"annotations": {"kubernetes.io/created-by": "{\"kind\":\"SerializedReference\",\"apiVersion\":\"v1\",\"reference\":{\"kind\":\"ReplicationController\",\"namespace\":\"default\",\"name\":\"app\",\"uid\":\"4af54540-5603-11e5-9b89fa163e78675c\",\"apiVersion\":\"v1\",\"resourceVersion\":\"943686\"}}"}},"spec": {"containers": [{"name": "app","image": "192.168.12.7:5000/app","ports": [{"containerPort": 8000,"protocol": "TCP"}],"resources": {},"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"}],"restartPolicy": "Always","dnsPolicy": "ClusterFirst","nodeName": "192.168.12.9"},"status": {"phase": "Running","conditions": [{"type": "Ready","status": "True"}],"hostIP": "192.168.12.9","podIP": "10.1.83.35","startTime": "2015-09-09T03:54:12Z","containerStatuses": [{"name": "app","state": {"running": {"startedAt": "2015-09-09T03:54:11Z"}},"lastState": {},"ready": true,"restartCount": 0,"image": "192.168.12.7:5000/app","imageID": "docker://b69a2ec1517d2bde1eabc8300246ba5cc38116af835c6383b1283013f1883280","containerID": "docker://870ab752c60859ba9ee01d4002bb72feaf2ab3f6f76db5800d160f17196c327d"}]}},{"metadata": {"name": "dev-03s84","generateName": "dev-","namespace": "default","selfLink": "/api/v1/namespaces/default/pods/dev-03s84","uid": "47edf449-5603-11e5-9b89-fa163e78675c","resourceVersion": "645635","creationTimestamp": "2015-09-08T08:26:19Z","labels": {"name": "dev"},"annotations": {"kubernetes.io/created-by": "{\"kind\":\"SerializedReference\",\"apiVersion\":\"v1\",\"reference\":{\"kind\":\"ReplicationController\",\"namespace\":\"default\",\"name\":\"dev\",\"uid\":\"47e1772f-5603-11e5-9b89fa163e78675c\",\"apiVersion\":\"v1\",\"resourceVersion\":\"645591\"}}"}},"spec": {"containers": [{"name": "dev","image": "192.168.12.7:5000/dev:p31000","ports": [{"containerPort": 80,"protocol": "TCP"}],"resources": {},"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"}],"restartPolicy": "Always","dnsPolicy": "ClusterFirst","nodeName": "192.168.12.9"},"status": {"phase": "Running","conditions": [{"type": "Ready","status": "True"}],"hostIP": "192.168.12.9","podIP": "10.1.83.31","startTime": "2015-09-08T08:26:21Z","containerStatuses": [{"name": "dev","state": {"running": {"startedAt": "2015-09-08T08:26:21Z"}},"lastState": {},"ready": true,"restartCount": 0,"image": "192.168.12.7:5000/dev:p31000","imageID": "docker://dc2ec8fe6d5d5eba3e34dcfc23ff619c5fc724e8a9b99320973bcc72204c1560","containerID": "docker://86b5c45fea9343892547b9b1b359083208c7aed8c29ea692cd1cf3661a8a80ea"}]}},{"metadata": {"name": "frontend-1n27i","generateName": "frontend-","namespace": "default","selfLink": "/api/v1/namespaces/default/pods/frontend-1n27i","uid": "15defa40-550c-11e5-820f-fa163e78675c","resourceVersion": "244963","creationTimestamp": "2015-09-07T02:56:49Z","labels": {"name": "frontend"},"annotations": {"kubernetes.io/created-by": "{\"kind\":\"SerializedReference\",\"apiVersion\":\"v1\",\"reference\":{\"kind\":\"ReplicationController\",\"namespace\":\"default\",\"name\":\"frontend\",\"uid\":\"15d2e7ac-550c-11e5-820ffa163e78675c\",\"apiVersion\":\"v1\",\"resourceVersion\":\"188536\"}}"}},"spec": {"containers": [{"name": "php-redis","image": "192.168.12.7:5000/kubernetes/example-guestbook-php-redis:v2","ports": [{"containerPort": 80,"protocol": "TCP"}],"resources": {},"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"}],"restartPolicy": "Always","dnsPolicy": "ClusterFirst","nodeName": "192.168.12.9"},"status": {"phase": "Running","conditions": [{"type": "Ready","status": "True"}],"hostIP": "192.168.12.9","podIP": "10.1.83.5","startTime": "2015-09-07T03:02:28Z","containerStatuses": [{"name": "php-redis","state": {"running": {"startedAt": "2015-09-07T06:26:43Z"}},"lastState": {"terminated": {"exitCode": 0,"startedAt": "2015-09-07T05:44:41Z","finishedAt": "2015-09-07T06:26:08Z","containerID": "docker://40c40661def8ed7c40eb8a2c5f93a9b295a3d3ccd6cf87ec35edeb7db618763d"}},"ready": true,"restartCount": 1,"image": "192.168.12.7:5000/kubernetes/example-guestbook-php-redis:v2","imageID": "docker://dac65295b08eabdd483c787bbd765e1b244e4f25ca1474329dc5b125e6a0f692","containerID": "docker://c3dd3e9184fb594cf3a7233cddb5beda444dc92f9b2a001cced0d2f7525ef5dc"}]}},{"metadata": {"name": "frontend-b1dlz","generateName": "frontend-","namespace": "default","selfLink": "/api/v1/namespaces/default/pods/frontend-b1dlz","uid": "276a0cb0-5524-11e5-8634-fa163e78675c","resourceVersion": "244952","creationTimestamp": "2015-09-07T05:49:07Z","labels": {"name": "frontend"},"annotations": {"kubernetes.io/created-by": "{\"kind\":\"SerializedReference\",\"apiVersion\":\"v1\",\"reference\":{\"kind\":\"ReplicationController\",\"namespace\":\"default\",\"name\":\"frontend\",\"uid\":\"15d2e7ac-550c-11e5-820ffa163e78675c\",\"apiVersion\":\"v1\",\"resourceVersion\":\"233666\"}}"}},"spec": {"containers": [{"name": "php-redis","image": "192.168.12.7:5000/kubernetes/example-guestbook-php-redis:v2","ports": [{"containerPort": 80,"protocol": "TCP"}],"resources": {},"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"}],"restartPolicy": "Always","dnsPolicy": "ClusterFirst","nodeName": "192.168.12.9"},"status": {"phase": "Running","conditions": [{"type": "Ready","status": "True"}],"hostIP": "192.168.12.9","podIP": "10.1.83.2","startTime": "2015-09-07T05:49:50Z","containerStatuses": [{"name": "php-redis","state": {"running": {"startedAt": "2015-09-07T06:26:43Z"}},"lastState": {"terminated": {"exitCode": 0,"startedAt": "2015-09-07T05:49:49Z","finishedAt": "2015-09-07T06:26:08Z","containerID": "docker://19e0046a411f65e6767736d48c05598e0e7d06ed23de8304d0455f479e54e624"}},"ready": true,"restartCount": 1,"image": "192.168.12.7:5000/kubernetes/example-guestbook-php-redis:v2","imageID": "docker://dac65295b08eabdd483c787bbd765e1b244e4f25ca1474329dc5b125e6a0f692","containerID": "docker://fb7ac3652823f883aaf0119e3ffcfe17b68987649e3b202da1ae90c731149a8f"}]}},{"metadata": {"name": "frontend-o33fq","generateName": "frontend-","namespace": "default","selfLink": "/api/v1/namespaces/default/pods/frontend-o33fq","uid": "15de6344-550c-11e5-820f-fa163e78675c","resourceVersion": "244957","creationTimestamp": "2015-09-07T02:56:49Z","labels": {"name": "frontend"},"annotations": {"kubernetes.io/created-by": "{\"kind\":\"SerializedReference\",\"apiVersion\":\"v1\",\"reference\":{\"kind\":\"ReplicationController\",\"namespace\":\"default\",\"name\":\"frontend\",\"uid\":\"15d2e7ac-550c-11e5-820ffa163e78675c\",\"apiVersion\":\"v1\",\"resourceVersion\":\"188536\"}}"}},"spec": {"containers": [{"name": "php-redis","image": "192.168.12.7:5000/kubernetes/example-guestbook-php-redis:v2","ports": [{"containerPort": 80,"protocol": "TCP"}],"resources": {},"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"}],"restartPolicy": "Always","dnsPolicy": "ClusterFirst","nodeName": "192.168.12.9"},"status": {"phase": "Running","conditions": [{"type": "Ready","status": "True"}],"hostIP": "192.168.12.9","podIP": "10.1.83.4","startTime": "2015-09-07T03:02:28Z","containerStatuses": [{"name": "php-redis","state": {"running": {"startedAt": "2015-09-07T06:26:43Z"}},"lastState": {"terminated": {"exitCode": 0,"startedAt": "2015-09-07T05:44:41Z","finishedAt": "2015-09-07T06:26:08Z","containerID": "docker://3583f1e182b6770ad5529e23a0ff23345410a0cd863aaccf05adf12a2351e2d5"}},"ready": true,"restartCount": 1,"image": "192.168.12.7:5000/kubernetes/example-guestbook-php-redis:v2","imageID": "docker://dac65295b08eabdd483c787bbd765e1b244e4f25ca1474329dc5b125e6a0f692","containerID": "docker://50bc7ebfaf843de380ca6a40f616dd41266bb8a9000eb0239181594ef73826bd"}]}},{"metadata": {"name": "kube-dns-v3-xtblt","generateName": "kube-dns-v3-","namespace": "default","selfLink": "/api/v1/namespaces/default/pods/kube-dns-v3-xtblt","uid": "278bfd47-5524-11e5-8634-fa163e78675c","resourceVersion": "1367393","creationTimestamp": "2015-09-07T05:49:07Z","labels": {"k8s-app": "kube-dns","kubernetes.io/cluster-service": "true","version": "v3"},"annotations": {"kubernetes.io/created-by": "{\"kind\":\"SerializedReference\",\"apiVersion\":\"v1\",\"reference\":{\"kind\":\"ReplicationController\",\"namespace\":\"default\",\"name\":\"kube-dns-v3\",\"uid\":\"9db5efa5-5507-11e5-820ffa163e78675c\",\"apiVersion\":\"v1\",\"resourceVersion\":\"233650\"}}"}},"spec": {"volumes": [{"name": "dns-token","secret": {"secretName": "token-system-dns"}}],"containers": [{"name": "etcd","image": "192.168.12.7:5000/google_containers/etcd:2.0.9","command": ["/usr/local/bin/etcd","-listen-client-urls","http://127.0.0.1:2379,http://127.0.0.1:4001","-advertise-client-urls","http://127.0.0.1:2379,http://127.0.0.1:4001","-initial-cluster-token","skydns-etcd"],"resources": {},"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"},{"name": "kube2sky","image": "192.168.12.7:5000/google_containers/kube2sky:1.9","args": ["-domain=kubernetes.local","-kubecfg_file=/etc/dns_token/kubeconfig","-v=2","-kube_master_url=https://${KUBERNETES_SERVICE_HOST}:${KUBERNETES_SERVICE_PORT}"],"resources": {},"volumeMounts": [{"name": "dns-token","readOnly": true,"mountPath": "/etc/dns_token"}],"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"},{"name": "skydns","image": "192.168.12.7:5000/google_containers/skydns:2015-03-11-001","args": ["-machines=http://localhost:4001","-addr=0.0.0.0:53","-domain=kubernetes.local.","-verbose=true"],"ports": [{"name": "dns","containerPort": 53,"protocol": "UDP"},{"name": "dns-tcp","containerPort": 53,"protocol": "TCP"}],"resources": {},"livenessProbe": {"exec": {"command": ["/bin/sh","-c","nslookup kubernetes.default.kubernetes.local localhost \u003e/dev/null"]},"initialDelaySeconds": 30,"timeoutSeconds": 5},"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"}],"restartPolicy": "Always","dnsPolicy": "Default","nodeName": "192.168.12.9"},"status": {"phase": "Running","conditions": [{"type": "Ready","status": "True"}],"hostIP": "192.168.12.9","podIP": "10.1.83.8","startTime": "2015-09-07T05:50:03Z","containerStatuses": [{"name": "skydns","state": {"running": {"startedAt": "2015-09-07T06:26:44Z"}},"lastState": {"terminated": {"exitCode": 2,"startedAt": "2015-09-07T05:50:03Z","finishedAt": "2015-09-07T06:23:49Z","containerID": "docker://36aff2f7a7e89415f43ef76ed1e23c7af14d39b02c6697d863706fd88ab7fae1"}},"ready": true,"restartCount": 1,"image": "192.168.12.7:5000/google_containers/skydns:2015-03-11-001","imageID": "docker://791ddf327076e0fd35a1125568a56c05ee1f1dfd7a165c74f4d489d8a5e65ac5","containerID": "docker://4ec888f0ecd0ee869125aab986c81140d310a6cd91042ad67ed402905bb10828"},{"name": "kube2sky","state": {"running": {"startedAt": "2015-09-07T06:26:44Z"}},"lastState": {"terminated": {"exitCode": 2,"startedAt": "2015-09-07T05:50:02Z","finishedAt": "2015-09-07T06:23:49Z","containerID": "docker://62df990afcaf969528d4b0cbf65f602776b200310416bb7841c42bea4ad500f6"}},"ready": true,"restartCount": 1,"image": "192.168.12.7:5000/google_containers/kube2sky:1.9","imageID": "docker://dbcdf588b1ed09a481f7ba0ec2ff15f552035a8b3a18afa2fa70d45164581d97","containerID": "docker://d8885f2d77a8d02f103baebee69a6a69fa0c994ee70c7e3d3216d6ca001c4b2f"},{"name": "etcd","state": {"running": {"startedAt": "2015-09-07T06:26:44Z"}},"lastState": {"terminated": {"exitCode": 0,"startedAt": "2015-09-07T05:49:56Z","finishedAt": "2015-09-07T06:23:49Z","containerID": "docker://aed7ed760868eae72094bf042f58bda7c4369bc348e7f1d257596eb31a59fc6c"}},"ready": true,"restartCount": 1,"image": "192.168.12.7:5000/google_containers/etcd:2.0.9","imageID": "docker://b6b9a86dc06aa1361357ca1b105feba961f6a4145adca6c54e142c0be0fe87b0","containerID": "docker://f126bcf61f39e0f183a3b1eb885f5416dbad57ecbc751819249dfd023f0d489c"}]}},{"metadata": {"name": "redis-master-4aqud","generateName": "redis-master-","namespace": "default","selfLink": "/api/v1/namespaces/default/pods/redis-master-4aqud","uid": "1ccbc6da-550c-11e5-820f-fa163e78675c","resourceVersion": "244973","creationTimestamp": "2015-09-07T02:57:01Z","labels": {"name": "redis-master"},"annotations": {"kubernetes.io/created-by": "{\"kind\":\"SerializedReference\",\"apiVersion\":\"v1\",\"reference\":{\"kind\":\"ReplicationController\",\"namespace\":\"default\",\"name\":\"redis-master\",\"uid\":\"1c244c76-550c-11e5-820ffa163e78675c\",\"apiVersion\":\"v1\",\"resourceVersion\":\"188599\"}}"}},"spec": {"containers": [{"name": "master","image": "192.168.12.7:5000/redis","ports": [{"containerPort": 6379,"protocol": "TCP"}],"resources": {},"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"}],"restartPolicy": "Always","dnsPolicy": "ClusterFirst","nodeName": "192.168.12.9"},"status": {"phase": "Running","conditions": [{"type": "Ready","status": "True"}],"hostIP": "192.168.12.9","podIP": "10.1.83.7","startTime": "2015-09-07T03:03:52Z","containerStatuses": [{"name": "master","state": {"running": {"startedAt": "2015-09-07T06:26:43Z"}},"lastState": {"terminated": {"exitCode": 0,"startedAt": "2015-09-07T05:44:41Z","finishedAt": "2015-09-07T06:23:50Z","containerID": "docker://481b2e4ac6c8ff43d95c1bb995e9f1a30c6c524ef6ac944d126fd43d703ecfa6"}},"ready": true,"restartCount": 1,"image": "192.168.12.7:5000/redis","imageID": "docker://95af5842ddb9b03f7c6ec7601e65924cec516fcedd7e590ae31660057085cf67","containerID": "docker://6c97a334dee2927577ee4b29822757074ee083fec9deb52c7fa57bc99bf4d624"}]}},{"metadata": {"name": "redis-slave-ohole","generateName": "redis-slave-","namespace": "default","selfLink": "/api/v1/namespaces/default/pods/redis-slave-ohole","uid": "38eb8838-550c-11e5-820f-fa163e78675c","resourceVersion": "244971","creationTimestamp": "2015-09-07T02:57:48Z","labels": {"name": "redis-slave"},"annotations": {"kubernetes.io/created-by": "{\"kind\":\"SerializedReference\",\"apiVersion\":\"v1\",\"reference\":{\"kind\":\"ReplicationController\",\"namespace\":\"default\",\"name\":\"redis-slave\",\"uid\":\"3895c897-550c-11e5-820ffa163e78675c\",\"apiVersion\":\"v1\",\"resourceVersion\":\"188832\"}}"}},"spec": {"containers": [{"name": "slave","image": "192.168.12.7:5000/kubernetes/redis-slave:v2","ports": [{"containerPort": 6379,"protocol": "TCP"}],"resources": {},"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"}],"restartPolicy": "Always","dnsPolicy": "ClusterFirst","nodeName": "192.168.12.9"},"status": {"phase": "Running","conditions": [{"type": "Ready","status": "True"}],"hostIP": "192.168.12.9","podIP": "10.1.83.6","startTime": "2015-09-07T03:06:30Z","containerStatuses": [{"name": "slave","state": {"running": {"startedAt": "2015-09-07T06:26:43Z"}},"lastState": {"terminated": {"exitCode": 0,"startedAt": "2015-09-07T05:44:41Z","finishedAt": "2015-09-07T06:26:08Z","containerID": "docker://26aa7816cc8b010376a4412562042b5733351147cf2a075a0e930fae83855332"}},"ready": true,"restartCount": 1,"image": "192.168.12.7:5000/kubernetes/redis-slave:v2","imageID": "docker://48a1e2b8f8599f083b09bd0b451dc062fe3ca54f4255d81e89699667cfddc0aa","containerID": "docker://f90b7d96ad4146886c386f3529e652c71f6bb566a443d972150f52b849ab748a"}]}},{"metadata": {"name": "redis-slave-rosj0","generateName": "redis-slave-","namespace": "default","selfLink": "/api/v1/namespaces/default/pods/redis-slave-rosj0","uid": "27d7e47e-5524-11e5-8634-fa163e78675c","resourceVersion": "244954","creationTimestamp": "2015-09-07T05:49:08Z","labels": {"name": "redis-slave"},"annotations": {"kubernetes.io/created-by": "{\"kind\":\"SerializedReference\",\"apiVersion\":\"v1\",\"reference\":{\"kind\":\"ReplicationController\",\"namespace\":\"default\",\"name\":\"redis-slave\",\"uid\":\"3895c897-550c-11e5-820ffa163e78675c\",\"apiVersion\":\"v1\",\"resourceVersion\":\"233664\"}}"}},"spec": {"containers": [{"name": "slave","image": "192.168.12.7:5000/kubernetes/redis-slave:v2","ports": [{"containerPort": 6379,"protocol": "TCP"}],"resources": {},"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"}],"restartPolicy": "Always","dnsPolicy": "ClusterFirst","nodeName": "192.168.12.9"},"status": {"phase": "Running","conditions": [{"type": "Ready","status": "True"}],"hostIP": "192.168.12.9","podIP": "10.1.83.3","startTime": "2015-09-07T05:49:49Z","containerStatuses": [{"name": "slave","state": {"running": {"startedAt": "2015-09-07T06:26:43Z"}},"lastState": {"terminated": {"exitCode": 0,"startedAt": "2015-09-07T05:49:49Z","finishedAt": "2015-09-07T06:26:08Z","containerID": "docker://dbde2db205b318c2a9d9f16a7dec6bbd9c5caac60edb98c036b74088c5cb21b2"}},"ready": true,"restartCount": 1,"image": "192.168.12.7:5000/kubernetes/redis-slave:v2","imageID": "docker://48a1e2b8f8599f083b09bd0b451dc062fe3ca54f4255d81e89699667cfddc0aa","containerID": "docker://4788ca0dbf08e4d0bdfae5ab2ea00f014b7c13b5a3b86a0d332b8441c3ca9172"}]}},{"metadata": {"name": "web-app-hu9po","generateName": "web-app-","namespace": "default","selfLink": "/api/v1/namespaces/default/pods/web-app-hu9po","uid": "cd07673b-5603-11e5-9b89-fa163e78675c","resourceVersion": "647661","creationTimestamp": "2015-09-08T08:30:02Z","labels": {"name": "web-app"},"annotations": {"kubernetes.io/created-by": "{\"kind\":\"SerializedReference\",\"apiVersion\":\"v1\",\"reference\":{\"kind\":\"ReplicationController\",\"namespace\":\"default\",\"name\":\"web-app\",\"uid\":\"ccf3572a-5603-11e5-9b89fa163e78675c\",\"apiVersion\":\"v1\",\"resourceVersion\":\"646563\"}}"}},"spec": {"containers": [{"name": "web-app","image": "training/webapp","ports": [{"containerPort": 5000,"protocol": "TCP"}],"resources": {},"terminationMessagePath": "/dev/termination-log","imagePullPolicy": "IfNotPresent"}],"restartPolicy": "Always","dnsPolicy": "ClusterFirst","nodeName": "192.168.12.9"},"status": {"phase": "Running","conditions": [{"type": "Ready","status": "True"}],"hostIP": "192.168.12.9","podIP": "10.1.83.33","startTime": "2015-09-08T08:34:13Z","containerStatuses": [{"name": "web-app","state": {"running": {"startedAt": "2015-09-08T08:34:12Z"}},"lastState": {},"ready": true,"restartCount": 0,"image": "training/webapp","imageID": "docker://02a8815912ca800f99b7d912485e8c618260e27c6de8d7a161b356b322d5c121","containerID": "docker://5e6b038129d7e032c4e404f16f57326d99618d0d697a008bfec1b1e52a4f4021"}]}}]}`)
	return body	
}



