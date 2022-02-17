package basecmd

import (
	"fmt"
	"log"
	"testing"
)

func Test_CmdAndChangeDirToRes(t *testing.T) {
	params := []string{}
	res, err := CmdAndChangeDirToRes("/Users/fuao/Downloads/chart-test", "ls", params)
	if err != nil {
		log.Println("cmd.StdoutPipe: ", err)
		return
	}
	fmt.Println(res)

}
