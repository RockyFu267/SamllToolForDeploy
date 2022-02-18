package basefunc

import (
	"log"
	"os"
)

func CheckDir(path string) (bool, error) {
	taget, err := os.Stat(path)
	if err != nil {
		log.Println(err)
		return false, err
	}
	return taget.IsDir(), nil
}

//CreateDir path是绝对路径
func CreateDir(path string) error {
	err := os.Mkdir(path, 0777)
	if err != nil {
		log.Println(err)
		return err
	}
	//如果不赋权限会有问题
	os.Chmod("path", 0777)
	return nil
}
