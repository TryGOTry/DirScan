package scan

import (
	"crypto/tls"
	"github.com/go-resty/resty/v2"
	"strings"
	"time"
)

//var (
//	title = `<title>([\s\S]+?)</title>`
//)

type Webinfo struct {
	StatusCode int
	//Title      string
	Server  string
	Powered string
	Body    string
	Res     string //成功的结果
	Bodylen int    //返回包长度
}

func Goscan(url string, dir string, errtime int64, timesleep int64) (Webinfo, error) {
	time.Sleep(time.Duration(timesleep) * time.Second) //设置延时时间
	var t string
	var Web Webinfo
	t = url
	a := t[len(t)-1:]
	if strings.Contains(t, "https://") {
		//log.Println("https")
	} else {
		t = "http://" + url
	}
	if a != "/" {
		t = url + "/" //判断结尾是否为/，如果不是，那就加上
	}
	client := resty.New().SetTimeout(time.Duration(errtime)* time.Second).SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}) //忽略https证书错误，设置超时时间
	client.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36")
	resp, err := client.R().EnableTrace().Get(t + dir) //开始请求扫描
	if err != nil {
		//log.Println(err)
		return Web, err
	}
	str := resp.Body()
	body := string(str)
	//re1 := regexp.MustCompile(title) //正则取标题
	//titlename := re1.FindAllStringSubmatch(body, 1)
	//fmt.Println(body)
	Web.StatusCode = resp.StatusCode()
	Web.Powered = resp.Header().Get("X-Powered-By")
	//Web.Title = titlename[0][1]
	Web.Server = resp.Header().Get("server")
	Web.Body = body
	Web.Res = t + dir
	Web.Bodylen = len(body)
	return Web, nil
}
