package root

import (
	"DirScan/dic"
	"DirScan/golimit"
	"DirScan/scan"
	"github.com/gookit/color"
	"strings"
	"sync"
	"time"
)

func GoDirScan(url string, filename string, num int, errtime int64, timesleep int64, openerrstr int, errstr string) {
	dicall := dic.Readfile(filename)
	r, err := scan.Goscan(url, "", errtime, 0)
	if err != nil {
		color.Warn.Println("[Err] 目标访问错误，可能被ban了！")
		return
	}
	color.Red.Println("[Info] Dirscan|Try| By T00ls.Net; 版本:0.0.2")
	if openerrstr == 1 {
		color.Red.Println("[Info] 已开启自定义错误关键词:", errstr)
	}
	color.Red.Println("[Info] 目标地址:", url)
	color.Red.Println("[Info] 当前线程:", num)
	color.Red.Println("[Info] 超时时间:", errtime)
	color.Red.Println("[Info] 目标相关容器:", r.Server, r.Powered)
	color.Red.Println("[Info] 加载字典数量:", len(dicall))
	color.Red.Println("[Info] 开始扫描中.")
	color.Red.Println("---------------------------------------")
	//writefile.Write(url, "[Info] 目标地址: "+url+"\n")
	//writefile.Write(url, "[Info] 目标相关容器: "+r.Server+r.Powered+"\n")
	g := golimit.NewG(num) //设置线程数量
	wg := &sync.WaitGroup{}
	beg := time.Now()
	for i := 0; i < len(dicall); i++ {
		wg.Add(1)
		task := dicall[i]
		g.Run(func() {
			respBody, err := scan.Goscan(url, task, errtime, timesleep)
			if err != nil {
				//color.Warn.Println("目标访问错误，可能被ban了！")
				wg.Done()
				return
			}
			if strings.Contains(respBody.Body, errstr) == false {
				if respBody.StatusCode == 200 {
					color.Info.Println("[200] ", respBody.Res+"   [len]", respBody.Bodylen)
					//writefile.Write(url, "[200] "+respBody.Res+"\n")
				} else if respBody.StatusCode == 403 {
					color.Warn.Println("[403] ", respBody.Res+"   [len]", respBody.Bodylen)
					//writefile.Write(url, "[403] "+respBody.Res+"\n")
				} else if respBody.StatusCode == 302 {
					color.Warn.Println("[302] ", respBody.Res+"   [len]", respBody.Bodylen)
					//writefile.Write(url, "[302] "+respBody.Res+"\n")
				}
			}
			wg.Done()
		})
	}
	wg.Wait()
	color.Red.Printf("[info] 扫描完成！当前用时: %fs", time.Now().Sub(beg).Seconds())
}
