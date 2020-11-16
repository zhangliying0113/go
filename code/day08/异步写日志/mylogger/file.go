package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

var (
	// MaxSize 日志通道缓冲区大小
	MaxSize = 50000
)

// FileLogger 文件日志结构体
type FileLogger struct {
	Level       LogLevel
	filePath    string // 日志文件保存的路径
	filename    string // 日志文件保存的文件名
	errFileName string // 错误日志单独记下来
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64 // 文件大小，超过文件大小进行切割
	logChan     chan *logMsg
}

type logMsg struct {
	level     LogLevel
	msg       string
	funcName  string
	fileName  string
	timestamp string
	line      int
}

// NewFileLogger 文件日志构造函数
func NewFileLogger(leverStr, fp, fn string, maxFileSize int64) *FileLogger {
	level, err := parseLogLevel(leverStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       level,
		filePath:    fp,
		filename:    fn,
		maxFileSize: maxFileSize,
		logChan:     make(chan *logMsg, MaxSize),
	}
	err = fl.initFile() // 按照文件路径和文件名将文件打开
	if err != nil {
		panic(err)
	}
	return fl
}

// initFile 根据指定的文件日志路径和文件名打开日志文件
func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.filename)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err:%v\n", err)
	}
	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err:%v\n", err)
	}
	// 日志文件全部打开
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	// 开启一个goroutine执行后台写日志的任务，开启多个 goroutine 会出现获取已经关闭文件信息的报错，加锁解决会耗费内存资源，所以用一个 goroutine
	go f.writeLogBackground()
	return nil
}

// 关闭文件
func (f *FileLogger) close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}

// 判断是否需要记录该日志
func (f *FileLogger) enable(logLevel LogLevel) bool {
	return logLevel >= f.Level
}

// 判断日志文件是否需要切割
func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat() // 拿到文件信息
	if err != nil {
		fmt.Printf("get file info failed, err:%v\n", err)
		return false
	}
	// 如果当前文件大小大于等于日志文件最大值，返回 true，需要切割文件
	return fileInfo.Size() > f.maxFileSize
}

// 文件切割
func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	nowStr := time.Now().Format("20060102140405000")
	fileInfo, err := file.Stat() // 拿到文件信息
	if err != nil {
		fmt.Printf("get file info failed, err:%v\n", err)
		return nil, err
	}

	logName := path.Join(f.filePath, fileInfo.Name())      // 拿到当前的日志文件完整路径
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr) // 拼接一个日志文件备份
	// 1. 关闭文件
	file.Close() // 文件关闭之后将不能拿到文件对象，所以先拿文件信息再关闭文件
	// // 2.备份一下 rename  xx.log  ->  xx.log.bak201908031709
	os.Rename(logName, newLogName)
	// 3. 打开一个新的日志文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new log file field, err:%v\n", err)
		return nil, err
	}
	// 4. 将打开的新日志文件对象赋值给 file
	return fileObj, nil
}

// 后台 goroutine 异步写日志
func (f *FileLogger) writeLogBackground() {
	for {
		if f.checkSize(f.fileObj) {
			newFile, err := f.splitFile(f.fileObj) // 日志文件
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		select {
		case logTmp := <-f.logChan:
			// 把日志拼出来
			logInfo := fmt.Sprintf("[%s] [%s] [%s:%s:%d] %s\n", logTmp.timestamp, getLogString(logTmp.level), logTmp.fileName, logTmp.funcName, logTmp.line, logTmp.msg)
			fmt.Fprintf(f.fileObj, logInfo)
			if logTmp.level >= ERROR {
				if f.checkSize(f.errFileObj) {
					newFile, err := f.splitFile(f.errFileObj) // 日志文件
					if err != nil {
						return
					}
					f.errFileObj = newFile
				}
				// 如果要记录的日志大于等于ERROR级别,我还要在err日志文件中再记录一遍
				fmt.Fprintf(f.errFileObj, logInfo)
			}
		default:
			// 取不到日志先休息500毫秒
			time.Sleep(time.Millisecond * 500)
		}
	}
}

// 记录日志的方法
func (f *FileLogger) log(lv LogLevel, format string, arg ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, arg...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		// 先把日志发送到通道中
		// 1. 造一个logMsg对象
		logTmp := &logMsg{
			level:     lv,
			msg:       msg,
			funcName:  funcName,
			fileName:  fileName,
			timestamp: now.Format("2006-01-02 15:04:05"),
			line:      lineNo,
		}
		select {
		case f.logChan <- logTmp:
		default:
			// 把日志就丢掉保证不出现阻塞
		}
	}
}

// Debug 级别日志记录
func (f *FileLogger) Debug(format string, arg ...interface{}) {
	f.log(DEBUG, format, arg...)
}

// Trace 级别日志记录
func (f *FileLogger) Trace(format string, arg ...interface{}) {
	f.log(TRACE, format, arg...)
}

// Info 级别日志记录
func (f *FileLogger) Info(format string, arg ...interface{}) {
	f.log(INFO, format, arg...)
}

// Warning 级别日志记录
func (f *FileLogger) Warning(format string, arg ...interface{}) {
	f.log(WARNING, format, arg...)
}

// Error 级别日志记录
func (f *FileLogger) Error(format string, arg ...interface{}) {
	f.log(ERROR, format, arg...)
}

// Fatal 级别日志记录
func (f *FileLogger) Fatal(format string, arg ...interface{}) {
	f.log(FATAL, format, arg...)
}
