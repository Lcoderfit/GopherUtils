package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

// Error 将命令，参数，错误信息通过格式化输出到标准错误输出
func Error(cmd *cobra.Command, args []string, err error) {
	fmt.Fprintf(os.Stderr, "当前命令: %s, 参数: %v, err: %v\n", cmd, args, err)
}

// convertArgsToFloat64Slice 将传入的参数列表解析为float64类型的切片
func convertArgsToFloat64Slice(args []string, errorHandlingType int) []float64 {
	result := make([]float64, 0, len(args))
	for _, arg := range args {
		// 将参数解析成float64类型
		value, err := strconv.ParseFloat(arg, 64)
		// 除了ExitOnParseError和PanicOnParseError特殊处理外，其他情况均正常返回result
		if err != nil {
			switch errorHandlingType {
			case ExitOnParseError:
				fmt.Fprintf(os.Stderr, "错误的参数: %s\n，err: %v\n", arg, err)
				os.Exit(1)
			case PanicOnParseError:
				panic(err)
			}
		} else {
			result = append(result, value)
		}
	}
	return result
}

func calc(values []float64, opType int) float64 {
	var result float64
	if len(values) == 0 {
		return result
	}

	result = values[0]
	for i := 1; i < len(values); i++ {
		switch opType {
		case ADD:
			result += values[i]
		case MINUS:
			result -= values[i]
		case MULTIPLY:
			result *= values[i]
		case DIVIDE:
			if values[i] == 0 {
				switch dividedByZeroHanding {
				case ReturnOnDivideByZero:
					return result
				case PanicOnDivideByZero:
					panic(errors.New("做除法运算时，参数不能为0"))
				}
			}
			result /= values[i]
		}
	}
	return result
}
