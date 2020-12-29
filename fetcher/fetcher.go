package fetcher

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
)

const u string = "https://www.microsoft.com/en-us/download/confirmation.aspx?id=56519"
const exp string = "https://download.*?\\.json"

//Fetch downloads the service tags manifest
func Fetch() (io.Reader, error) {
	var client http.Client
	var blank io.Reader

	// This first page contains a burried link that we need to find to download
	// the actual public IP ranges file
	page, err := client.Get(u)
	if err != nil {
		return blank, err
	}
	defer page.Body.Close()

	body, err := ioutil.ReadAll(page.Body)
	if err != nil {
		return blank, err
	}

	// Extract the URI of the current ranges file
	re := regexp.MustCompile(exp)
	u := re.FindString(string(body))
	uri, err := url.Parse(u)
	if err != nil {
		return blank, err
	}

	// Fetch and return the file proper
	tags, err := client.Get(uri.String())
	if err != nil {
		return blank, err
	}

	return tags.Body, nil
}
