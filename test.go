package main
import "fmt"

type Object interface {
    typeNum()
}

type Arr struct {
    Object
    val []int
}

type Int struct {
    Object
    val int
}

func main(){
    var stack []Object
    fmt.Println(stack)
    stack = append(stack, Int{val: 34})
    stack = append(stack, Arr{val: []int{1,2,3}})
    fmt.Println(stack)
}
