# 日志
- zerolog库
  - log官方库实际使用并不方便
  - 官网 https://zerolog.io/

  - 默认Json格式输出到标准错误输出strerr，有级别
  - 格式
    - level   消息的级别
      - 快捷方法，本质上调用的是log.Logger的方法
        - log.Debug() 方法产生消息带上消息级别

```txt
                        return Logger.Debug()
                    log.Fatal().Msg
```

          - Msg用来挂内容
          - Debug、Fatal等设置该消息的级别
    - time    很重要，业务数据没有时间就没有意义
    - message 关系到内容

  - 日志输出都需要一个日志记录器Logger
  - log.Logger
    - 缺省logger
      - 不需要你显式创建，就可以直接使用的日志记录器
    - 为了方便一般会提供一个缺省的日志记录器
      - log包内建的缺省的全局导出的Logger

```txt
                var Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
```

      - 通过日志记录器产生消息（不同消息级别）输出到指定的 文件对象（stdout、stderr、file对象）
      - 缺省logger的级别是 Trace -1，不能拦着任何消息，消息最低级别是 -1
    - 自定义 Logger
      - 缺省logger不能满足时使用

```txt
                    1、zerolog.New(io.Writer)
```

        - 2、基于父logger.Xxx() 返回基于父logger的子logger

  - New
    - zerolog.New(io.Writer)  默认logger的级别是 Trace -1
    - 完全构造一个新logger，区别于logger.Level的构造一个子logger

  - level级别
    - 等级
      - DebugLevel Level = 0
      - InfoLevel   1
      - WarnLevel   2
      - ErrorLevel  3
      - FatalLevel  4
      - PanicLevel  5
      - NoLevel     6
      - Disabled    7
      - TraceLevel Level = -1
    - 消息级别
      - logger.Debug().Msg()通过日志记录器产生一个某级别的消息
      - 表示消息的重要程度
    - 日志记录器的级别
      - logger.GetLevel() zerolog.Level 返回当前logger的级别，是一个整数
      - logger.Level()                  返回一个新的子logger，拥有用户设置的级别
      - 调整日志记录器级别为了控制哪些消息输出，如果级别过低，回导致很多消息输出，从而使得磁盘文件膨胀过快
    - gLevel
      - zerolog的全局level，影响所有logger，提高所有logger的最低门槛
    - 日志输出条件
      - msg level >= max(logger level, gLevel) 才会输出
  - 加字段
    - 1、临时加字段 在当前输出方法上跟logger.Debug().Str()
    - 2、定义在logger上

  - 总结
    - 日志库输出需要Logger（为了方便，往往库会提供快捷方法调用缺省logger）
    - 调用logger的不同方法会产生不同消息级别的消息

  - 存储到文件
    - 写，打开文件的写能力（只写）
    - os.OpenFile => f, err f就是读或写能力的对象

```txt
            zerolog.New(w io.Writer)
```

      - w对象要实现Writer接口 的方法 Writer(p []bytr) (n int, err error)
