package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var (
		cmd *exec.Cmd
		err error
		output []byte
	)

	cmd = exec.Command("C:\\Users\\84576\\.babun\\cygwin\\bin\\bash.exe", "-c", "echo 1; echo 2; sleep 3; echo 3; ls;")

	// 执行了命令，捕获了子进程的输出（pipe）
	if output, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(output))
}
