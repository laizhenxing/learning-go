// Dup1 prints the text of echo line that appears more than
// once in the standard input, preceded by its count
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		// 控制退出
		if input.Text() == "end" {
			break
		}
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}
