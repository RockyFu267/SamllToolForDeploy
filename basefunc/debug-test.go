package basefunc

import (
	"log"
	"os"

	yamlV3 "gopkg.in/yaml.v3"
)

// type StructA struct {
// 	A string `yaml:"a"`
// }

// type StructB struct {
// 	StructA
// 	C StructA
// 	B string `yaml:"b"`
// }

// var data = `
// a: a string from struct A
// b: a string from struct B
// c:
//   a: a string from c.a
// `

var data = `
aaa: aaa
bbb: bbb
ccc: 
  cc1c: 
    cc1c2c: cc1c2c
# ddd: ddd
`

//DebugTestFunc 测试
func DebugTestFunc(path string) error {
	var b TestTest
	x := 22

	err := yamlV3.Unmarshal([]byte(data), &b)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
		return err
	}

	yfile, err := os.Create(path)
	if err != nil {
		log.Println("cannot create data: ", err)
		return err
	}
	defer yfile.Close()

	yencoder := yamlV3.NewEncoder(yfile)
	defer yencoder.Close()
	yencoder.Encode(b)
	yencoder.Encode(b)
	yencoder.Encode(x)
	yencoder.Encode(444444)
	return nil
}
