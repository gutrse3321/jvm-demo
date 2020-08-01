package classpath

import (
	"io/ioutil"
	"path/filepath"
)

/*
	DirEntry：表示目录形式的类路径
 */

type DirEntry struct {
	absDir string //存放目录的绝对路径
}

func newDirEntry(path string) *DirEntry {
	//将参数转换成绝对路径
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &DirEntry{absDir}
}

//将目录和class文件名拼成一个完整的路径，然后调用ioutil包的ReadFile函数读取class文件内容
func (d *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(d.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, d, err
}

func (d *DirEntry) String() string {
	return d.absDir
}
