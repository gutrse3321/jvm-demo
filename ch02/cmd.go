package main

import (
	"flag"
	"fmt"
	"os"
)

/*
	定义的cmd结构
*/
type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	XjreOption  string //bootstrap classpath，启动类路径寻找和加载Java标准库中的类
	class       string
	args        []string
}

/*
	os包定义了一个Args变量，其中存放传递命令行的全部参数。
	如果直接处理os.Args变量，需要写很多代码。Go中内置flag包，可以帮助
	我们处理命令行选项。
*/
func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "打印帮助信息")
	flag.BoolVar(&cmd.helpFlag, "?", false, "打印帮助信息")
	flag.BoolVar(&cmd.versionFlag, "version", false, "打印版本信息")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	//bootstrap classpath
	flag.StringVar(&cmd.XjreOption, "Xjre", "/Library/Java/JavaVirtualMachines/jdk1.8.0_251.jdk/Contents/Home", "jre的路径")
	flag.Parse()

	args := flag.Args()
	cmd.class = "java.lang.Object"
	if len(args) > 0 {
		//cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

/*
	如果调用Parse()函数解析失败，他就调用printUsage()把命令的用法打印到控制台。
	解析成功的话，flag.Args()函数可以捕获其他没有被解析的参数。
	第一个是主类名，剩下的是要传递给主类的参数
*/
func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
