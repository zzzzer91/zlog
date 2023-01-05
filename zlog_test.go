package zlog

import (
	"context"
	"github.com/pkg/errors"
	"testing"
)

func f1(x int) error {
	return errors.Errorf("this is an error, x: %d", x)
}

func f2(x int) error {
	return f1(x - 1)
}

func f3(x int) error {
	return f2(x - 1)
}

func TestDefaultLogger(t *testing.T) {
	err := f3(10)
	Ctx(context.Background()).WithError(err).Errorf("failed")
}
