package main
//my logic of defer is it uses stact and pushes the defer code into stack and pops after the //function execution
import (
"fmt"
)
func main(){
fmt.Println("start")
defer  fmt.Println("middle")
fmt.Println("end")
test()


}
func test(){
fmt.Println("test start")
defer  fmt.Println("test middle")
fmt.Println("test end")
a:="before_defer"
defer fmt.Println(a)
a="after_defer"
}