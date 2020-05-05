package monster

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type Monster struct {
	Name string
	Age int
	Skill string
}

func (this *Monster) Store() bool {
	// 序列化
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println("marshal err = ", err)
		return false
	}

	// 保存到文件
	filePath := "d:/monster.sef"
	ioutil.WriteFile(filePath, data, 0666)

	if err != nil {
		fmt.Println("write file err = ", err)
		return false
	}

	return true
}

func (this *Monster) ReStore() bool {

	// 先从文件中，读取序列化的字符串
	filePath := "d:/monster.sef"

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("ReadFile err = ", err)
		return false
	}

	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Println("UnMarshal err = ", err)
		return false
	}

	return true


}
