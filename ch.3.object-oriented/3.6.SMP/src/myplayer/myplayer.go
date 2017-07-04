package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"mlib"
	"mp"
)

var lib *mlib.MusicManger
var id int = 1
var ctrl1, singal chan int

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.Get(i)
			fmt.Println(i+1, "i", e.Name, e.Artist, e.Source, e.Type)
		}
	case "add":
		if len(tokens) == 6 {
			//这里看来是利用传入的数组新建了一个MusicEntry的变量
			lib.Add(&mlib.MusicEntry{strconv.Itoa(id), tokens[2], tokens[3], tokens[4],
				tokens[5]})
			id++
		} else {
			fmt.Println("USAGE: lib add <name><artist><source><type>")
		}
	case "remove":
		if len(tokens) == 3 {
			if lib.RemoveByName(tokens[2]) != nil {
				//更新id
				id--
			}

		} else {
			fmt.Println("USAGE: lib remove <name>")
		}
	default:
		fmt.Println("Unrecongnized lib conmmand:", tokens[1])

	}
}

func handlePlayCommands(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("USAGE: play <name>")
		return
	}

	_, e := lib.Find(tokens[1])
	if e == nil {
		fmt.Println("the music", tokens[1], "doesn't exist")
		return
	}

	mp.Play(e.Source, e.Type)
}

func main() {
	fmt.Println("Enter following comands to control the player:")
	fmt.Println("lib list -- view the existing music lib")
	//go的转行如何实现
	fmt.Println("lib add <name><artist><source><type> -- Add a music from the lib")
	fmt.Println("lib remove <name> -- Remove the specified music from the lib")
	fmt.Println("play <name> -- play the specified music")

	lib = mlib.NewMusicManger()

	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter command->")
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)
		if line == "q" || line == "e" {
			break
		}
		//将line分割 返回切片数组
		tokens := strings.Split(line, " ")

		if tokens[0] == "lib" && len(tokens) > 1 {
			handleLibCommands(tokens)
		} else if tokens[0] == "play" && len(tokens) > 1 {
			handlePlayCommands(tokens)
		} else {
			fmt.Println("Unrecognized command:", tokens[0])
		}
	}

}
