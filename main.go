package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 3{
		fmt.Println("exe-link <source> <destination>")
		return
	}

	src,err := filepath.Abs(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	dest,err:= filepath.Abs(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	end:= dest[len(dest)-4:]
	if end != ".exe" {
		dest = dest + ".exe"
	}

	fmt.Println(dest)
	fmt.Println(strings.Repeat(" ",(len(src)+len(dest))/4),"↓")
	fmt.Println(src)

	src = strings.ReplaceAll(src,"\\","\\\\")

	code:=fmt.Sprintf(`
package main
import (
	"os"
	"os/exec"
)
	func main(){
		cmd :=exec.Command("%s")
		cmd.Stdout=os.Stdout
		cmd.Stdin=os.Stdin
		cmd.Stderr=os.Stderr
		cmd.Args = os.Args
		err:=cmd.Run()
		if err != nil {
			println(err.Error())
		}
		os.Exit(0)
}`,src)
	out:="win-link-maker-out.go"

	err=ioutil.WriteFile(out,[]byte(code),666)
	defer os.Remove(out)
	if err != nil {
		fmt.Println(err)
		return
	}

	cmd:=exec.Command("go","build","-ldflags","-w","-o",dest,out)
	cmd.Stdout=os.Stdout
	cmd.Stderr=os.Stderr
	cmd.Stdin=os.Stdin
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

}
