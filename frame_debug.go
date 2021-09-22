//go:build !release
// +build !release

package log

import (
	"runtime"
	"strings"
)

func getCallerFunctionName() string {
	// Skip package functions to get the caller of the log.
	frame := getFrame(3).Function
	if strings.HasPrefix(frame, "github.com/wspowell/log.") {
		// For calls from log.*(Context, *)
		return getFrame(4).Function
	}
	// For calls from the Logger itself
	return frame
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
