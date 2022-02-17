package main

import (
	"SamllToolForDeploy/basecmd"
	"flag"
	"fmt"
	"log"
)

var source_path *string = flag.String("sourcepath", "./", "Use -sourcepath <source_path>")

func main() {
	//获取参数
	flag.Parse()
	fmt.Println("hi buddy")
	params := []string{}
	err := basecmd.CmdAndChangeDirToShow(*source_path, "ls", params)
	fmt.Println(*source_path)
	if err != nil {

		log.Println("cmd.StdoutPipe: ", err)
		return
	}

}
