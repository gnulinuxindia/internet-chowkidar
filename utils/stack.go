package utils

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/go-errors/errors"
)

func GetStack(err *errors.Error, colors bool) string {
	basePath := GetBasePath() + "/"
	stack := ""

	boldBlack := color.New(color.FgHiBlack, color.Bold)
	blue := color.New(color.FgBlue)
	black := color.New(color.FgBlack)
	cyan := color.New(color.FgCyan)

	if !colors {
		boldBlack.DisableColor()
		blue.DisableColor()
		black.DisableColor()
		cyan.DisableColor()
	}

	allFrames := err.StackFrames()
	frames := make([]errors.StackFrame, 0, len(allFrames))
	for _, frame := range allFrames {
		if !strings.Contains(frame.Package, "runtime") {
			frames = append(frames, frame)
		}
	}
	frameCount := len(frames)

	for i, frame := range frames {
		var file string
		if strings.Contains(frame.File, basePath) {
			file = strings.Replace(frame.File, basePath, "", 1)
		} else {
			file = black.Sprint(frame.File)
		}

		frameStr := fmt.Sprintf("%s%s%d", file, blue.Sprint(":"), frame.LineNumber)

		source, err := frame.SourceLine()
		stack += boldBlack.Sprint(fmt.Sprintf("%2d: ", frameCount-i))
		if err != nil {
			stack += fmt.Sprintf("%s\n", frameStr)
		} else {
			var arrow string
			var fmtdSource string
			if strings.Contains(frame.File, basePath) {
				arrow = cyan.Sprint(">> ")
				fmtdSource = source
			} else {
				arrow = black.Sprint("> ")
				fmtdSource = black.Sprint(source)
			}

			stack += fmt.Sprintf("%s\n    %s%s\n", frameStr, arrow, fmtdSource)
		}
	}

	return " " + err.Error() + "\n" + stack
}
