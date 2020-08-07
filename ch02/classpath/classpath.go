package classpath

import (
	"os"
	"path/filepath"
)

/*
	Classpath结构体有三个字段，分别存放三种类路径。
*/

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

//parse函数使用-Xjre选项解析启动类路径和扩展类路径，使用-classpath/-cp选项解析用户类路径
func Parse(jreOption, cpOption string) *Classpath {
	cp := new(Classpath)
	//优先使用用户输入的-Xjre选项作为jre目录
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

//ReadClass方法一次从启动类路径、扩展类路径和用户类路径中搜索class文件
func (c *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"

	if data, entry, err := c.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}

	if data, entry, err := c.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}

	return c.userClasspath.readClass(className)
}

func (c *Classpath) String() string {
	return c.bootClasspath.String()
}

//-Xjre选项解析启动类路径和扩展类路径
func (c *Classpath) parseBootAndExtClasspath(jreOption string) {
	//获取jre目录地址
	jreDir := getJreDir(jreOption)

	//jre/lib/*
	//Join连接路径，返回已经clean过的路径
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	c.bootClasspath = newWildcardEntry(jreLibPath)

	//jre/lib/ext/*
	//Join连接路径，返回已经clean过的路径
	jreExtPath := filepath.Join(jreDir, "lib", "*")
	c.extClasspath = newWildcardEntry(jreExtPath)
}

//优先-Xjre选项作为jre目录，如果没有输入，则在当前目录下寻找，再使用JAVA_HOME环境变量
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}

	if exists("./jre") {
		return "./jre"
	}

	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}

	panic("没有找到jre目录")
}

//判断目录是否存在
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

//-classpath/-cp选项解析用户类路径
//如果用户没有提供-classpath/-cp选项，则使用当前目录作为用户类路径
func (c *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}

	c.userClasspath = newEntry(cpOption)
}
