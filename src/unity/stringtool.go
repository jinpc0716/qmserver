package unity

import (
	"os"
)

//截取字符串
func Substr(str string, start, length int) string {

    rs := []rune(str)
    rl := len(rs)
    end := 0 
    if start < 0 {
        start = rl - 1 + start
    }
    end = start + length
    if start > end {
        start, end = end, start
    }
    if start < 0 {
        start = 0
    }
    if start > rl {
        start = rl
    }
    if end < 0 {
        end = 0
    }
    if end > rl {
        end = rl
    }
    return string(rs[start:end])

}

//判断文件是否存在
func Exist(filename string) bool {

    _, err := os.Stat(filename)
    return err == nil || os.IsExist(err)

}