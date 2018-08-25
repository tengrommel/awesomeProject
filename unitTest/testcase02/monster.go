package testcase02

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

// 给Monster绑定方法Store，可以将一个Monster变量（对象），序列化后保存到文件中
func (this *Monster) Store() bool {
	// 先序列化
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println("marshal err =", err)
		return false
	}
	// 保存到文件
	filePath := "/Users/tengzhou/go/src/awesomeProject/unitTest/testingdemo01/testcase01/Monster.ser"
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Println("write file err =", err)
		return false
	}
	return true
}

func (this *Monster) ReStore() bool {
	filePath := "/Users/tengzhou/go/src/awesomeProject/unitTest/testingdemo01/testcase01/Monster.ser"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("ReadFile err =", err)
		return false
	}
	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Println("UnMarshal err =", err)
	}
	return false
}
