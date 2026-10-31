// 2 types of scope
//local--- declare within function 
//global ---- declare outside all functions can be used everywhere in the same package
package main

import (
	"fmt"
	"net/http"
)
//globbal scope variables be here

// var shortgolang = "helloo "
// 	var canyou = "yo wsp"
// 	var potato = "yuppie"
//     var taskItems = []string{shortgolang, potato, canyou, "potato"}  //this is whole fucking slice yk
func main(){
	
  //var taskms = [4]string{shortgolang, potato, canyou, "potato"} //this becomes ann array why because it has a fixed size yk
	fmt.Println("#### welcome !#### ")
	http.HandleFunc("/", hellouser)



	http.ListenAndServe(":8080",nil)








	fmt.Println("learn go lang ")
	var shortgolang = "helloo "
	var canyou = "yo wsp"
	var potato = "yuppie"
    var taskItems = []string{shortgolang, potato, canyou, "potato"}  //this is whole fucking slice yk
	 
	
	
	printTasks(taskItems)
	printSt()
	fmt.Println()
	taskItems = addTasks(taskItems, "hello my name is claap ")
	printTasks(taskItems)
	}

func hellouser(writer http.ResponseWriter, Request *http.Request){
	var greeting = "hello there"
	fmt.Fprint(writer, greeting)
}

	func printTasks(taskItems []string){
		fmt.Println("learn go lang ")
	 
	//range is the keyword to iterate over elements for different data structures
	for index, task := range taskItems{
		//fmt.Println(index+1,".", task)// toignore index use underscoree
		fmt.Printf( " %d. %s\n", index+1, task)
		}
	}


	func printSt(){
		fmt.Println("hie potato")
	}


	func addTasks(taskItems []string, newTask string) []string {
		var addingnewTasks = append(taskItems, newTask)
		return addingnewTasks

	}

	





