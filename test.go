package main
import "fmt"

type Object interface {
    typeNum() int
}

type Arr struct {
    Object
    val []int
}

func (n Arr) typeNum() int {
	return 1
}

func (n Int) typeNum() int {
	return 0
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
    fmt.Println(stack[0].typeNum())
    fmt.Println(stack[1].typeNum())
}
