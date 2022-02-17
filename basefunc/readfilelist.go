package basefunc

import (
	"SamllToolForDeploy/basecmd"
	"fmt"
	"log"
	"os/exec"
)

func ReadTGZFileList(dir string) ([]string, error) {
	res, err := ReadFileList(dir)
	if err != nil {
		log.Println(err)
		return res, err
	}
	for _, v := range res {
		fmt.Println(len(v))
	}
	return res, nil
}

//ReadFileList 获取所有文件列表
func ReadFileList(dir string) ([]string, error) {
	params := []string{"-l |grep -v \"^d\""}
	res, err := basecmd.CmdAndChangeDirToRes(dir, "ls", params)
	if err != nil {
		log.Println(err)
		return res, err
	}
	return res, nil
}

func TestList() string {
	cmd := "ls -l |grep -v ^d"
	cmd01 := exec.Command("bash", "-c", cmd)
	cmd01.Dir = "/Users/fuao/Downloads/chart-test/test01/test01"
	out, err := cmd01.Output()
	if err != nil {
		return fmt.Sprintf("Failed to execute command: %s", cmd)
	}
	return string(out)
}
