package parse

type Stack struct {
	data []*SQLExpression
}

func NewInsStack() *Stack {
	data := make([]*SQLExpression, 0, 0)

	return &Stack{
		data: data,
	}
}

// Pop implements Stack
func (s *Stack) Pop() *SQLExpression {
	if len(s.data) == 0 {
		return nil
	}

	result := s.data[len(s.data)-1]

	if len(s.data) == 1 {
		s.data = make([]*SQLExpression, 0, 0)
	} else {
		s.data = s.data[:len(s.data)-1]
	}

	return result
}

func (s *Stack) Len() int {
	return len(s.data)
}

func (s *Stack) Push(i *SQLExpression) {
	s.data = append(s.data, i)
}
