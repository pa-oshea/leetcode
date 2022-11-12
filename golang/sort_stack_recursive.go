package main

type Stack struct {
	arr []int
}

func (s *Stack) pop() (r int) {
	r = s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]

	return
}

func (s *Stack) push(el int) {
	s.arr = append(s.arr, el)
}

func (s *Stack) top() int {
	return s.arr[len(s.arr) - 1]
}

func (s *Stack) length() int {
	return len(s.arr)
}

func sort_stack(stack *Stack) *Stack {
	if stack.length() == 1 {
		return stack
	}

	temp := stack.pop()

	sort_stack(stack)

	insert_stack(stack, temp)
	
	return stack
}

func insert_stack(stack *Stack, temp int) *Stack {
	if stack.length() == 0 || temp >= stack.top() {
		stack.push(temp)
		return stack
	}
	return stack
}