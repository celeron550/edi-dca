package stack

import "errors"

type ArrayStack struct {
	v []int
	top int
}

func (stack *ArrayStack) Init(size int) {
	stack.v = make([]int, size)
}

func (stack *ArrayStack) Size() int{
	return stack.top
}

func (stack *ArrayStack) Peek() (int, error){
	if stack.top == 0 {
		return -1, errors.New("stack vazia")
	}
	return stack.v[stack.top-1], nil
}

func (stack *ArrayStack) Push(value int) {
	if stack.top == len(stack.v) {
		stack.doubleV()
	}
	stack.v[stack.top] = value
	stack.top++

}

func (stack *ArrayStack) Pop() (int, error){
	if stack.top==0 {
		return -1, errors.New("stack vazia")
	}
	stack.top--
	return stack.v[stack.top], nil
}

func (stack *ArrayStack) doubleV(){
	newV := make([] int, len(stack.v)*2)
	for i := 0; i<stack.Size(); i++{
		newV[i] = stack.v[i]
	}
	stack.v = newV
}

func (stack *ArrayStack) IsEmpty() bool {
	return stack.top == 0
}