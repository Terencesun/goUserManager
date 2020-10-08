package behave

import (
	"flag"
	"model"
)

type Behave struct {
	Command string
	Value *model.User
}

// 解析命令
// 指令
// -h 帮助
// -c 创建新用户
// -d 删除用户
// -l 列出条件用户
// -u 更新用户
func GetParams() (behave *Behave, err error) {
	var (
		c bool
		d bool
		l bool
		u bool
		id int
		username string
		gender string
		age int
	)
	flag.BoolVar(&c, "create", false, "create new user")
	flag.BoolVar(&d, "delete", false, "delet a user")
	flag.BoolVar(&l, "list", false, "list the user")
	flag.BoolVar(&u, "update", false, "update a user")
	flag.IntVar(&id, "id", -1, "the params of user id")
	flag.StringVar(&username, "username", "", "the params of user username")
	flag.StringVar(&gender, "gender", "", "the params of user gender")
	flag.IntVar(&age, "age", -1, "the params of user age")
	flag.Parse()
	switch {
	case c:
		// 用户名，性别，年龄
		if username == "" {
			err = model.PARAMS_ERROR
			return
		}
		if gender == "" || !(gender == "M" || gender == "F") {
			err = model.PARAMS_ERROR
			return
		}
		if age == -1 {
			err = model.PARAMS_ERROR
			return
		}
		tmp := &model.User{
			Id: 0,
			Username: username,
			Gender: gender,
			Age: age,
		}
		behave = &Behave{
			Command: "create",
			Value: tmp,
		}
	case d:
		// id
		if id == -1 {
			err = model.PARAMS_ERROR
			return
		}
		tmp := &model.User{
			Id: id,
		}
		behave = &Behave{
			Command: "delete",
			Value: tmp,
		}
	case l:
		behave = &Behave{
			Command: "list",
			Value: nil,
		}
	case u:
		// id和其他值
		if username == "" {
			err = model.PARAMS_ERROR
			return
		}
		if gender == "" || !(gender == "M" || gender == "F") {
			err = model.PARAMS_ERROR
			return
		}
		if id == -1 {
			err = model.PARAMS_ERROR
			return
		}
		if age == -1 {
			err = model.PARAMS_ERROR
			return
		}
		tmp := &model.User{
			Id: id,
			Username: username,
			Gender: gender,
			Age: age,
		}
		behave = &Behave{
			Command: "update",
			Value: tmp,
		}
	default:
		err = model.FLAG_ERROR
	}
	return
}
