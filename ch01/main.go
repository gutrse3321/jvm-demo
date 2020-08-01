package main

import "fmt"

func main() {
	cmd := parseCmd()

	if cmd.versionFlag {
		fmt.Println("tomo version \"1.8.0_251\"\nTomo(TM) SE Runtime Environment (build 1.8.0_251-b08)\nTomo HotSpot(TM) 64-Bit Server VM (build 25.251-b08, mixed mode)")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	fmt.Printf("classpath:%s class:%s args:%v\n", cmd.cpOption, cmd.class, cmd.args)
}
