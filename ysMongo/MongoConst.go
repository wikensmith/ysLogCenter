package ysMongo

var ChannelNum = 10 // mongo 连接池的连接数量

const (
	ConnectURL = "127.0.0.1:27017" // 连接地址
	DBName     = "test"
)

type LogInfo struct {
	Group     string // 小组名称
	Project   string //  项目名称
	Level     string // 日志等级
	Title     string // 日志抬头
	Msg       string // 日志信息
	Context   string // 上下文
	Timestamp string // 时间
}
