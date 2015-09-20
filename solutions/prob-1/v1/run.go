package main

import "os"
import "fmt"
import "io/ioutil"
import "encoding/json"

type File struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type Payload struct {
	Language string `json:"language"`
	Files    []File `json:"files"`
}

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}
	content := string(bytes)
	v := &Payload{
		"cpp",
		[]File{
			File{"a.cpp", content},
			File{"STDIN", "1 2"},
		},
	}
	b, err := json.Marshal(v)
	fmt.Println(string(b))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
