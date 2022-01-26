package main
import "fmt"


func multiply(res int) int {
	fmt.Println("Result: ", res)
	return 0
}
func show() {
	fmt.Println("show")
}
func tellYourName(){
	fmt.Println("tell name")
}
func main() {
	multiply(1)
	defer multiply(2)
	defer multiply(3)
	show()
	tellYourName()
	
}