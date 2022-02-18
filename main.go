package main

import (
	"SamllToolForDeploy/basefunc"
	"flag"
	"fmt"
	"log"
)

var source_path *string = flag.String("sourcepath", "./", "Use -sourcepath <source_path>")
var target_path *string = flag.String("targetpath", "./output", "Use -targetpath <source_path>")

func main() {
	//获取参数
	flag.Parse()
	fmt.Println("hi buddy")
	//获取所有文件列表
	tgzFileList, err := basefunc.ReadFileListNew(*source_path)
	if err != nil {
		log.Println(err)
		//***结束***
		return
	}
	//检查目标输出路径
	err = basefunc.CheckDir(*target_path)
	if err != nil {
		log.Println(err)
		//***结束***
		return
	}
	fmt.Println(tgzFileList)
	fmt.Println(*target_path)
}
