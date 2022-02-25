package basefunc

import (
	"fmt"
	"log"
	"testing"
)

func Test_DryRun(t *testing.T) {
	err := DryRun("/Users/fuao/Desktop/code/github/SamllToolForDeploy/output/", "testtest")
	if err != nil {
		log.Println("cmd.StdoutPipe: ", err)
		return
	}
	fmt.Println("what")

}
func Test_GetChartName(t *testing.T) {
	res, err := GetChartName("/Users/fuao/Downloads/chart-test/test01/test000/asd01/Chart.yaml")
	if err != nil {
		log.Println("cmd.StdoutPipe: ", err)
		return
	}
	fmt.Println(res)
}

func Test_ReadFileTypeDemo(t *testing.T) {
	err := ReadFileTypeDemo("/Users/fuao/Desktop/code/github/SamllToolForDeploy/output/chartyaml/asd.yaml")
	if err != nil {
		log.Println("cmd.StdoutPipe: ", err)
		return
	}
	fmt.Println("ok")
}
