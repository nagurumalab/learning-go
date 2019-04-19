package stack

//Borrowed this code from this generous person:  https://gist.github.com/marz619/a0e6d3884a4d0d271ddb89ab76368690
import (
	"reflect"
	"sync"
)

// Stack implements Push & Pop
type Stack interface {
	Push(interface{})
	Pop() interface{}
	Size() int
}

// stack is our internal implementation of a Stack. It is concurrent and type
// safe
type stack struct {
	l *sync.Mutex
	// slice to hold the data
	s   []interface{}
	pos int          // current position
	t   reflect.Type // current type
}

func (s stack) pushTypeCheck(v interface{}) {
	typ := reflect.TypeOf(v)
	if s.t != nil && typ.PkgPath()+"#"+typ.Name() != s.t.PkgPath()+"#"+s.t.Name() {
		panic("pushing different type onto stack")
	}
}

func (s stack) setType() {
	switch s.pos {
	case 0:
		s.t = nil
	case 1:
		s.t = reflect.TypeOf(s.s[0])
	}
}

// Pop implements Stack. It panics if the stack size is 0
func (s stack) Pop() (v interface{}) {
	if len(s.s) == 0 {
		panic("cannot pop empty stack")
	}
	// lock/unlock
	s.l.Lock()
	defer s.l.Unlock()
	defer s.setType()
	// pop
	s.pos--
	v, s.s = s.s[s.pos], s.s[:s.pos]
	return
}

// Push implements Stack. It panics if v is nil or does not match the type
// of the value currently on the stack
func (s stack) Push(v interface{}) {
	if v == nil {
		panic("cannot push nil onto stack")
	}
	// lock/unlock
	s.l.Lock()
	defer s.l.Unlock()
	// type check
	s.pushTypeCheck(v)
	// type set
	defer s.setType()
	// push
	s.s = append(s.s, v)
	s.pos++
}

// Size implements Seq
func (s stack) Size() int {
	s.l.Lock()
	defer s.l.Unlock()
	// return the size
	return len(s.s)
}

// NewStack returns a Stack
func NewStack() Stack {
	return stack{
		l: new(sync.Mutex),
		s: make([]interface{}, 1),
	}
}
