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
	if c1.Err() != nil {
		t.Fatal()
	}
	if c2, _ := context.WithCancel(context.Background()); fmt.Sprint(c1) != fmt.Sprint(c2) {
		t.Fatal()
	}
	var err = errors.New("custom")
	cancel(err)
	if c1.Err() != err {
		t.Fatal()
	}
}
