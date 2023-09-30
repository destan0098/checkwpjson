package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/destan0098/checkwpjson/cmd/Part1"
	"os"
	"strings"
)

var path string
var fo *os.File

var err error

func main() {
	//Receive input from the user.
	input := flag.String("i", "input.txt", "Input List")
	output := flag.String("o", "output.txt", "Output List")
	help := flag.Bool("h", false, "Show help")
	flag.Parse()
	if *help {
		fmt.Println("-h : To Show Help")
		fmt.Println("-i : To Input File Address")
		fmt.Println("-o : To OutPut File Address")
	}
	//Open OutPut File in Directory
	fo, err = os.OpenFile(*output, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {

		fmt.Println(color.Colorize(color.Red, "[-]  Error:"+err.Error()))
		fmt.Println(color.Colorize(color.Red, "[-] Error In Permission To Open File"))
	}
	//To Skip Insecure SSL

	InputFile := *input
	//Open Input File
	InputWebs, err := os.Open(InputFile)
	if err != nil {
		fmt.Println(color.Colorize(color.Red, "[-]  Error:"+err.Error()))
		recover()
	}
	//Scan All Line of input file
	InputBuf := bufio.NewScanner(InputWebs)
	InputBuf.Split(bufio.ScanLines)

	for InputBuf.Scan() {
		//Check End Of File Or not
		InputText := InputBuf.Text()

		if !strings.HasSuffix(InputText, "/") {

			//It puts the values of the file line by line into the variable.
			path = fmt.Sprintf(InputText+"/%s", "wp-json/wp/v2/users")

			// Add WP-json Directory To your Address
		} else {

			path = fmt.Sprintf(InputText+"%s", "wp-json/wp/v2/users")

		}
		Part1.Part1(path, fo, InputText)
	}

}
