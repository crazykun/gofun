package log

import (
	"fmt"
	"gofun/conf"
	. "gofun/pkg/tools"
	"log"
	"os"
	"time"
)

// Log 记录一般日志，定期自动删除
func Log(_txt string) {

	// 创建文件夹
	filepath := conf.Config.AppPath + "/log/"
	// 创建日期文件夹
	has, _ := HasFile(filepath)
	if !has {
		err := os.Mkdir(filepath, os.ModePerm)
		if err != nil {
			fmt.Printf("不能创建文件夹1=[%v]\n", err)
			return
		}
	}

	// 文件信息
	fileName := "log_" + GetTimeDate("Ymd") + ".log"
	file, err := os.OpenFile(filepath+fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("文件创建失败", filepath, err)
		//panic(err)
		return
	}
	// 延迟关闭文件
	defer file.Close()

	// 写入文件内容
	_, err = file.WriteString(_txt + "\n")
	if err != nil {
		log.Println("文件写入失败", filepath, err)
		return
	}

	// 删除昨天日志文件
	delFile := "log_" + time.Now().AddDate(0, 0, -1).Format("20060102") + ".log"
	delFilepath := filepath + delFile
	err1 := os.RemoveAll(delFilepath)
	if err1 != nil {
		log.Println("旧日志删除失败=", err1)
	}

}

// Error 记录错误日志
func Err(_txt string) {
	// 创建文件夹
	filepath := conf.Config.AppPath + "/log/"
	// 创建日期文件夹
	has, _ := HasFile(filepath)
	if !has {
		err := os.Mkdir(filepath, os.ModePerm)
		if err != nil {
			fmt.Printf("不能创建文件夹1=[%v]\n", err)
		}
	}

	// 文件信息
	fileName := "err_" + GetTimeDate("Ymd") + ".log"
	file, err := os.OpenFile(filepath+fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	// 延迟关闭文件
	defer file.Close()
	// 写入文件内容
	file.WriteString(_txt + "\n")

	// 删除老文件夹
	delFile := "err_" + time.Now().AddDate(0, 0, -1).Format("20060102") + ".log"
	delFilepath := filepath + delFile
	err1 := os.RemoveAll(delFilepath)
	if err1 != nil {
		log.Println("旧错误删除失败=", err1)
	}

}
