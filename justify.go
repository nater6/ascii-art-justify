package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func PrintStr(a string, b string, c string, d string, e string, f string, g string, h string) {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run . [STRING] [BANNER] [OPTION]")
		fmt.Println("EX: go run . something standard --align=right")
		return
	}

	file, err := os.Open(os.Args[2] + ".txt")
	if err != nil {
		fmt.Println("Usage: go run . [STRING] [BANNER] [OPTION]")
		fmt.Print("EX: go run . something standard --align=right")
		return
	}
	// file to slice of string by line

	var lttrlines []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lttrlines = append(lttrlines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	// find a string in a scanned file
	if len(os.Args) != 4 || len(os.Args[3]) < 9 {
		fmt.Println("Usage: go run . [STRING] [BANNER] [OPTION]")
		fmt.Print("EX: go run . something standard --align=right")
		return
	}

	//Get the align command
	loc := os.Args[3]
	loc = loc[8:]

	if loc != "right" && loc != "left" && loc != "center" && loc != "justify" {
		fmt.Println("Usage: go run . [STRING] [BANNER] [OPTION]")
		fmt.Print("EX: go run . something standard --align=right")
		return
	}
	// Get terminal lenght
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()

	if err != nil {
		log.Fatal(err.Error())
	}
	//Remove white spaces to get int from []byte
	T := strings.Fields(strings.TrimSpace(string(out)))
	wid := T[1]
	w, _ := strconv.Atoi(wid)

	//Strings
	ln1 := ""
	ln2 := ""
	ln3 := ""
	ln4 := ""
	ln5 := ""
	ln6 := ""
	ln7 := ""
	ln8 := ""

	//Stirng Slices
	sln1 := []string{}
	sln2 := []string{}
	sln3 := []string{}
	sln4 := []string{}
	sln5 := []string{}
	sln6 := []string{}
	sln7 := []string{}
	sln8 := []string{}

	if loc == "center" || loc == "right" {
		// Splits on newlines by default.

		line := 0
		Output := os.Args[1]

		if strings.Contains("\"", Output) {
			s := []rune(Output)
			emptys := []rune{}
			for _, r := range s {
				if r == '"' {
					emptys = append(emptys, '\\')
					emptys = append(emptys, r)
				} else {
					emptys = append(emptys, r)
				}
			}
			Output = string(emptys)
		}

		SlcOutput := []rune(Output)
		dash := strings.Index(Output, `\n`)

		for i := 0; i < len(SlcOutput); i++ {
			if i == dash && dash >= 0 {
				if ln1 != "" && ln2 != "" && ln3 != "" && ln4 != "" && ln5 != "" && ln6 != "" && ln7 != "" && ln8 != "" {
					if loc == "center" {
						CenterPrinter(ln1, w)
						CenterPrinter(ln2, w)
						CenterPrinter(ln3, w)
						CenterPrinter(ln4, w)
						CenterPrinter(ln5, w)
						CenterPrinter(ln6, w)
						CenterPrinter(ln7, w)
						CenterPrinter(ln8, w)
					} else if loc == "right" {
						RightPrinter(ln1, w)
						RightPrinter(ln2, w)
						RightPrinter(ln3, w)
						RightPrinter(ln4, w)
						RightPrinter(ln5, w)
						RightPrinter(ln6, w)
						RightPrinter(ln7, w)
						RightPrinter(ln8, w)
					}

				} else {
					fmt.Println()
				}

				ln1 = ""
				ln2 = ""
				ln3 = ""
				ln4 = ""
				ln5 = ""
				ln6 = ""
				ln7 = ""
				ln8 = ""

				dash = dash + strings.Index(Output[dash+1:], `\n`) + 1
				i++
			} else {
				for j, s := range lttrlines {
					line = j
					if len(s) > 0 && s == string(SlcOutput[i]) {
						break
					}

				}

				for i := line + 1; i <= line+8; i++ {
					if i == line+1 {
						ln1 = ln1 + lttrlines[i]
					}
					if i == line+2 {
						ln2 = ln2 + lttrlines[i]
					}
					if i == line+3 {
						ln3 = ln3 + lttrlines[i]
					}
					if i == line+4 {
						ln4 = ln4 + lttrlines[i]
					}
					if i == line+5 {
						ln5 = ln5 + lttrlines[i]
					}
					if i == line+6 {
						ln6 = ln6 + lttrlines[i]
					}
					if i == line+7 {
						ln7 = ln7 + lttrlines[i]
					}
					if i == line+8 {
						ln8 = ln8 + lttrlines[i]
					}

				}
				if i == len(Output)-1 {
					if loc == "center" {
						CenterPrinter(ln1, w)
						CenterPrinter(ln2, w)
						CenterPrinter(ln3, w)
						CenterPrinter(ln4, w)
						CenterPrinter(ln5, w)
						CenterPrinter(ln6, w)
						CenterPrinter(ln7, w)
						CenterPrinter(ln8, w)
					} else if loc == "right" {
						RightPrinter(ln1, w)
						RightPrinter(ln2, w)
						RightPrinter(ln3, w)
						RightPrinter(ln4, w)
						RightPrinter(ln5, w)
						RightPrinter(ln6, w)
						RightPrinter(ln7, w)
						RightPrinter(ln8, w)
					}
				}
			}
		}
	} else if loc == "justify" {

		line := 0
		Output := os.Args[1]
		SlcOutput := []rune(Output)
		dash := strings.Index(Output, `\n`)

		for i := 0; i < len(SlcOutput); i++ {
			if (i == dash) && (dash >= 0) {
				fmt.Println("!")
				if (ln1 == "" || ln2 != "" || ln3 != "" || ln4 != "" || ln5 != "" || ln6 != "" ||ln7 != "" || ln8 != "") {

					sln1 = append(sln1, ln1)
					sln2 = append(sln2, ln2)
					sln3 = append(sln3, ln3)
					sln4 = append(sln4, ln4)
					sln5 = append(sln5, ln5)
					sln6 = append(sln6, ln6)
					sln7 = append(sln7, ln7)
					sln8 = append(sln8, ln8)
					ln1 = ""
					ln2 = ""
					ln3 = ""
					ln4 = ""
					ln5 = ""
					ln6 = ""
					ln7 = ""
					ln8 = ""

					Printer(sln1, w)
					Printer(sln2, w)
					Printer(sln3, w)
					Printer(sln4, w)
					Printer(sln5, w)
					Printer(sln6, w)
					Printer(sln7, w)
					Printer(sln8, w)

				} else {
					fmt.Println("EMpty Line")
				}

				sln1 = []string{}
				sln2 = []string{}
				sln3 = []string{}
				sln4 = []string{}
				sln5 = []string{}
				sln6 = []string{}
				sln7 = []string{}
				sln8 = []string{}
				dash = dash + strings.Index(Output[dash+1:], `\n`) + 1
				i++
			} else {
				for j, s := range lttrlines {
					line = j
					if len(s) > 0 && s == string(SlcOutput[i]) {
						break
					}

				}
				if lttrlines[line] != " " {
					for i := line + 1; i <= line+8; i++ {
						if i == line+1 {
							ln1 = ln1 + lttrlines[i]
						}
						if i == line+2 {
							ln2 = ln2 + lttrlines[i]
						}
						if i == line+3 {
							ln3 = ln3 + lttrlines[i]
						}
						if i == line+4 {
							ln4 = ln4 + lttrlines[i]
						}
						if i == line+5 {
							ln5 = ln5 + lttrlines[i]
						}
						if i == line+6 {
							ln6 = ln6 + lttrlines[i]
						}
						if i == line+7 {
							ln7 = ln7 + lttrlines[i]
						}
						if i == line+8 {
							ln8 = ln8 + lttrlines[i]
						}
					}
				}
				if lttrlines[line] == " " || i == len(SlcOutput)-1 {
					sln1 = append(sln1, ln1)
					sln2 = append(sln2, ln2)
					sln3 = append(sln3, ln3)
					sln4 = append(sln4, ln4)
					sln5 = append(sln5, ln5)
					sln6 = append(sln6, ln6)
					sln7 = append(sln7, ln7)
					sln8 = append(sln8, ln8)
					ln1 = ""
					ln2 = ""
					ln3 = ""
					ln4 = ""
					ln5 = ""
					ln6 = ""
					ln7 = ""
					ln8 = ""

				}

			}
			if i == len(Output)-1 {

				Printer(sln1, w)
				Printer(sln2, w)
				Printer(sln3, w)
				Printer(sln4, w)
				Printer(sln5, w)
				Printer(sln6, w)
				Printer(sln7, w)
				Printer(sln8, w)
			}
		}

	} else {
		// Splits on newlines by default.
		ln1 := ""
		ln2 := ""
		ln3 := ""
		ln4 := ""
		ln5 := ""
		ln6 := ""
		ln7 := ""
		ln8 := ""

		line := 0
		Output := os.Args[1]
		SlcOutput := []rune(Output)
		dash := strings.Index(Output, `\n`)

		for i := 0; i < len(SlcOutput); i++ {
			if i == dash && dash >= 0 {
				if ln1 != "" && ln2 != "" && ln3 != "" && ln4 != "" && ln5 != "" && ln6 != "" && ln7 != "" && ln8 != "" {
					PrintStr(ln1, ln2, ln3, ln4, ln5, ln6, ln7, ln8)
				} else {
					fmt.Println()
				}

				ln1 = ""
				ln2 = ""
				ln3 = ""
				ln4 = ""
				ln5 = ""
				ln6 = ""
				ln7 = ""
				ln8 = ""

				dash = dash + strings.Index(Output[dash+1:], `\n`) + 1
				i++
			} else {
				for j, s := range lttrlines {
					line = j
					if len(s) > 0 && s == string(SlcOutput[i]) {
						break
					}

				}

				for i := line + 1; i <= line+8; i++ {
					if i == line+1 {
						ln1 = ln1 + lttrlines[i]
					}
					if i == line+2 {
						ln2 = ln2 + lttrlines[i]
					}
					if i == line+3 {
						ln3 = ln3 + lttrlines[i]
					}
					if i == line+4 {
						ln4 = ln4 + lttrlines[i]
					}
					if i == line+5 {
						ln5 = ln5 + lttrlines[i]
					}
					if i == line+6 {
						ln6 = ln6 + lttrlines[i]
					}
					if i == line+7 {
						ln7 = ln7 + lttrlines[i]
					}
					if i == line+8 {
						ln8 = ln8 + lttrlines[i]
					}

				}
				if i == len(Output)-1 {
					PrintStr(ln1, ln2, ln3, ln4, ln5, ln6, ln7, ln8)
				}
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			// Handle the error
		}
		file.Close()
	}
}

func CenterPrinter(a string, w int) {
	length := len(a)
	x := w - length
	x /= 2
	fmt.Println(strings.Repeat(" ", x) + a)

}

func RightPrinter(a string, w int) {
	length := len(a)
	x := w - length
	//Works with print but not println
	fmt.Print(strings.Repeat(" ", x) + a)
}

func Printer(a []string, w int) {
	length := 0
	if len(a) == 1 {
		fmt.Println(a[0])
	} else {
		for _, c := range a {
			length = length + len(c)
		}
		width := w - length
		width = width / (len(a) - 1)

		str := ""
		for i, c := range a {
			if i == len(a)-1 {
				str = str + c
			} else {
				str = str + c + strings.Repeat(" ", width-1)
			}

		}

		fmt.Println(str)
	}

}
