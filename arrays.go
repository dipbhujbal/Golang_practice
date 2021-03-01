package main

import (
	"fmt"
)
func main() {
//ways to declare an array
arr := [3] int {23,45,67}
fmt.Printf("%v\n",arr)

arr2:=[...] int {21,22}
fmt.Printf("%v\n",arr2)


a:=[3] int {44,11,22}
b:=a
b[1]=99
fmt.Printf("%v\n" , a)
fmt.Printf("%v\n",b)

a1:=[3] int {44,11,22}
b1:=&a1
b1[1]=99
fmt.Printf("%v\n" , a1)
fmt.Printf("%v\n",b1)



/////// SLICES
s:=[] int {98,97,96}
fmt.Println(s)
fmt.Println(len(s))
fmt.Println(cap(s))


//a11:=[]int{1,2,3,4,5,6,7}
//b11:=a11[:]
//c11:=a11[2:5]
//make function
a3:=make([]int,4,100)
fmt.Println(a3)
fmt.Println(len(a3))
fmt.Println(cap(a3))
a4:=[] int{}
fmt.Println(a4)
fmt.Println(len(a4))
fmt.Println(cap(a4))
a4=append(a4,2,3)
fmt.Println(a4)
fmt.Printf("len %v\n" ,len(a4))
fmt.Println(cap(a4))
a4=append(a4,21,32,43,22)
fmt.Println(a4)
fmt.Printf("len %v\n" ,len(a4))
fmt.Println(cap(a4))

}