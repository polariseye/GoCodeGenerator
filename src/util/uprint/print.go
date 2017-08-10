package uprint

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

var out io.Writer

func init() {
	out = os.Stdout
}

func SetOutput(_out io.Writer) {
	out = _out
}

func Print(args ...interface{}) {
	val := bytes.NewBufferString(fmt.Sprint(args...))
	out.Write(val.Bytes())
}

func Println(args ...interface{}) {
	val := bytes.NewBufferString(fmt.Sprintln(args...))
	out.Write(val.Bytes())
}

func Printf(format string, args ...interface{}) {
	val := bytes.NewBufferString(fmt.Sprintf(format, args...))
	out.Write(val.Bytes())
}
