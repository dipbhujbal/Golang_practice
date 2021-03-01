package main
import (
	"fmt"
	"strconv"
)
var i string = "my"
var (
firstname string = "dipali"
lastname string ="bhujbal"
rollnum int = 2

)

func main(){
var j float32=42.3
var i int=int(j)
var z string =string(i)
var z1 string=strconv.Itoa(i)
fmt.Println(j)
fmt.Println(i)

fmt.Println(z)
fmt.Println(z1)
fmt.Println(rollnum)

a:=10
b:= 3
fmt.Println(a & b)
fmt.Println(a | b)
fmt.Println(a ^ b)
fmt.Println(a &^ b)

var num complex64=1+2i
fmt.Println(real(num))
fmt.Println(imag(num))
var num2 complex64=complex(1,2)
fmt.Println(real(num2))
fmt.Println(imag(num2))

//strings
//rune
var r rune='a'
fmt.Printf("%v,%T",r,r)

}