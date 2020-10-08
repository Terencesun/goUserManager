package main

import (
	"behave"
	"fmt"
	mgr "manager"
	"model"
)

func main()  {
	if opt, err := behave.GetParams(); err != nil {
		panic(err)
	} else {
		manager := &mgr.Manager{}
		manager.Init()
		switch opt.Command {
		case "create":
			fmt.Println("创建成功")
			manager.Create(opt.Value.Username, opt.Value.Gender, opt.Value.Age)
		case "update":
			fmt.Println("修改成功")
			manager.Update(opt.Value.Id, opt.Value.Username, opt.Value.Gender, opt.Value.Age)
		case "delete":
			fmt.Println("删除成功")
			manager.Delete(opt.Value.Id)
		case "list":
			manager.List()
		default:
			panic(model.COMMAND_ERROR)
		}
	}
}
