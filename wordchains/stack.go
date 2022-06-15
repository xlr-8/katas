package wordchains

// stack is an implementation of a FIFO like queue.
// Elements can be push at the end, while any pop
// will return the first element of the queue.
type stack []string

// pop removes and returns the first element of the stack.
func (s *stack) pop() string {
	if len(*s) > 0 {
		e := (*s)[0]
		*s = (*s)[1:]
		return e
	}
	return ""
}

// push adds and element at the end of the stack.
func (s *stack) push(e string) {
	*s = append(*s, e)
}
