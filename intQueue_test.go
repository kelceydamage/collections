package collections

import (
	"strconv"
	"testing"
)

var iQ IntQueue
var iQueue = iQ.New(6)

// IntQueue testss
//---------------------------------------------------------------------------------------------------- <-100

func TestQPop(t *testing.T) {
	t.Log("Testing: IntSlice for Pop")
	iQueue.Content = IntSlice{1, 4, 2, 5, 6}
	v, ok := iQ.Pop()
	if v.(int) != 1 && ok == false {
		t.Error("Pop() not expected value of 1: " + strconv.Itoa(v.(int)))
	}
	if len(iQ.All()) != 4 {
		t.Error("Pop() not expected value of 4: " + strconv.Itoa(len(iQ.All())))
	}
}

func TestQBuffer(t *testing.T) {
	t.Log("Testing: negative values IntSlice for Max")
	v, ok := iQ.Buffer(5)
	if v.(int) != -1 && ok == false {
		t.Error("Max() not expected value of 1: " + strconv.Itoa(v.(int)))
	}
	if len(iQ.All()) != 5 {
		t.Error("Max() not expected value of 4: " + strconv.Itoa(len(iQ.All())))
	}
}
