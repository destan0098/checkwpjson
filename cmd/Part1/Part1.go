package Part1

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/TwiN/go-color"
	"io/ioutil"
	"net/http"
	"strings"
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
var path string

func Part1(domains string) ([]string, string) {
	fmt.Println(domains)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	// Do not verify certificates, do not follow redirects.
	client := &http.Client{
		Transport: tr,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}}
	if !strings.HasSuffix(domains, "/") {

		//It puts the values of the file line by line into the variable.
		path = fmt.Sprintf(domains+"/%s", "wp-json/wp/v2/users")

		// Add WP-json Directory To your Address
	} else {

		path = fmt.Sprintf(domains+"%s", "wp-json/wp/v2/users")

	}
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
			fmt.Println(color.Colorize(color.Red, "[-] Line 77 Error:"+err.Error()))
			recover()
		}
		fmt.Println(color.Colorize(color.Green, "[+] Find In : "+path))

		for _, author := range authors {
			outlast = append(outlast, author.Slug)
			fmt.Println(color.Colorize(color.Green, "[+] UserNames : "+author.Slug))

		}

		return outlast, domains
		//Save In output File

	} else {
		return nil, ""

	}

}
