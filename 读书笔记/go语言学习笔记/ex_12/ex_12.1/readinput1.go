package main
import "fmt"

var (
	firstName, lastName, s string
	i int
	f float32
	input = "56.12 / 5212 / Go"
	format = "%f / %d / %s"
)
func main() {
	fmt.Println("Please enter your full name:")
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("%s %s\n", firstName, lastName)
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)
}