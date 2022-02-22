package basefunc

import (
	"SamllToolForDeploy/basecmd"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	yamlV2 "gopkg.in/yaml.v2"
)

//DryRun 通过helm dryrun 来获取渲染后的yaml
func DryRun(targetpath string, namespace string) error {
	dirList, err := ReadDirList(targetpath)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("step1")
	for _, v := range dirList {
		chartNameTmp, err := GetChartName(targetpath + "/" + v + "/" + "Chart.yaml")
		if err != nil {
			continue
		}
		fmt.Println("step2")
		//***使用debug***
		//err = basecmd.CmdAndChangeDirToResFile(targetpath+"/chartyaml/"+chartNameTmp+".yaml", targetpath, "helm install -n testtest "+chartNameTmp+" ./"+v+" --dry-run")
		//***使用template***
		err = basecmd.CmdAndChangeDirToResFile(targetpath+"/chartyaml/"+chartNameTmp+".yaml", targetpath, "helm template -n "+namespace+" "+chartNameTmp+" ./"+v)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println("step3")
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

func ReadFileTypeDemo(targetpath string) error {
	// var tmpStruct []map[string]interface{}
	var tmpStruct TestTest
	yamlFile, err := ioutil.ReadFile(targetpath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yamlV2.Unmarshal(yamlFile, &tmpStruct)
	if err != nil {
		return err
	}
	// result := []string{}
	// s := string(yamlFile)
	// for _, lineStr := range strings.Split(s, "\n") {
	// 	lineStr = strings.TrimSpace(lineStr)
	// 	if lineStr == "" {
	// 		continue
	// 	}
	// 	result = append(result, lineStr)
	// }
	// for _, v := range result {
	// 	fmt.Println(v)
	// }
	tmpStruct.Ccc.Cc1c.Cc1c2c = "fuckyou"
	A, err := yamlV2.Marshal(tmpStruct)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println(tmpStruct, "11111")
	res2A, _ := json.Marshal(tmpStruct)
	fmt.Println(string(res2A), "22222")
	fmt.Println(string(A), "33333")
	return nil
}
