package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	//path:="D:\\Documents\\app\\go\\src\\hugo\\site"
	cmd := exec.Command("powershell" ,"git.exe add .")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out.String())
}
