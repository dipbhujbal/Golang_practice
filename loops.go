package main
import (
"fmt"
)
func main(){
for i:=1;i<=5;i++{
fmt.Println(i)

}
s := "hello dips"
// 1 you can have both keys and values printed 
for k,v := range s{

fmt.Println(k,string(v))
}
// 2. You can have only keys to be printed 

for k := range s{

fmt.Println("keys ",k)
}

// you cant have only values to be returned from range function if you want to print only //values then
for _ ,v := range s{
fmt.Println("values",string(v))

}

}