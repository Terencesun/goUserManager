package mgr

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"model"
	"os"
	"path/filepath"
)
// 冒泡
func Sort(data *[]int) {
	for i := 0; i < len(*data); i++{
		for j := i; j < len(*data); j++ {
			if (*data)[i] > (*data)[j] {
				(*data)[i], (*data)[j] = (*data)[j], (*data)[i]
			}
		}
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

type Manager struct {
	userInfos map[int]*model.User
	maxId int
	filePath string
}

func (p *Manager) printList()  {
	tmp := make([]int, 0)
	for v := range p.userInfos {
		tmp = append(tmp, v)
	}
	Sort(&tmp)
	fmt.Printf("ID	UserName	Gender		Age	\n")
	for _,v := range tmp {
		fmt.Printf("%v	%v		%v		%v	\n", p.userInfos[v].Id, p.userInfos[v].Username, p.userInfos[v].Gender, p.userInfos[v].Age)
	}
}

// 初始化manager
// 检查本地文件，user是否存在，不存在创建一个新的
// 检查本地缓存文件是否存在，不存在创建一个新的
// 读取user，并加载进userInfos
func (p *Manager) Init() {
	// 初始化
	p.userInfos = make(map[int]*model.User)
	if filePath, err := filepath.Abs("./users"); err != nil {
		panic(model.FILE_PATH_ERROR)
	} else {
		isExist, err := PathExists(filePath)
		if err != nil {
			panic(model.FILE_READ_ERROR)
		}
		p.filePath = filePath
		if isExist {
			// 加载数据
			fmt.Println("存在")
			p.readBuffFile()
		} else {
			// 创建文件
			fmt.Println("不存在")
			p.createBuffFile()
		}
	}
}

func (p *Manager) createBuffFile()  {
	f, err := os.Create(p.filePath)
	defer func() {
		closeErr := f.Close()
		if closeErr != nil {
			panic(model.FILE_CLOSE_ERROR)
		}
	}()
	if err != nil {
		panic(model.FILE_CREATE_ERROR)
	}
	_, writeError := f.Write([]byte(""))
	if writeError != nil {
		panic(model.FILE_CREATE_ERROR)
	}
	p.maxId = 0
}

func (p *Manager) readBuffFile()  {
	file, err := os.Open(p.filePath)
	defer func() {
		fileCloseErr := file.Close()
		if fileCloseErr != nil {
			panic(model.FILE_CLOSE_ERROR)
		}
	}()
	if err != nil {
		panic(model.FILE_OPEN_ERROR)
	}
	content, readError := ioutil.ReadAll(file)
	if readError != nil {
		panic(model.FILE_READ_ERROR)
	}
	if len(content) == 0 {
		// 文件内容为空
		p.maxId = 0
	} else {
		//读取文件，并赋值
		unmarshalErr := json.Unmarshal(content, &p.userInfos)
		if unmarshalErr != nil {
			panic(model.FILE_READ_ERROR)
		}
		p.getMaxId()
	}
}

func (p *Manager) updateBuffFile()  {
	byteData, err := json.Marshal(p.userInfos)
	if err != nil {
		panic(model.UPDATE_BUFF_ERROR)
	}
	buf := new(bytes.Buffer)
	err2 := binary.Write(buf, binary.BigEndian, byteData)
	if err2 != nil {
		panic(model.UPDATE_BUFF_ERROR)
	}
	// 清空文件
	truncateError := os.Truncate(p.filePath, 0)
	if truncateError != nil {
		panic(model.FILE_TRUNC_ERROR)
	}
	f, err3 := os.OpenFile(p.filePath, os.O_WRONLY, os.ModePerm)
	defer func() {
		closeErr := f.Close()
		if closeErr != nil {
			panic(model.FILE_CLOSE_ERROR)
		}
	}()
	if err3 != nil {
		panic(model.FILE_OPEN_ERROR)
	}
	_, err4 := f.Write(buf.Bytes())
	if err4 != nil {
		panic(model.FILE_WRITE_ERROR)
	}
}

func (p *Manager) getMaxId()  {
	for i := range p.userInfos{
		if i > p.maxId {
			p.maxId = i
		}
	}
}

func (p *Manager) compareField(id int, username string, gender string, age int) (fieldKey []string) {
	originalVal, _ := p.userInfos[id]
	if username != "" && username != originalVal.Username {
		fieldKey = append(fieldKey, "Username")
	}
	if gender != "" && gender != originalVal.Gender && (gender == "M" || gender == "F") {
		fieldKey = append(fieldKey, "Gender")
	}
	if age != -1 && age != originalVal.Age {
		fieldKey = append(fieldKey, "Age")
	}
	return
}

func (p *Manager) Create(username string, gender string, age int) {
	id := p.maxId + 1
	obj := &model.User{
		Id: id,
		Username: username,
		Gender: gender,
		Age: age,
	}
	p.userInfos[id] = obj
	p.updateBuffFile()
	p.printList()
}

func (p *Manager) Delete(id int) {
	_, exist := p.userInfos[id]
	if exist {
		delete(p.userInfos, id)
		fmt.Println(p.userInfos)
		p.updateBuffFile()
		p.printList()
	} else {
		fmt.Println("不存在该ID")
	}
}

func (p *Manager) List() {
	p.printList()
}

func (p *Manager) Update(id int, username string, gender string, age int) {
	_, exist := p.userInfos[id]
	if exist {
		fieldKey := p.compareField(id, username, gender, age)
		for _,v := range fieldKey{
			switch v {
			case "Username":
				p.userInfos[id].Username = username
			case "Gender":
				p.userInfos[id].Gender = gender
			case "Age":
				p.userInfos[id].Age = age
			default:
				panic(model.FIELD_TYPE_ERROR)
			}
		}
		p.updateBuffFile()
		p.printList()
	} else {
		fmt.Println("不存在该ID")
	}
}
