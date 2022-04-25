package ctx

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

func TestWithCancelX(t *testing.T) { XTestWithCancelX(t) }

func XTestWithCancelX(t *testing.T) {
	c1, cancel := WithCancel(context.Background())
	defer cancel(nil)
	if c1.Err() != nil {
		t.Fatal()
	}
	c2, cancel2 := context.WithCancel(context.Background())
	defer cancel2()
	if fmt.Sprint(c1) != fmt.Sprint(c2) {
		t.Fatal()
	}
	var err = errors.New("custom")
	cancel(err)
	if c1.Err() != err {
		t.Fatal()
	}
}
