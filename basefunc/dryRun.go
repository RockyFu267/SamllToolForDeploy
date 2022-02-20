package basefunc

import (
	"errors"
	"io/ioutil"
	"log"

	yamlV2 "gopkg.in/yaml.v2"
)

func DryRun() error {
	return nil
}

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
		return "", errors.New("name is empty")
	}
	return tmpStruct.Name, nil
}
