package basefunc

import (
	"SamllToolForDeploy/basecmd"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	yamlV2 "gopkg.in/yaml.v2"
)

//DryRun 通过helm dryrun 来获取渲染后的yaml
func DryRun(targetpath string) error {
	dirList, err := ReadDirList(targetpath)
	if err != nil {
		log.Println(err)
		return err
	}
	for _, v := range dirList {
		chartNameTmp, err := GetChartName(targetpath + "/" + v + "/" + "Chart.yaml")
		if err != nil {
			continue
		}
		resTmp, err := basecmd.CmdAndChangeDirToResAllInOne(targetpath, "helm install -n testtest "+chartNameTmp+" ./"+v+" --dry-run")
		if err != nil {
			log.Println(err)
			return err
		}
		fmt.Println(resTmp)
	}
	return nil
}

//GetChartName 获取chart包的name，过滤chart.yaml中name为空的包，需要用户手动填写准备
func GetChartName(targetpath string) (string, error) {
	var tmpStruct ChartYaml
	yamlFile, err := ioutil.ReadFile(targetpath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yamlV2.Unmarshal(yamlFile, &tmpStruct)
	if err != nil {
		return "", err
	}
	if len(tmpStruct.Name) == 0 {
		return "", errors.New("chart-name is empty")
	}
	return tmpStruct.Name, nil
}
