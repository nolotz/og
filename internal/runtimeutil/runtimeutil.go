package runtimeutil

import "runtime"

// FailingPC returns the program counter from the failing function
// in case of an error by skipping 3 Frames.
// It returns 0 if not enough Frames are available.
func FailingPC() uintptr {
	pc := make([]uintptr, 1)
	if runtime.Callers(3, pc) == 0 {
		return 0
	}

	return pc[0]
}

// FramePC returns a Frame for a given program counter
func FramePC(pc uintptr) *runtime.Frame {
	f, _ := runtime.CallersFrames([]uintptr{pc}).Next()

	return &f
}
