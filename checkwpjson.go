package main

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/TwiN/go-color"
)

type Author struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Slug        string `json:"slug"`
	AvatarURLs  struct {
		Size24 string `json:"24"`
		Size48 string `json:"48"`
		Size96 string `json:"96"`
	} `json:"avatar_urls"`
	Meta []interface{} `json:"meta"`
	ACF  []interface{} `json:"acf"`
	// Add more fields here as needed
}

var authors []Author
var outlast []string

var fo *os.File

func main() {
	input := flag.String("i", "input.txt", "Input List")
	output := flag.String("o", "output.txt", "Output List")
	help := flag.Bool("h", false, "Show help")
	flag.Parse()
	if *help {
		fmt.Println("-h : To Show Help")
		fmt.Println("-i : To Input File Address")
		fmt.Println("-o : To OutPut File Address")
	}

	fo, err := os.OpenFile(*output, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {

		fmt.Println(color.Colorize(color.Red, "[-] Line 51 Error:"+err.Error()))
		fmt.Println(color.Colorize(color.Red, "[-] Error In Permission To Open File"))
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	// Do not verify certificates, do not follow redirects.
	client := &http.Client{
		Transport: tr,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}}

	InputFile := *input
	InputWebs, err := os.Open(InputFile)
	if err != nil {
		fmt.Println(color.Colorize(color.Red, "[-] Line 67 Error:"+err.Error()))
		recover()
	}
	InputBuf := bufio.NewScanner(InputWebs)
	InputBuf.Split(bufio.ScanLines)

	for InputBuf.Scan() {
		InputText := InputBuf.Text()

		path := fmt.Sprintf(InputText+"/%s", "wp-json/wp/v2/users")
		req, _ := http.NewRequest("GET", path, nil)
		resp, erer := client.Do(req)

		if erer != nil {
			fmt.Println(color.Colorize(color.Red, "[-] Line 80 Error:"+erer.Error()))
			continue
		}
		if resp.StatusCode == 200 {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(color.Colorize(color.Red, "[-] Line 86 Error:"+err.Error()))
				continue
			}
			err = json.Unmarshal(body, &authors)
			if err != nil {
				fmt.Println(color.Colorize(color.Red, "[-] Line 91 Error:"+err.Error()))
				continue
			}
			fmt.Println(color.Colorize(color.Green, "[+] Find In : "+path))
			outlast = append(outlast, path+"\n")
			for _, author := range authors {
				outlast = append(outlast, " User Name :\n"+author.Slug+"\n")
				fmt.Println(color.Colorize(color.Green, "[+] UserNames : "+author.Slug))

			}
			outlast = append(outlast, "********************************************\n")

			_, err = fmt.Fprint(fo, outlast)
			if err != nil {
				fmt.Println(color.Colorize(color.Red, "[-] Line 100 Error:"+err.Error()))
				continue
			}
		}
	}

}
