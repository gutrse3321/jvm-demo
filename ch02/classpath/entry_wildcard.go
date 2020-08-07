package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

/*
	* 通配符入口
	WildcardEntry实际上也是CompositeEntry，不需要再定义新的类型
*/

func newWildcardEntry(path string) *CompositeEntry {
	//去掉路径通配符*
	baseDir := path[:len(path)-1] //删除通配符 *
	var compositeEntry CompositeEntry

	//调用Walk函数遍历baseDir创建ZipEntry
	filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}

		//根据后缀名选出JAR文件，并返回SkipDir跳过子目录(通配符类路径不能递归匹配子目录下的JAR文件)
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}

		return nil
	})

	return &compositeEntry
}
