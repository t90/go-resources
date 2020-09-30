//go:generate resource $GOPATH/src/resource/res/res.go $GOPATH/src/resource/res_template.gtp
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"resource/res"
	"strings"
)

type ResourceData struct {
	Name string
	Data string
}

func main(){
	if len(os.Args) < 3 {
		println("USAGE: resource <output filename.go> <input file 1> [input file 2, 3, ...] ")
		return
	}
	inputFiles := os.Args[2:]
	resources := make(chan ResourceData)

	for _, fileName := range inputFiles {
		go encodeResource(fileName, resources)
	}
	count := len(inputFiles)

	f,err := os.Create(os.Args[1])
	defer f.Close()

	if err != nil{
		panic(err)
	}
	var buf bytes.Buffer
	template := string(res.Resources["res_template.gtp"])

	for ;count > 0; count--{
		res := <- resources
		fmt.Fprintf(&buf, "\t\"%s\": {%s},\n", res.Name, res.Data)
	}
	template = strings.Replace(template, "%LINE_ITEMS%", buf.String(),-1)
	f.WriteString(template)
}

func encodeResource(name string, resources chan ResourceData) {
	resData := ResourceData{
		Name: path.Base(name),
	}
	data, err := ioutil.ReadFile(name)
	if err != nil{
		panic(err)
	}
	size := len(data)
	bytes := make([]string,size)
	for i,v := range data{
		bytes[i] = fmt.Sprintf("0x%x", v)
	}
	resData.Data = strings.Join(bytes, ",")
	resources <- resData
}