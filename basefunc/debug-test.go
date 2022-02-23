package basefunc

import (
	"log"
	"os"

	yamlV3 "gopkg.in/yaml.v3"
)

type StructA struct {
	A string `yaml:"a"`
}

type StructB struct {
	StructA
	C StructA
	B string `yaml:"b"`
}

var data = `
a: a string from struct A
b: a string from struct B
c:
  a: a string from c.a
`

func DebugTestFunc() {
	var b StructB
	x := 22

	err := yamlV3.Unmarshal([]byte(data), &b)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}

	yfile, _ := os.Create("test.yaml")
	defer yfile.Close()

	yencoder := yamlV3.NewEncoder(yfile)
	defer yencoder.Close()
	yencoder.Encode(b)
	yencoder.Encode(b)
	yencoder.Encode(x)
	yencoder.Encode(44)
}
