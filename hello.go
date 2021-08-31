package main
import (
	"fmt"

	"github.com/go-programs/greetings"
)

func main()	{
	// Get a greeting message and print it.
	message := greetings.Hello("Illayaraja")
	fmt.Println(message)
}

