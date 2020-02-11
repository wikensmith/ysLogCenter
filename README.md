-日志格式
参考：http://ulog.colipu.com/<br>
```golang
{
	"group": "test",  # 团队名称
	"project": "wiken",  # 项目名称
	"leve": "info",   # 日志等级info, error, warn
	"title": "wiken_test",  # 日志抬头

     // 日志格式， 如果返回是异常， success 为false， 否则为true
	"msg": {"success": "true", "msg": "test_msg"}, 
	"context": "",  # 当前日志的上下文， 可以是程序的传入信息等。
	"timestamp": "",  # 时间戳 2020-02-11T16:53:37+08:00(东8区时间)	
}
```