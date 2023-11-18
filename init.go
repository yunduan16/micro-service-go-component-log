package log

import (
    "errors"
    "fmt"
    "github.com/lestrrat-go/file-rotatelogs"
    "github.com/sirupsen/logrus"
    "os"
    "path/filepath"
    "runtime"
    "strconv"
)

type Fields map[string]interface{}

// LoggerObj 全局的日志结构体
type LoggerObj struct {
    logger     *logrus.Logger
    hostName   string
    pid        int
    callerDeep int
}

// 默认的日志文件
var defaultLogger *LoggerObj

func init(){
    // step1 初始化log
    defaultLogger = &LoggerObj{
        logger: logrus.New(),
    }
}

func New(opts ...Option) (*LoggerObj, error) {
    l := &LoggerObj{
        logger: logrus.New(),
    }
    l.hostName, _ = os.Hostname()
    if l.hostName == "" {
        l.hostName = "unKnowHost"
    }
    l.pid = os.Getpid()

    // step1 获取参数options
    var options = defaultOptions
    for _, opt := range opts {
        opt(&options)
    }
    l.logger.SetLevel(options.logLevel)
    l.logger.SetFormatter(&HHZJSONFormatter{
        TimestampFormat:   options.timestampFormat,
        DisableHTMLEscape: true,
    })
    l.callerDeep = options.callerDeep

    if options.fileName == "" {
        return nil, errors.New("fileName 不能为空")
    }
    // step2 如果日志不存在，创建日志文件，使用rotatelogs
    _, ok := IsExists(filepath.Dir(options.fileName))
    if !ok {
        if err := os.MkdirAll(filepath.Dir(options.fileName), 0766); err != nil {
            return nil, errors.New("创建日志文件的目录失败！")
        }
    }
    rotatePath := fmt.Sprintf("%s", options.fileName)+".%Y%m%d%H"
    logWriter, _ := rotatelogs.New(rotatePath, rotatelogs.WithLinkName(options.fileName),
        rotatelogs.WithRotationTime(options.rotationTime), rotatelogs.WithMaxAge(options.maxAge))
    l.logger.SetOutput(logWriter)
    return l, nil
}

func (l *LoggerObj) Trace(logFields Fields, args ...interface{}) {
    l.logger.WithFields(l.GetCommonFields(logFields)).Trace(args...)
}

func (l *LoggerObj) Info(logFields Fields, args ...interface{}) {
    l.logger.WithFields(l.GetCommonFields(logFields)).Info(args...)
}

func (l *LoggerObj) Error(logFields Fields, args ...interface{}) {
    l.logger.WithFields(l.GetCommonFields(logFields)).Error(args...)
}

func (l *LoggerObj) Warn(logFields Fields, args ...interface{}) {
    l.logger.WithFields(l.GetCommonFields(logFields)).Warn(args...)
}

func (l *LoggerObj) Debug(logFields Fields, args ...interface{}) {
    l.logger.WithFields(l.GetCommonFields(logFields)).Debug(args...)
}

// GetCommonFields 讲一些公共参数返回
func (l *LoggerObj) GetCommonFields(logFields Fields) logrus.Fields {
    fields := logrus.Fields(logFields)
    fields["hostName"] = l.hostName
    fields["pid"] = l.pid
    _, file, line, _ := runtime.Caller(l.callerDeep)
    fields["path"] = file + ":" + strconv.Itoa(line)
    return fields
}