package main

import (
	"SamllToolForDeploy/basefunc"
	"flag"
	"fmt"
	"log"
)

var source_path *string = flag.String("sourcepath", "./", "Use -sourcepath <source_path>")

func main() {
	//获取参数
	flag.Parse()
	fmt.Println("hi buddy")
	tgzFileList, err := basefunc.ReadFileListNew(*source_path)
	if err != nil {
		log.Println(err)
		//***结束***
		return
	}
	fmt.Println(tgzFileList)
}
