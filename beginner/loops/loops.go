package main
import "fmt"
//for --> only construct in go for looping
func main(){
	//we only use for loop in go we can't use any other loop in go
	// we can't use i as var in for loop
	// we need to use the gofer syntax
// for as a while loop

i := 1
for i <= 3 {
	fmt.Println(i)
	i = i + 1
}

//infinite loop
/* for {
fmt.Println("1")
}*/


// classic for loop

for i = 0; i<3; i++{
	fmt.Println(i)
}

// for i:= range 20 {
// 	fmt.Println(i)
// }


//continue = skip this particular turn and jump to the next one



// break = breaks the whole loop


for value := range 20 {

		if value %2 == 1 {
			continue  //now this gonna skip odd numbers fair
		}

		println(value)
}}