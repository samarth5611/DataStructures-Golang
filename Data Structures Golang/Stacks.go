package main
import "fmt"


type Stack struct{
	items []int

}
//BuildStack
func BuildStack()Stack{
	return Stack{items: make([]int, 0)}
}
//Push
func (stack *Stack)Push(value int){
	stack.items = append(stack.items , value)
}
//Pop
func(stack *Stack)Pop(){
	if len(stack.items) == 0{
		fmt.Println("Can'pop, stack is already empty!")
	}
	l := len(stack.items)
	stack.items = stack.items[:l-1]
}
//Top
func(stack *Stack)Top()int{
	if len(stack.items) == 0{
		fmt.Println("Stack is already empty!")
	}
	return stack.items[len(stack.items)-1]
}
//Size
func (stack*Stack)Size()int{
	return len(stack.items)
}
//Print Stack
func (stack *Stack)Print(){
	for i:=0 ; i< len(stack.items); i++{
		fmt.Print(stack.items[i], " ")
	}
	fmt.Println("")
}

func main(){
	mystack := BuildStack()
	mystack.Push(11)
	mystack.Push(23)
	mystack.Push(56)
	mystack.Push(98)

	mystack.Print()

	mystack.Pop()

	mystack.Print()
}