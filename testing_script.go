package testing_script

import (
    "fmt"
	"github.com/bitfield/script"
)

func main() {
    fmt.Println("Hello")
	script.Stdin().Column(1).Freq().First(10).Stdout()
}