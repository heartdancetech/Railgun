package common

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"
)

// NewReverseProxy 创建反向代理处理方法
func NewReverseProxy() *httputil.ReverseProxy {

	var target = struct{
		ServiceAddress string
		ServicePort int
	}{
		ServiceAddress: "localhost",
		ServicePort: 8080,
	}
	//创建Director
	director := func(req *http.Request) {

		//查询原始请求路径，如：/arithmetic/calculate/10/5
		reqPath := req.URL.Path
		if reqPath == "" {
			return
		}
		//按照分隔符'/'对路径进行分解，获取服务名称serviceName
		pathArray := strings.Split(reqPath, "/")
		serviceName := pathArray[1]
		fmt.Println(serviceName)

		//重新组织请求路径，去掉服务名称部分
		destPath := strings.Join(pathArray[2:], "/")

		//随机选择一个服务实例
		//tgt := result[rand.Int()%len(result)]
		//logger.Log("service id", tgt.ServiceID)

		//设置代理服务地址信息
		req.URL.Scheme = "http"
		req.URL.Host = fmt.Sprintf("%s:%d", target.ServiceAddress, target.ServicePort)
		req.URL.Path = "/" + destPath
	}

	modifyResponse := func(res *http.Response) error {
		fmt.Println(res.Body)
		return nil
	}

	return &httputil.ReverseProxy{Director: director, ModifyResponse: modifyResponse}

}
