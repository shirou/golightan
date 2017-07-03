package lexer

type Stack struct {
	stack []int
}

const (
	StackEmpty = -1
)

func NewStack(capacity int) *Stack {
	return &Stack{
		stack: make([]int, 0, capacity),
	}
}

func (s *Stack) Push(n int) {
	s.stack = append(s.stack, n)
}
func (s *Stack) Pop() int {
	if s.Len() == 0 {
		return StackEmpty
	}
	l := s.Len()
	value := s.stack[l-1]
	s.stack = s.stack[:l-1]
	return value
}

func (s *Stack) Last() int {
	return s.stack[s.Len()-1]
}

func (s *Stack) Len() int {
	return len(s.stack)
}
