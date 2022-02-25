package basefunc

import (
	"fmt"
	"log"
	"testing"
)

func Test_DebugTestFunc(t *testing.T) {
	err := DebugTestFunc("/Users/fuao/Desktop/code/github/SamllToolForDeploy/output/chartyaml/asd123.yaml")
	if err != nil {
		log.Println("cmd.StdoutPipe: ", err)
		return
	}
	fmt.Println("ok")
}
