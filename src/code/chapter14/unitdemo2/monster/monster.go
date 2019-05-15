package monster

import(
	"encoding/json"
	"io/ioutil"
	"fmt"
)

type Monster struct{
	Name string
	Age int
	Skill string
}

// 保存文件
func (this *Monster) Store() bool {
	// 序列化
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println("marshal err = ", err)
		return false
	}

	// 保存文件
	filePath := "D:/opt/test/monster.txt"
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil{
		fmt.Println("write err = ", err)
		return false
	}
	return true
}

// 解析文件

func (this *Monster) Restore() bool {
	filePath := "D:/opt/test/monster.txt"
	data, err := ioutil.ReadFile(filePath)
	if err != nil{
		fmt.Println("read err = ", err)
		return false
	}
	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Println("unmarshal err = ", err)
		return false
	}
	fmt.Println(*this)
	return true


}