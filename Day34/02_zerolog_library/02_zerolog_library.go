package main

import (
	//"github.com/rs/zerolog"	// 全局配置

	"encoding/json"
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log" // 冲突：使用这个log就不能使用官方库的log，反之亦然。想使用需要给其中一个加别名
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// 时间字段
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs // 以毫秒为单位的时间戳
	if err := os.MkdirAll("logs", os.ModePerm); err != nil {
		panic(err)
	}

	// 自定义错误输出的字段
	zerolog.ErrorFieldName = "e"

	// roll滚动日志。日志文件会不断膨胀，要进行切割：按时间、按大小等，支持压缩，保留文件
	ll := &lumberjack.Logger{
		Filename:   "logs/a.log",
		MaxSize:    1,    // 兆
		MaxBackups: 3,    // 除去当前写入的foo.log外，历史文件有几个
		MaxAge:     28,   // 天
		Compress:   true, // 是否将历史文件进行压缩
	}
	defer ll.Close()
	// 基本用法
	log.Print("这是第一个测试的消息") // log提供的导出快捷方法，相当于log.Debug().Msg()
	//log.Fatal().Msg("这是一个fatal消息")
	//log.Panic()
	fmt.Println(zerolog.GlobalLevel(), "###") // 默认trace级别-1
	log.Trace().Caller().Msg("这是一个trace消息")   // -1，Caller：展示哪个文件的哪行代码调用

	zerolog.SetGlobalLevel(zerolog.ErrorLevel) // 定义级别为error
	fmt.Println(zerolog.GlobalLevel(), "###")  // 级别改为error

	// 自定义 Logger
	defalt := log.Logger
	//defalt.Level() // 不是设置当前logger，而是创建一个子logger
	fmt.Println(defalt.GetLevel(), int(defalt.GetLevel()))

	// 自定义级别
	child := defalt.Level(zerolog.WarnLevel)
	child.Debug().Msg("这是一个子logger") // 产生一个debug消息级别的消息
	child.Info().Msg("child info msg")
	child.Warn().Msg("child warm msg")
	child.Error().Msg("child error msg") // msg level >= logger level 才会输出

	// 构造全新logger
	child2 := zerolog.New(os.Stdout) // -1
	child2.Trace().Msg("child2 trace msg")

	// 添加不同类型的key-value字段（添加上下文）
	child.Error().Str("school", "gxmu").Msg("child error msg with field")
	child.Error().Bool("success", true).Msg("child error msg")
	child.Error().Str("student", "Faye").Ints("score", []int{100, 100, 100}).Msg("student")
	child.Error().Str("student", "Faye").Ints("score", []int{100, 100, 100}).Send() // 直接发送

	// 反序列化（单行）
	s := `{"level":"error","student":"Faye","score":[100,100,100],"time":1774152174379}`
	var data = make(map[string]any)
	err2 := json.Unmarshal([]byte(s), &data)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Printf("%T %[1]v\n", data)

	// 存储到文件
	f3, err3 := os.OpenFile("logs/a.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err3 != nil {
		panic(err3)
	}
	defer f3.Close()
	out := zerolog.New(f3).With().Timestamp().Logger()
	out.Error().Msg("out error msg")

	// zerolog内部实现了群发
	writers := zerolog.MultiLevelWriter(ll, f3, os.Stdout, os.Stderr) // LevelWriter实现Writer方法，内部就要遍历
	loggers := zerolog.New(writers).With().Timestamp().Logger()
	loggers.Error().Msg("loggers error msg")

}
