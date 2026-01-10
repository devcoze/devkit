package term

import (
	"fmt"
	"io"
)
import "github.com/moby/term"

// TerminalSize 返回当前终端的宽度和高度（以字符为单位）。
// 如果给定的 writer 不是终端，则返回错误。通常情况下，writer 必须是 os.Stdout。
func TerminalSize(w io.Writer) (width, height int, err error) {
	outFd, isTerminal := term.GetFdInfo(w)
	if !isTerminal {
		return 0, 0, fmt.Errorf("the given writer is not a terminal")
	}
	winSize, err := term.GetWinsize(outFd)
	if err != nil {
		return 0, 0, err
	}
	return int(winSize.Width), int(winSize.Height), nil
}
