### logger 是golang的日志库 ，基于对golang内置log的封装。

打印日志有4个方法 Debug，Info，Warn, Error

```
llog.Xxxxf() 表示格式化输出,和fmt.Printf()类似
llog.Xxxx()  表示常规输出
```

可自己设置输出方式
```
llog.SetLevel()
设置日志等级

llog.SetPath()
llog.SetLogFile()
设置日志文件路径

llog.SetPut() 
llog.SetWrite()
设置输出到终端或者写入文件里
```

### 输出日志
```
llog.Debugf("debug test %s","Debug")
llog.Infof("debug test %s","Info")

llog.Debug("debug test Debug")
llog.Info("debug test Info")
```


这个算是一个小的模块吧,对自己学习过程的一些总结