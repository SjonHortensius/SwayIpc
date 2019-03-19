package main

import (
	"./swayIpc"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var root swayIpc.Root

	cmd := exec.Command("swaymsg", "-t", "get_tree")
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		panic(err.Error())
	}

	if err := cmd.Start(); err != nil {
		panic(err.Error())
	}

	if err := json.NewDecoder(stdout).Decode(&root); err != nil {
		panic(err.Error())
	}

	if err := cmd.Wait(); err != nil {
		panic(err.Error())
	}

	if false {
		b, err := json.MarshalIndent(root, "", "  ")

		if err != nil {
			panic(err.Error())
		}

		os.Stdout.Write(b)

		fmt.Printf("\n---\n")
	}

	fmt.Printf("* %s [width:%d height:%d]\n", root.Type, root.Rect.Width, root.Rect.Height)
	for _, o := range root.Nodes {
		fmt.Printf("\t* %s: [make:%s model:%s]\n", o.Type, o.Make, o.Model)
		for _, w := range o.Nodes {
			fmt.Printf("\t\t* %s: [name:%s representation:%s]\n", w.Type, w.Name, w.Representation)
			for _, fc := range w.FloatingNodes {
				fmt.Printf("\t\t\t* %s: [name:%s pid:%d]\n", fc.Type, fc.Name, fc.Pid)
			}

			for _, c := range w.Nodes {
				fmt.Printf("\t\t\t* %s: [name:%s pid:%d]\n", c.Type, c.Name, c.Pid)
			}
		}
	}
}
