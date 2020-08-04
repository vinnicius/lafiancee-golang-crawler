package main

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/gocolly/colly/v2"
)

// TODO: This clearly doesnt look well made.
type siteData struct {
	TpaWidgetNativeInitData struct {
		TPAMultiSectionjenllqhb struct {
			WixCodeProps struct {
				Product struct {
					ID    string `json:"id"`
					Name  string `json:"name"`
					Price int    `json:"price"`
					Media []struct {
						FullURL string `json:"fullUrl"`
					} `json:"media"`
					Options []struct {
						Option []interface{} `json:"selections"`
					} `json:"options"`
				} `json:"product"`
			} `json:"wixCodeProps"`
		} `json:"TPAMultiSection_jenllqhb"`
	} `json:"tpaWidgetNativeInitData"`
}

func main() {

	baseURL := "https://www.lafiancee.com.br/vestidos-de-noiva"
	baseCollector := colly.NewCollector()
	dressCollector := colly.NewCollector()

	// Get "Visualização Rápida" link
	baseCollector.OnHTML("._34sIs", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Link found: %s \n", link)
		// Then pass the link to the dress collector
		dressCollector.Visit(link)
	})

	dressCollector.OnResponse(func(r *colly.Response) {
		// Initialize regexp with js search parameters
		regexp := regexp.MustCompile(`var warmupData = ({.*});`)

		var dress siteData

		// Try to unmarshal json subdata
		if err := json.Unmarshal(regexp.FindSubmatch(r.Body)[1], &dress); err != nil {
			fmt.Printf("Error unmarshaling data json: %s \n", err)
		}

		fmt.Println("Data: \n", dress)
	})

	baseCollector.Visit(baseURL)
}
