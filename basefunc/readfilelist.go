package basefunc

import (
	"SamllToolForDeploy/basecmd"
	"log"
)

//ReadFileList 获取所有文件列表
func ReadFileList(dir string) ([]string, error) {
	params := []string{}
	res, err := basecmd.CmdAndChangeDirToRes(dir, "ls", params)
	if err != nil {
		log.Println(err)
		return res, err
	}
	return res, nil
}
