package pastebin

import (
	"encoding/xml"
	"errors"
	"fmt"
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

func Client(devKey string, username string, passwrod string) (*client, error) {

	apiDevKey, err := connect(username, passwrod, devKey)

	if err != nil {
		log.Fatal("Unable to connect to PasteBin API", err)
		return nil, errors.New("Unable to connect to PasteBin API")
	}

	return &client{devKey, apiDevKey}, nil
}

func AnonymousClient(devKey string) client {

	return client{devKey, ""}
}

func (client client) CreatePaste(paste model.BasicPaste) (string, error) {

	pasteDto, err := paste.ToDTO()

	if err != nil {
		log.Fatalln("Unable to convert input data: ", err)
		return "", errors.New("Unable to convert input data")
	}

	req := structToValues(*pasteDto)
	req.Add("api_dev_key", client.devKey)
	req.Add("api_option", "paste")

	if client.userKey != "" {
		req.Add("api_user_key", client.userKey)
	}

	res, err := doCall(apiPostUrl, req)

	if err != nil {
		log.Fatal("Error during call to PasteBin API", err)
		return "", err
	}

	generatePasteUrl := extractStringResponse(res)
	return generatePasteUrl, nil
}

func (client client) DeletePaste(pasteKey string) error {

	req := url.Values{
		"api_dev_key":   {client.devKey},
		"api_paste_key": {pasteKey},
		"api_option":    {"delete"},
	}

	if client.userKey != "" {
		req.Add("api_user_key", client.userKey)
	}

	res, err := doCall(apiPostUrl, req)

	if err != nil {
		log.Fatal("Error during call to PasteBin API", err)
		return err
	}

	generatePasteUrl := extractStringResponse(res)

	//TODO check better the errors
	log.Println(generatePasteUrl)

	return nil
}

func (client client) GetPastes() ([]model.Paste, error) {

	req := url.Values{
		"api_dev_key":       {client.devKey},
		"api_user_key":      {client.userKey},
		"api_results_limit": {"100"},
		"api_option":        {"list"},
	}

	res, err := doCall(apiPostUrl, req)

	if err != nil {
		log.Fatal("Error during call to PasteBin API", err)
		return nil, err
	}

	pastesRes := extractXmlResponse(res)

	pastes := make([]model.Paste, len(pastesRes))
	for i, paste := range pastesRes {
		pastes[i] = model.From(paste)
	}

	return pastes, nil
}

func connect(username string, password string, devKey string) (string, error) {

	req := url.Values{
		"api_dev_key":       {devKey},
		"api_user_name":     {username},
		"api_user_password": {password},
	}

	res, err := doCall(apiLoginUrl, req)

	if err != nil {
		log.Fatal("Error during call to PasteBin API", err)
		return "", err
	}

	return extractStringResponse(res), nil
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
		log.Fatalln("Error reading body", err)
	}

	defer res.Body.Close()

	return string(b)
}

func extractXmlResponse(res *http.Response) []dto.Paste {

	b, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	//This beacuse response from pastebin API is an invalid XML
	result := "<pastes>" + string(b) + "</pastes>"

	pastes := dto.PasteWrapper{}

	xml.Unmarshal([]byte(result), &pastes)
	return pastes.Pastes
}

func doCall(url string, data url.Values) (*http.Response, error) {

	res, err := http.PostForm(url, data)

	if err != nil {
		return nil, fmt.Errorf("Error during call to PasteBin API: %w", err)
	}

	return res, nil
}
