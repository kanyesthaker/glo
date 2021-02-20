package main
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Malgo v0.0.1")
	fmt.Println("Last updated 2-17-2021")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("glo > ")
		s, _ := reader.ReadString('\n')
		fmt.Print(rep(s))
	}
}

func READ(str string) string {
	return str
}

func EVAL(ast string, env string) string {
	return ast
}

func PRINT(str string) string {
	return str
}

func rep(str string) string {
	return PRINT(EVAL(READ(str), ""))
}