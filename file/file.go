package file

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var (
	IsDirectoryError     = errors.New("IsDirectoryError: must open a file,not a dir")
	ParamDimensionsError = errors.New("param `dimensions` is invalid")
)

const (
	GrepContain = iota
	GrepNotContain
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
func FindNotExistKeys(path string, keys []string) (result []string, err error) {
	if keys == nil || len(keys) == 0 {
		// 可以使用斜杠对引号进行转义
		// return nil, errors.New("param \"dimensions\" is invalid")
		// 反引号表示原生字符串,不可转义,可包含多行
		// return nil, errors.New(`param "dimensions" is invalid`)
		return nil, ParamDimensionsError
	}
	// 通过grep命令查找文件中是否包含关键字, 如果原始命中中包含空格，则必须分成多个参数传入
	// 如果有多条命令，则必须每个命令使用一个exec
	// bash -c `command string....` command string最好是外层用单引号,内层字符串用引号,
	// 因为如果内层用单引号,Linux命令可能会把字符串识别成命令
	for _, key := range keys {
		cmd := exec.Command(
			"bash", "-c",
			fmt.Sprintf(`grep -q %s_ds %s && echo %d || echo %d`, key, path, GrepContain, GrepNotContain),
		)
		outputByte, err := cmd.Output()
		if err != nil {
			return nil, err
		}
		output := strings.TrimSpace(string(outputByte))
		if output == strconv.Itoa(GrepNotContain) {
			result = append(result, key)
		}
	}
	return result, nil
}
