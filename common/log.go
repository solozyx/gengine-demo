package common

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"path"
	"runtime"
)

func InitLog(writer ...io.Writer) {
	log.SetFormatter(&log.JSONFormatter{})
	log.AddHook(&ContextHook{})
	log.SetLevel(log.TraceLevel) // 设置需要记录的日志级别
	for _, item := range writer {
		log.SetOutput(item)
	}
}

type ContextHook struct{}

func (p *ContextHook) Levels() []log.Level {
	return log.AllLevels
}

func (p *ContextHook) Fire(entry *log.Entry) error {
	if pc, file, line, ok := runtime.Caller(10); ok {
		funcName := runtime.FuncForPC(pc).Name()
		entry.Data["source"] = fmt.Sprintf("%s:%v:%s", path.Base(file), line, path.Base(funcName))
	}
	return nil
}
