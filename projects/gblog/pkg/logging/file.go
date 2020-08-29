package logging

import (
	"fmt"
	"gblog/pkg/file"
	"os"
	"time"

	"gblog/pkg/setting"
)

var logSetting = setting.AppSetting

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", logSetting.RuntimeRootPath, logSetting.LogSavePath)
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", logSetting.LogSaveName, time.Now().Format(logSetting.TimeFormat), logSetting.LogFileExt)

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		logSetting.LogSaveName,
		time.Now().Format(logSetting.TimeFormat),
		logSetting.LogFileExt)
}

func openLogFile(filePath, fileName string) (*os.File, error){
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	perm := file.CheckPermission(src)
	if perm {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err = file.IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := file.Open(src + fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile: %v", err)
	}

	return f, nil
}