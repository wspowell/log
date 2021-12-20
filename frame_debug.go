//go:build !release
// +build !release

package log

import (
	"runtime"
	"strconv"
	"strings"
)

func getCallerFunctionName() string {
	// Skip package functions to get the caller of the log.
	frame := getFrame(3)
	if strings.HasPrefix(frame.Function, "github.com/wspowell/log.") {
		// For calls from log.*(Context, *)
		frame := getFrame(4)
		return frame.Function + ":" + strconv.Itoa(frame.Line)
	}
	// For calls from the Logger itself
	return frame.Function + ":" + strconv.Itoa(frame.Line)
}

func getFrame(skipFrames int) runtime.Frame {
	// We need the frame at index skipFrames+2, since we never want runtime.Callers and getFrame
	targetFrameIndex := skipFrames + 2

	// Set size to targetFrameIndex+2 to ensure we have room for one more caller than we need
	programCounters := make([]uintptr, targetFrameIndex+2)
	n := runtime.Callers(0, programCounters)

	frame := runtime.Frame{Function: "unknown"}
	if n > 0 {
		frames := runtime.CallersFrames(programCounters[:n])
		for more, frameIndex := true, 0; more && frameIndex <= targetFrameIndex; frameIndex++ {
			var frameCandidate runtime.Frame
			frameCandidate, more = frames.Next()
			if frameIndex == targetFrameIndex {
				frame = frameCandidate
			}
		}
	}

	return frame
}
