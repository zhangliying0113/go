package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
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
}

// NewFileLogger 文件日志构造函数
func NewFileLogger(leverStr, fp, fn string, maxSize int64) *FileLogger {
	level, err := parseLogLevel(leverStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       level,
		filePath:    fp,
		filename:    fn,
		maxFileSize: maxSize,
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

// 记录日志的方法
func (f *FileLogger) log(lv LogLevel, format string, arg ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, arg...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		if f.checkSize(f.fileObj) {
			newFile, err := f.splitFile(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s %s %d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), funcName, fileName, lineNo, msg)
		if lv >= TRACE {
			if f.checkSize(f.errFileObj) {
				newErrFile, err := f.splitFile(f.errFileObj)
				if err != nil {
					return
				}
				f.errFileObj = newErrFile
			}
			// 如果记录日志大于等于 ERROR 级别，我还要在 err 日志文件中再记录一遍
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s %s %d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), funcName, fileName, lineNo, msg)
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
