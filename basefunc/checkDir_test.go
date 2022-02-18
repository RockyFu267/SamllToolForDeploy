package basefunc

import (
	"fmt"
	"log"
	"testing"
)

func Test_CheckDir(t *testing.T) {
	res, err := CheckDir("/Users/fuao/Downloads/chart-test/test01")
	if err != nil {
		log.Println("cmd.StdoutPipe: ", err)
		return
	}
	fmt.Println(res)

}

func Test_CreateDir(t *testing.T) {
	err := CreateDir("/Users/fuao/Downloads/chart-test/test01/asdasd")
	if err != nil {
		log.Println("cmd.StdoutPipe: ", err)
		return
	}
}
