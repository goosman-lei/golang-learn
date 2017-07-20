package stack

import (
    "bytes"
    "strconv"
)

type Stack struct {
    frames []int
    head int
    cap int
}

func NewStack(cap int) *Stack {
    s := &Stack{make([]int, cap), 0, cap}
    return s
}

func (s *Stack) Push(v int) bool {
    if s.head >= s.cap {
        return false
    }

    s.frames[s.head] = v
    s.head ++
    return true
}

func (s *Stack) Pop() int {
    if s.head == 0 {
        return -1
    }

    s.head --
    return s.frames[s.head]
}

func (s *Stack) String() string {
    var buffer bytes.Buffer

    buffer.WriteString("Stack content from head to tail:\n")
    for idx := s.head - 1; idx >= 0; idx -- {
        buffer.WriteString("\t" + strconv.Itoa(idx) + ":" + strconv.Itoa(s.frames[idx]) + "\n")
    }
    buffer.WriteString("Stack content end\n")

    return buffer.String()
}