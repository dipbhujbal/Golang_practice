//composition is has-a 
//bird has a animal like character
package main
import (
	"fmt"
)
type Animal struct {
Name string
Origin string

}

type Bird struct{
Animal
Speed float32
CanFly bool


}
func main(){
b:=Bird{}
b.Name = "Crow"
b.Origin ="India"
b.Speed = 23
b.CanFly = true
fmt.Println(b)



}