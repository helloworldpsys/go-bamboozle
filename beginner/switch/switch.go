package main
import "fmt"
import "time"

 func main(){

	switch time.Now().Weekday(){
	case time.Sunday, time.Saturday:
		fmt.Println("weekend")
		default:
		fmt.Println("workday")
	}
}