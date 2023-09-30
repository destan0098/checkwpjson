package Part1

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/destan0098/checkwpjson/cmd/Part2"
	"io/ioutil"
	"net/http"
	"os"
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

var outlast []string
var authors []Author
var err error

func Part1(path string, fo *os.File, InputText string) {
	fmt.Println(path)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	// Do not verify certificates, do not follow redirects.
	client := &http.Client{
		Transport: tr,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}}

	req, _ := http.NewRequest("GET", path, nil)
	//Send Request To Address
	resp, erer := client.Do(req)

	if erer != nil {
		fmt.Println(color.Colorize(color.Red, "[-]  Error:"+erer.Error()))
		recover()
	}
	if resp.StatusCode == 200 {
		//Check Response status Code
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(color.Colorize(color.Red, "[-]  Error:"+err.Error()))
			recover()
		}
		err = json.Unmarshal(body, &authors)
		//Parse Json Values To Show
		if err != nil {
			fmt.Println(color.Colorize(color.Red, "[-] Line 91 Error:"+err.Error()))
			recover()
		}
		fmt.Println(color.Colorize(color.Green, "[+] Find In : "+path))
		outlast = append(outlast, path+"\n")
		for _, author := range authors {
			outlast = append(outlast, " User Name :\n"+author.Slug+"\n")
			fmt.Println(color.Colorize(color.Green, "[+] UserNames : "+author.Slug))

		}
		outlast = append(outlast, "********************************************\n")

		_, err = fmt.Fprint(fo, outlast)
		//Save In output File
		if err != nil {
			fmt.Println(color.Colorize(color.Red, "[-] Line 100 Error:"+err.Error()))
			recover()
		}
	} else {
		Part2.Part2(fo, InputText)
	}
}
