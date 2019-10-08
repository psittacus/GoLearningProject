package main

import (
	"fmt"
	"io/ioutil"
	_ "log"
	"os/exec"
	"strings"
	"time"
)

var (
	path string = "/home/psittacus/Dokumente/Uni/"
	fach string = ""

	yellow string = "\033[1;33m"
	green  string = "\033[0;32m"
	red    string = "\033[0;31m"
	nc     string = "\033[0m"
	blue   string = "\033[0;34m"
	purple string = "\033[0;35m"
	brown  string = "\033[0;33m"
)

func main() {
	input := ""
	fmt.Println("UniGo\n-----")

	fmt.Print("UniGo-$ ")
	fmt.Scanln(&input)
	for input != "x" {
		HandleInput(input)
		fmt.Print("UniGo-$ ")
		fmt.Scanln(&input)
	}
}

func HandleInput(input string) {
	switch input {
	case "n":
		if fach == "" {
			fmt.Println(red + "Bitte Fach (f) eingeben" + nc)
			return
		}
		filename := time.Now().Format("20060102") + ".tex"
		kopie := exec.Command("cp", path+"layout.tex", path+fach+"/"+filename)
		kopie.Run()
		cmd := exec.Command("st", "vim", path+fach+"/"+filename)
		cmd.Run()
	case "ls":
		if fach == "" {
			fmt.Println(red + "Bitte Fach (f) eingeben" + nc)
			return
		}
		listDir(fach)
	case "f":
		fmt.Println("Welches Fach:\n" +
			" - " + green + "(E)nglisch I" + nc + "\n" +
			" - " + blue + "(M)athe I" + nc + "\n" +
			" - " + purple + "(P)rogrammieren I" + nc + "\n" +
			" - " + brown + "(B)WL I" + nc + "\n" +
			" - " + red + "(G)DI I" + nc)
		fmt.Print("UniGo-Fach-$ ")
		fmt.Scanln(&fach)
		switch fach {
		case "E", "e":
			fach = "Englisch1"
		case "M", "m":
			fach = "Mathe1"
		case "P", "p":
			fach = "Prog1"
		case "B", "b":
			fach = "BWL1"
		case "G", "g":
			fach = "GDI1"
		default:
			fmt.Println(red + "Ungültiges Fach!" + nc)
			return
		}
	default:
		listDir(input)
	}
}

func listDir(fach string) {
	files, err := ioutil.ReadDir(path + fach)
	if err != nil {
		fmt.Println(red + "Kein Verzeichnis / Command" + nc)
		return
	}

	for _, f := range files {
		if f.IsDir() {
			fmt.Println(yellow + f.Name() + "/" + nc)
		}
		if strings.HasSuffix(f.Name(), ".tex") {
			fmt.Println(green + f.Name() + nc)
		}
	}
}
