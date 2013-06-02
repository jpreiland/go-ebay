package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/xml"
)

type Item struct {
	ItemID string `xml:"itemId"`
	Title string `xml:"title"`
}

type Response struct {
	XMLName xml.Name `xml:"findItemsByKeywordsResponse"`
	Items []Item `xml:"searchResult>item"`
}


// build the url for HTTP GET 
func buildRequest(appid, keywords, n string) string {
	keywords = keywordConvert(keywords)

	url := "http://svcs.ebay.com/services/search/FindingService/v1?OPERATION-NAME=findItemsByKeywords&SERVICE-VERSION=1.0.0&SECURITY-APPNAME="
	url += appid
	url += "&GLOBAL-ID=EBAY-US&RESPONSE-DATA-FORMAT=XML&REST-PAYLOAD&keywords="
	url += keywords
	url += "&paginationInput.entriesPerPage="
	url += n

	return url
}

// URL-encode any spaces in keywords
// all " " (spaces) should be converted to "%20"
func keywordConvert(keywords string) string {
	for i := 0; i < len(keywords); i++ {
    	if keywords[i] == ' ' {
    		keywords = keywords[:(i)] + "%20" + keywords[(i+1):]
    	}
    }
    return keywords
}


func main() {
	// build request url
	appID := "" //YOUR APPID HERE
	query := "" //YOUR SEARCH QUERY HERE
	num_items := "10" //NUMBER OF RESULTS TO RETURN
	url := buildRequest(appID, query, num_items)

	// send request and store response
	r, _ := http.Get(url)
	response, _ := ioutil.ReadAll(r.Body)
	v := Response{}
	err := xml.Unmarshal([]byte(response), &v)
	if err != nil {
   		fmt.Printf("error: %v", err)
   		return
    }

    // output information
    for i := 0; i < len(v.Items); i++ {
    	fmt.Println(v.Items[i].Title)
    }
}