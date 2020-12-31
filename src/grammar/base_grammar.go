package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filePath := "/Users/jhs/go/go_learning/src/grammar/test.txt"
	printFileContentV1(filePath)
	printFileContentV2(filePath)

}

//if else在go中的简单使用
func printFileContentV1(fileName string) {
	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

//在go中if判断中也可以像java for一样使用
func printFileContentV2(fileName string) {
	if contents, err := ioutil.ReadFile(fileName); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}
