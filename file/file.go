package file

import (
	"errors"
	"os"
	"os/exec"
)

var (
	IsDirectoryError     = errors.New("IsDirectoryError: must open a file,not a dir")
	ParamDimensionsError = errors.New("param `dimensions` is invalid")
)

// Open 打开一个文件，如果打开时报错或者读取的是一个目录而不是文件，均会返回error
func Open(path string) (*os.File, error) {
	// 读取文件信息
	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	// 读取的文件不能为目录
	if fi.IsDir() {
		return nil, IsDirectoryError
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// FindNotExistKeys 在文件中查找不存在于维度列表中的维度
func FindNotExistKeys(path string, keys []string) (string, error) {
	if keys == nil || len(keys) == 0 {
		// 可以使用斜杠对引号进行转义
		// return nil, errors.New("param \"dimensions\" is invalid")
		// 反引号表示原生字符串,不可转义,可包含多行
		// return nil, errors.New(`param "dimensions" is invalid`)
		return "", ParamDimensionsError
	}
	// 通过grep命令查找文件中是否包含关键字
	cmd := exec.Command("grep", "d", "/home/learnGoroutine/file/src.log", "&&", "echo", "contain")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
