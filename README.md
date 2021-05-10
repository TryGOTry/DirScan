# Dirscan
一个简单的目录扫描器
# 说明
第一次用golang上手写目录扫描器，原理很简单，就不多说了。
（本来最开始写了自动保存扫描结果的，但是后面测试发现好像没必要，就直接注释了。）

实现的功能:1.可自定义超时时间(默认为3s) 2.可自定义延时请求时间(默认0) 3.可自定义线程(默认5)4.可自定义错误关键词判断(如果包含当前设置的关键词，则跳过该页面.)
未实现的功能:1.设置代理 2.自定义ua头 3.待定

当前测试好像还没有发现bug，表哥们有啥建议可以留言。
# 使用说明
```
     -errstr string
            自定义错误关键词,必须先-o 1
      -f string
            加载字典（必填）
      -o int
            1表示开始自定义错误关键词,默认关闭
      -s int
            线程数 (default 5)
      -timeout int
            超时时间 (default 3)
      -timesleep int
            延时时间
      -u string
            url地址 (必填)
```
# 截图
![运行截图](https://www.nctry.com/wp-content/uploads/2021/05/6663.png)
