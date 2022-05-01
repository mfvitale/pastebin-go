package pastebin

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/mfvitale/pastebin-go/client/dto"
	"github.com/mfvitale/pastebin-go/model"
)

const apiPostUrl string = "https://pastebin.com/api/api_post.php"
const apiLoginUrl string = "https://pastebin.com/api/api_login.php"

type client struct {
	devKey  string
	userKey string
}

func Client(devKey string, username string, passwrod string) client {

	apiDevKey := connect(username, passwrod, devKey)

	return client{devKey, apiDevKey}
}

func AnonymousClient(devKey string) client {

	return client{devKey, ""}
}

func (client client) CreatePaste(paste model.BasicPaste) string {

	request := dto.PasteBinRequest{Text: "npm run", Format: "bash"}

	req := structToValues(request) //TODO try with github.com/dranikpg/dto-mapper
	req.Add("api_dev_key", client.devKey)
	req.Add("api_option", "paste")

	if client.userKey != "" {
		req.Add("api_user_key", client.userKey)
	}

	res := doCall(apiPostUrl, req)

	generatePasteUrl := extractStringResponse(res)
	return generatePasteUrl
}

func (client client) GetPastes() []model.Paste {

	req := url.Values{
		"api_dev_key":       {client.devKey},
		"api_user_key":      {client.userKey},
		"api_results_limit": {"100"},
		"api_option":        {"list"},
	}

	res := doCall(apiPostUrl, req)

	pastesRes := extractXmlResponse(res)

	pastes := make([]model.Paste, len(pastesRes))
	for i, paste := range pastesRes {
		pastes[i] = model.From(paste)
	}

	return pastes
}

func connect(username string, password string, devKey string) string {

	req := url.Values{
		"api_dev_key":       {devKey},
		"api_user_name":     {username},
		"api_user_password": {password},
	}

	res := doCall(apiLoginUrl, req)
	return extractStringResponse(res)
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

	//TODO check status
	b, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	return string(b)
}

func extractXmlResponse(res *http.Response) []dto.Paste {

	b, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	//This beacuse is an invalid XML
	result := "<pastes>" + string(b) + "</pastes>"

	pastes := dto.PasteWrapper{}

	xml.Unmarshal([]byte(result), &pastes)
	return pastes.Pastes
}

func doCall(url string, data url.Values) *http.Response {

	res, err := http.PostForm(url, data)

	if err != nil {
		log.Fatal(err)
	}
	return res
}
