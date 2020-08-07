package main

import (
	"fmt"
	"jvm/ch02/classpath"
	"strings"
)

func main() {
	cmd := parseCmd()

	fmt.Println(cmd.XjreOption)

	//if cmd.versionFlag {
	//	fmt.Println("tomo version \"1.8.0_251\"\nTomo(TM) SE Runtime Environment (build 1.8.0_251-b08)\nTomo HotSpot(TM) 64-Bit Server VM (build 25.251-b08, mixed mode)")
	//} else if cmd.helpFlag || cmd.class == "" {
	//	printUsage()
	//} else {
	startJVM(cmd)
	//}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("\nclasspath:%s\nclass:%s args:%v\n", cp, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Println("没有找到或加载主函数", cmd.class)
		return
	}

	fmt.Printf("类数据: %v\n", classData)
}
