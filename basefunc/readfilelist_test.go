package basefunc

import (
	"fmt"
	"log"
	"testing"
)

func Test_ReadFileList(t *testing.T) {
	res, err := ReadFileList("/Users/fuao/Downloads/chart-test")
	if err != nil {
		log.Println("cmd.StdoutPipe: ", err)
		return
	}
	fmt.Println(res)

}
