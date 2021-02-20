package main
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Glo v0.0.1")
	fmt.Println("Last updated 2-17-2021")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("glo > ")
		text, _ := reader.ReadString('\n')
		fmt.Print(text)
	}
}