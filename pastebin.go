package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/mfvitale/pastebin-go/client/dto"
	"github.com/mfvitale/pastebin-go/model"
)

const API_POST string = "https://pastebin.com/api/api_post.php"
const API_DEV_KEY string = "t5h5nUzpRUFj87qd5wBxpzTC11IJa891"

func main() {

	paste := model.New("Test paste bin", model.Public, "npm run", "10M", "bash")

	generatePasteUrl := createPaste(paste)

	fmt.Println(generatePasteUrl)
}

func createPaste(paste model.Paste) string {

	request := dto.PasteBinRequest{Text: "npm run", Format: "bash"}

	req := structToValues(request)
	req.Add("api_dev_key", API_DEV_KEY)
	req.Add("api_option", "paste")

	res := doCall(req)

	generatePasteUrl := extractStringResponse(res)
	return generatePasteUrl
}

func structToValues(request dto.PasteBinRequest) (values url.Values) {

	toReturn := url.Values{}

	for key, values := range request.ToMap() {
		for _, value := range values {
			toReturn.Add(key, value)
		}
	}

	return toReturn
}

func extractStringResponse(res *http.Response) string {

	b, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	return string(b)
}

func doCall(data url.Values) *http.Response {

	res, err := http.PostForm(API_POST, data)

	if err != nil {
		log.Fatal(err)
	}
	return res
}
