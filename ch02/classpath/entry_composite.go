package classpath

import (
	"errors"
	"strings"
)

/*
	由更小的Entry组成
*/

type CompositeEntry []Entry

/*
	把路径列表按分隔符分成小路径，然后把每个小路径都转换成具体的Entry实例
*/
func newCompositeEntry(pathList string) CompositeEntry {
	var compositeEntry []Entry

	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}

	return compositeEntry
}

/*
	调用小的entry的readClass方法获取
*/
func (c *CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range *c {
		data, from, err := entry.readClass(className)
		if err != nil {
			return data, from, nil
		}
	}

	return nil, nil, errors.New("没有找到class: " + className)
}

/*
	调用每一个子路径的String()方法，然后用路径分隔符拼接起来
*/
func (c *CompositeEntry) String() string {
	strs := make([]string, len(*c))

	for i, entry := range *c {
		strs[i] = entry.String()
	}

	return strings.Join(strs, pathListSeparator)
}
