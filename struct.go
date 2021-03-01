package main
import (
	"fmt"
)
type Doctor struct{
	name string
	spec  string
	num  int
}
func main(){

	aDoc := Doctor {
		name : "vaibhav",
		spec:  "physian",
		num : 23,
}

fmt.Println(aDoc)
//Another way to declare struct
bDoc:=struct{name string}{name : "raj"}
fmt.Println(bDoc)

}