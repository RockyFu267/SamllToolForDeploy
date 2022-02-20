package basefunc

import (
	"fmt"
	"log"
	"testing"
)

func Test_DryRun(t *testing.T) {
	err := DryRun("/Users/fuao/Downloads/chart-test/test01/test000/")
	if err != nil {
		log.Println("cmd.StdoutPipe: ", err)
		return
	}

}
func Test_GetChartName(t *testing.T) {
	res, err := GetChartName("/Users/fuao/Downloads/chart-test/test01/test000/asd01/Chart.yaml")
	if err != nil {
		log.Println("cmd.StdoutPipe: ", err)
		return
	}
	fmt.Println(res)
}
