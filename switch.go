package main
import (
	"fmt"
)
func main(){
//traditional way of switch which cany have overlapping cases 
switch i:=2+3 ; i {
case 1,3,5 :
	fmt.Println("1,3,5")
case 2:
	fmt.Println("2")
default :
	fmt.Println("default")


}

//another way to decalre switch is tagless ,allows oerlapping cases 
i:=10
switch{
case i<=10:
	fmt.Println("yess")
//when you intetionally want to execute next case
	fallthrough
case i<=15:
	fmt.Println("ohh yess")
default:
	fmt.Println("default")





}



}