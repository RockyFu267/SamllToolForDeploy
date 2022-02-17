package basecmd

import (
	"log"
	"testing"
)

func Test_CmdAndChangeDirToShow(t *testing.T) {
	params := []string{}
	err := CmdAndChangeDirToShow("/Users/fuao/Downloads/chart-test", "ls", params)
	if err != nil {
		log.Println("cmd.StdoutPipe: ", err)
		return
	}

}
