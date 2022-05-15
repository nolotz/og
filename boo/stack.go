package boo

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/nolotz/og/internal/runtimeutil"
)

const (
	errIsNil     = "err is nil"
)

type StackError struct {
	PC    uintptr
	Cause error
}

func (e StackError) Error() string {
	errError := errIsNil
	if e.Cause != nil {
		errError = e.Cause.Error()
	}

	if e.PC == 0 {
		return errError
	}

	f := runtimeutil.FramePC(e.PC)
	if f.File == "" {
		return errError
	}

	var l string
	if f.Function != "" {
		i := strings.LastIndex(f.Function, "/")
		l += f.Function[i+1:] + "/"
	}

	_, file := filepath.Split(f.File)
	l += file + "@" + strconv.Itoa(f.Line)

	return fmt.Sprintf("%s: %s", l, errError)
}

func (e StackError) Unwrap() error {
	return e.Cause
}
