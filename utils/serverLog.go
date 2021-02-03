package utils

import (
	"github.com/hhkbp2/go-logging"
	"os"
	"time"
)

func WriteLog(message string)  {
	filePath := "./serverLog.log"
	fileMode := os.O_APPEND
	bufferSize := 0
	buffersFlushTime := 30 * time.Second
	inputChanSize := 1
	fileMaxByte := uint64(100*1024*1024)
	backupCount := uint32(9)
	handler := logging.MustNewRotatingFileHandler(
		filePath,fileMode,bufferSize,buffersFlushTime,inputChanSize,fileMaxByte,backupCount)
	format := "%(asctime)s %(levelname)s : %(message)s"
	dateFormat := "%Y-%m-%d %H:%M:S.%3n"
	formatter := logging.NewStandardFormatter(format, dateFormat)
	handler.SetFormatter(formatter)
	logger := logging.GetLogger("")
	logger.SetLevel(logging.LevelInfo)
	logger.AddHandler(handler)
	defer logging.Shutdown()
	logger.Infof(" %s",message)
}
