package main 
import "fmt"

// Heap Struct
// Insert Heap
// GetMax Heap

type Heap struct{
	array []int
}
// Insert element inside heap
func(h *Heap)Insert(key int){
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array)-1)
}
// Get maximum element from heap
func(h *Heap) GetMax()int{
	return h.array[0]
}
// Pop maximum element from heap
func(h* Heap) Pop(){
	h.array[0] = h.array[len(h.array)-1]
	h.maxHeapifyDown(0)
}

//maxHeapifyUp
func(h * Heap)maxHeapifyUp(index int){
	for h.array[parent(index)] < h.array[index]{
		h.swap(parent(index), index)
		index = parent(index)
	}
}
//maxHeapifyDown
func(h* Heap)maxHeapifyDown(index int){
	 lastIndex := len(h.array) - 1
	 l,r := left(index), right(index)
	 childToCompare := 0
	 if l > lastIndex{
		return
	 }
	 for l <= lastIndex{
		if l == lastIndex{
			childToCompare = l
		}else if h.array[l] > h.array[r]{
			childToCompare = l
		}else{
			childToCompare = r
		}
		// compare values with larger child, and swap
		if h.array[childToCompare] > h.array[index]{
			h.swap(index, childToCompare)
			index = childToCompare
			l,r = left(index), right(index)
		}else{
			return 
		}
	}

}
// parent
func parent(index int)int{
	return (index-1)/2
}
// left
func left(index int)int{
	return 2*index+1
}
// right
func right(index int)int{
	return 2*index+2
}
// swap
func (h *Heap)swap(i, j int){
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

func main(){
	h := Heap{array: make([]int, 0)}
	h.Insert(5)
	h.Insert(6)
	h.Insert(10)
	h.Insert(8)
	h.Insert(19)
	h.Insert(4)

	fmt.Println(h.array)
	h.Pop()
	fmt.Println(h.array)
	h.Pop()
	fmt.Println(h.array)
	h.Pop()
	fmt.Println(h.array)
	h.Pop()
	fmt.Println(h.array)

}