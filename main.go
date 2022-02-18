package main

import (
	"SamllToolForDeploy/basefunc"
	"flag"
	"fmt"
	"log"
	"os"
)

var source_path *string = flag.String("sourcepath", "./", "Use -sourcepath <source_path>")
var target_path *string = flag.String("targetpath", "", "Use -targetpath <source_path>")

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
		targetpathTmp = pwdPath + "/output"
	}
	//获取所有文件列表
	tgzFileList, err := basefunc.ReadFileListNew(*source_path)
	if err != nil {
		log.Println(err)
		//***结束***
		return
	}
	//检查目标输出路径
	err = basefunc.CheckDir(targetpathTmp)
	if err != nil {
		log.Println(err)
		//***结束***
		return
	}
	//解压文件到目标路径
	err = basefunc.TarTarget(tgzFileList, *source_path, targetpathTmp)
	if err != nil {
		log.Println(err)
		//***结束***
		return
	}
	// fmt.Println(tgzFileList)
	// fmt.Println(*target_path)
}
