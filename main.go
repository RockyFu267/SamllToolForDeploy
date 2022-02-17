package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

var source_path *string = flag.String("sourcepath", "./", "Use -sourcepath <source_path>")

func main() {
	//获取参数
	flag.Parse()
	fmt.Println("hi buddy")
	// cmd01 := exec.Command("ls", "./")
	// output, err := cmd.Output()
	params := []string{}
	err := CmdAndChangeDirToShow(*source_path, "pwd", params)
	fmt.Println(*source_path)
	if err != nil {

		log.Println("cmd.StdoutPipe: ", err)
		return
	}

}

func CmdAndChangeDirToShow(dir string, commandName string, params []string) error {

	cmd := exec.Command(commandName, params...)
	log.Println("CmdAndChangeDirToFile", dir, cmd.Args)
	//StdoutPipe方法返回一个在命令Start后与命令标准输出关联的管道。Wait方法获知命令结束后会关闭这个管道，一般不需要显式的关闭该管道。
	stdout, err := cmd.StdoutPipe()
	if err != nil {

		log.Println("cmd.StdoutPipe: ", err)
		return err
	}
	cmd.Stderr = os.Stderr
	cmd.Dir = dir
	err = cmd.Start()
	if err != nil {

		return err
	}
	//创建一个流来读取管道内内容，这里逻辑是通过一行一行的读取的
	reader := bufio.NewReader(stdout)
	//实时循环读取输出流中的一行内容
	for {

		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {

			break
		}
		log.Println(line)
	}
	err = cmd.Wait()
	return err
}
