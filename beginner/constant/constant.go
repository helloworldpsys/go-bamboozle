package main
import "fmt"

const age int = 30
var age2 int = 30   // we can't use short hand mehod out of func 
func main(){
	const name string = "potato"

	const (
		port = 5000
		host = "localhost"

	)
	// name string = "saira"           can't change value

	fmt.Println(name, age, age2 , port ,host)
}