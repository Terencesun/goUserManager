package model

import "errors"

var (
	FLAG_ERROR  = errors.New("指令错误")
	PARAMS_ERROR  = errors.New("参数错误")
	COMMAND_ERROR = errors.New("命令错误")
	FILE_PATH_ERROR = errors.New("文件路径错误")
	FILE_READ_ERROR = errors.New("文件读取错误")
	FILE_CREATE_ERROR = errors.New("文件创建失败")
	FILE_CLOSE_ERROR = errors.New("文件关闭失败")
	FILE_OPEN_ERROR = errors.New("文件打开失败")
	FILE_WRITE_ERROR = errors.New("文件写入失败")
	UPDATE_BUFF_ERROR = errors.New("缓冲区更新失败")
	FIELD_TYPE_ERROR = errors.New("字段类型错误")
	FILE_TRUNC_ERROR = errors.New("文件清空错误")
)
