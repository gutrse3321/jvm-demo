package classpath

import (
	"os"
	"strings"
)

/*
	实现类路径
	启动类路径、扩展类路径、用户类路径
	套用组合模式(composite pattern)实现
*/

//就是冒号 : ，存放路径分隔符
const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	//负责寻找和加载class文件，参数是class文件的相对路径，路径之间用斜线/分隔，文件名有.class后缀
	readClass(className string) ([]byte, Entry, error)
	//相当于toString()，用于返回变量的字符串表示
	String() string
}

//根据参数创建不同类型的Entry实例
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}

	return newDirEntry(path)
}
