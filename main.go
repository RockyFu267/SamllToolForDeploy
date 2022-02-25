package main

import (
	"SamllToolForDeploy/basefunc"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
)

var source_path *string = flag.String("sourcepath", "./", "Use -sourcepath <source_path>")
var target_path *string = flag.String("targetpath", "", "Use -targetpath <target_path>")
var namespace *string = flag.String("namespace", "default", "Use -namespace <namespace>")

func main() {
	//获取参数
	flag.Parse()
	fmt.Println("hi buddy")
	//获取当前路径
	pwdPath, err := os.Getwd()
	if err != nil {
		log.Println(err)
		//***结束***
		return
	}
	fmt.Println("当前路径：", pwdPath)
	//赋值
	targetpathTmp := *target_path
	if len(*target_path) == 0 {
		targetpathTmp = pwdPath + "/output/"
	}

	fmt.Println(targetpathTmp)
	//获取所有文件列表
	tgzFileList, err := basefunc.ReadFileListNew(*source_path)
	if err != nil {
		log.Println(err)
		//***结束***
		return
	}
	fmt.Println("1111111111111111111111111111111111111111111111111111")
	//检查目标输出路径
	err = basefunc.CheckDir(targetpathTmp)
	if err != nil {
		log.Println(err)
		//***结束***
		return
	}
	fmt.Println("2222222222222222222222222222222222222222222222222222")
	//解压文件到目标路径
	err = basefunc.TarTarget(tgzFileList, *source_path, targetpathTmp)
	if err != nil {
		log.Println(err)
		//***结束***
		return
	}
	fmt.Println("333333333333333333333333333333333333333333333333333")
	err = basefunc.DryRun(targetpathTmp, *namespace)
	if err != nil {
		log.Println(err)
		//***结束***
		return
	}
	fmt.Println("44444444444444444444444444444444444444444444444444")
	fmt.Println(tgzFileList)
	fmt.Println(*target_path)
	fmt.Println(os.Hostname())
	fmt.Println(runtime.GOARCH) //系统构架 386、amd64
	fmt.Println(runtime.GOOS)   //系统版本 windows
}
