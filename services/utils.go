package services

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// func inTimeSpan(start, end, check time.Time) bool {
//     return check.After(start) && check.Before(end)
// }

// func main() {
//     start, _ := time.Parse(time.RFC822, "01 Jan 15 10:00 UTC")
//     end, _ := time.Parse(time.RFC822, "01 Jan 16 10:00 UTC")

//     in, _ := time.Parse(time.RFC822, "01 Jan 15 20:00 UTC")
//     out, _ := time.Parse(time.RFC822, "01 Jan 17 10:00 UTC")

//     if inTimeSpan(start, end, in) {
//         fmt.Println(in, "is between", start, "and", end, ".")
//     }

//     if !inTimeSpan(start, end, out) {
//         fmt.Println(out, "is not between", start, "and", end, ".")
//     }
// }

func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("==> Output: %s\n", string(outs))
	}
}
