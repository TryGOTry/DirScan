package main
/*
* @Author: Try
* @Date:   2021/05/04
 */
import (
	"DirScan/root"
	"flag"
)

func main() {
	url := flag.String("u", "", "url地址")
	filename := flag.String("f", "", "加载字典")
	xiancheng := flag.Int("s", 5, "线程数")
	openerrstr := flag.Int("o", 0, "1表示开始自定义错误关键词,默认关闭")
	errtime := flag.Int64("timeout", 3, "超时时间")
	errstr := flag.String("errstr", "", "自定义错误关键词,必须先-o 1")
	timesleep := flag.Int64("timesleep", 0, "延时时间")
	flag.Parse()
	if *url != "" && *filename != "" && *openerrstr == 0 {
		root.GoDirScan(*url, *filename, *xiancheng, *errtime, *timesleep, *openerrstr, "@#!#asddddddddd122222222222asd")
	} else if *url != "" && *filename != "" && *openerrstr == 1 && *errstr != "" {
		root.GoDirScan(*url, *filename, *xiancheng, *errtime, *timesleep, *openerrstr, *errstr)
	} else {
		flag.Usage()
	}
}
