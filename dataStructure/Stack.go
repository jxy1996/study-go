package Stack

type (
    Stack struct {
        top    *node
        length int
    }
    node struct {
        val  interface{}
        prev *node
    }
)

func New() *Stack {
    return &Stack{
        top:    nil,
        length: 0,
    }
}

func (s *Stack) Len() int {
    return s.length
}

func (s *Stack) Peek() interface{} {
    if s.length == 0 {
        return nil
    }
    return s.top.val
}

func (s *Stack) Pop() interface{} {
    if s.length == 0 {
        return nil
    }
    n := s.top
    s.top = n.prev
    s.length--
    return n.val
}

func (s *Stack) Push(val interface{}) {
    n := &node{
        val:  val,
        prev: s.top,
    }
    s.top = n
    s.length++
}
