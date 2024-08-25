/*
Copyright © 2024 PRATIK GAYEN pratik.neel02@gmail.com
*/
package main

import (
	"os"

	"todo/cmd"
)

func main() {
	home, _ := os.UserHomeDir()
	os.Chdir(home + "/Documents")
	cmd.Execute()
}
