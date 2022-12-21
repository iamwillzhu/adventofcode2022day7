package main

type DirectoryStack []*Directory

// IsEmpty: check if stack is empty
func (s *DirectoryStack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *DirectoryStack) Length() int {
	return len(*s)
}

// Push a new value onto the stack
func (s *DirectoryStack) Push(directory *Directory) {
	*s = append(*s, directory)
}

// return top element of stack. Return false if stack is empty
func (s *DirectoryStack) Top() (*Directory, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	return (*s)[len(*s)-1], true
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *DirectoryStack) Pop() (*Directory, bool) {
	if s.IsEmpty() {
		return nil, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}
