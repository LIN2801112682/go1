package Utils

type Stack struct {
	size int
	top int
	data []interface{} //**
}

func (s *Stack) InitStack() {
	//s := Stack{}
	s.size = 100
	s.top = -1
	s.data = make([]interface{},100)
	//return s
}
func (s *Stack) Size() int {
	len := s.top + 1
	return len
}
func (s *Stack) IsEmpty() bool {
	return s.top == -1
}
func (s *Stack) Push(data interface{}) {
	s.top++
	s.data[s.top] = data
}
func (s *Stack) Pop() interface{} {
	if s.IsEmpty(){
		return nil
	}
	temp := s.data[s.top]
	s.top--
	return temp
}
func (s *Stack) Top() interface{}{
	if s.IsEmpty(){
		return nil
	}
	temp := s.data[s.top]
	return temp
}