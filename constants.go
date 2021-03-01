//constants
package main
import(
	"fmt"
)
//enumerated constant
const e = iota
const (
f=iota
g=iota

)
const (
_=iota
h
i


)
const a int16=23
func main(){
//typed constant 
const myConst int = 42 
const a int=21
fmt.Printf("%v, %T\n",myConst,myConst)
fmt.Println(a)
fmt.Println(e)
fmt.Println(f)
fmt.Println(g)
fmt.Println(h)
}