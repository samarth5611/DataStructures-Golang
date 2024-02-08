package main
import "fmt"

type Queue struct{
	items []int
}
// Enqueue
func (queue *Queue)Enqueue(value int){
	queue.items = append(queue.items, value)
}
//Dequeue
func (queue* Queue)Dequeue(){
	queue.items = queue.items[1:]
}

func(queue* Queue)Print(){
	for _, x := range(queue.items){
		fmt.Print(x , " ")
	}
	fmt.Println()
}


func main(){
	q := Queue{items:make([]int, 0)}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Print()

	q.Dequeue()
	q.Print()


}