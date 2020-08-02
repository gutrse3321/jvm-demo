package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

/*
	ZipEntry表示ZIP或JAR文件形式的类路径
*/

type ZipEntry struct {
	absPath string //村黄ZIP或JAR文件的绝对路径
}

func newZipEntry(path string) Entry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &ZipEntry{absPath}
}

/*
	打开zip文件，遍历包，看能否找到class文件，找到
	打开class文件，读取内容，并返回
*/
func (z *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	err := errors.New("没有找到class: " + className)
	r, err := zip.OpenReader(z.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()

	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			//这样defer调用会造成资源无法回收，不是执行一次循环就回收一次，而是循环结束后才开始
			//defer rc.Close()

			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}

			rc.Close()
			return data, z, nil
		}
	}

	return nil, nil, err
}

func (z *ZipEntry) String() string {
	return z.absPath
}
