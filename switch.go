// 12. Switch

package main 
import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Print("This Go runs on:")

	switch os := runtime.GOOS; os {
	case "darwin" :
		fmt.Println("OS X.")
	case "linux" :
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.", os)
	}

	// cases run from top to bootom stoping and first successfull case

	// switch with no case is the same as switch true
	t := time.Now()
	switch {
	case t.Hour() < 12: 
		fmt.Println("Good morning!")
	case t.Hour() < 17: 
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}


}