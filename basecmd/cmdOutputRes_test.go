package basecmd

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

func Test_CmdAndChangeDirToRes(t *testing.T) {
	params := []string{}
	res, err := CmdAndChangeDirToRes("/Users/fuao/Downloads/chart-test/test01", "ls", params)
	if err != nil {
		log.Println("cmd.StdoutPipe: ", err)
		return
	}
	fmt.Println(res[0], len(res[0]))
	str := strings.Replace(res[0], "\n", "", -1)
	fmt.Println(str, len(str))

}

func Test_CmdAndChangeDirToResAllInOne(t *testing.T) {
	res, err := CmdAndChangeDirToResAllInOne("/Users/fuao/Downloads/chart-test/test01", "ls")
	if err != nil {
		log.Println("cmd.StdoutPipe: ", err)
		return
	}
	fmt.Println(res[0], len(res[0]))
	str := strings.Replace(res[0], "\n", "", -1)
	fmt.Println(str, len(str))

}
