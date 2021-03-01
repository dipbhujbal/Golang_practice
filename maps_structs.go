package main
import (
	"fmt"
)
func main(){
states:=make(map[string]int)
states = map [string]int{
"MH": 1,
"GJ": 2,
"KN": 3,
}

fmt.Println(states)
delete(states,"GJ")
fmt.Println(states)
}