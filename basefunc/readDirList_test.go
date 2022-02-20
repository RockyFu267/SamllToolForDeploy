package basefunc

import (
	"fmt"
	"log"
	"testing"
)

func Test_ReadDirList(t *testing.T) {
	res, err := ReadDirList("/Users/fuao/Downloads/chart-test/test01/test000")
	if err != nil {
		log.Println("cmd.StdoutPipe: ", err)
		return
	}
	fmt.Println(res)
}
