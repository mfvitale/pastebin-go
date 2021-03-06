package pastebin

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"

	"github.com/mfvitale/pastebin-go/internal/client/dto"
	"github.com/mfvitale/pastebin-go/model"
)

const apiPostUrl string = "https://pastebin.com/api/api_post.php"
const apiLoginUrl string = "https://pastebin.com/api/api_login.php"

type Client struct {
	devKey  string
	userKey string
}

func NewClient(devKey string, username string, passwrod string) (*Client, error) {

	apiDevKey, err := connect(username, passwrod, devKey)

	if err != nil {
		return nil, fmt.Errorf("Unable to connect to PasteBin API: %w", err)
	}

	return &Client{devKey, apiDevKey}, nil
}

func NewAnonymousClient(devKey string) Client {

	return Client{devKey, ""}
}

func (client Client) CreatePaste(paste model.BasicPaste) (string, error) {

	pasteDto, err := paste.ToDTO()

	if err != nil {
		return "", fmt.Errorf("Unable to convert input data: %w", err)
	}

	req := structToValues(*pasteDto)
	req.Add("api_dev_key", client.devKey)
	req.Add("api_option", "paste")

	if client.userKey != "" {
		req.Add("api_user_key", client.userKey)
	}

	res, err := doCall(apiPostUrl, req)

	if err != nil {
		return "", fmt.Errorf("Error during call to 'CreatePaste' on PasteBin API: %w", err)
	}

	generatePasteUrl := extractStringResponse(res)
	return generatePasteUrl, nil
}

func (client Client) DeletePaste(pasteKey string) (string, error) {

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
		return "", fmt.Errorf("Error during call to 'DeletePaste' on PasteBin API: %w", err)
	}

	return extractStringResponse(res), nil
}

func (client Client) GetPastes() ([]model.Paste, error) {

	req := url.Values{
		"api_dev_key":       {client.devKey},
		"api_user_key":      {client.userKey},
		"api_results_limit": {"100"},
		"api_option":        {"list"},
	}

	res, err := doCall(apiPostUrl, req)

	if err != nil {
		return nil, fmt.Errorf("Error during call to 'GetPastes' on PasteBin API: %w", err)
	}

	pastesRes := dto.PasteWrapper{}

	extractXmlResponse(res, &pastesRes)

	pastes := make([]model.Paste, len(pastesRes.Pastes))
	for i, paste := range pastesRes.Pastes {
		pastes[i] = model.NewPaste(paste)
	}

	return pastes, nil
}

func (client Client) UserInfo() (*model.User, error) {

	req := url.Values{
		"api_dev_key":  {client.devKey},
		"api_user_key": {client.userKey},
		"api_option":   {"userdetails"},
	}

	res, err := doCall(apiPostUrl, req)

	if err != nil {
		return nil, fmt.Errorf("Error during call to 'UserInfo' on PasteBin API: %w", err)
	}

	userRes := dto.User{}

	extractXmlResponse(res, &userRes)

	return model.NewUser(userRes), nil
}

func (client Client) GetRawPaste(pasteKey string) (string, error) {

	req := url.Values{
		"api_dev_key":   {client.devKey},
		"api_user_key":  {client.userKey},
		"api_paste_key": {pasteKey},
		"api_option":    {"show_paste"},
	}

	res, err := doCall(apiPostUrl, req)

	if err != nil {
		return "", fmt.Errorf("Error during call to 'GetRawPaste' on PasteBin API: %w", err)
	}

	return extractStringResponse(res), nil
}

func connect(username string, password string, devKey string) (string, error) {

	req := url.Values{
		"api_dev_key":       {devKey},
		"api_user_name":     {username},
		"api_user_password": {password},
	}

	res, err := doCall(apiLoginUrl, req)

	if err != nil {
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

func extractStringResponse(body []byte) string {

	return string(body)
}

func extractXmlResponse[T any](body []byte, class *T) {

	var result = string(body)

	if reflect.TypeOf(class).Elem() == reflect.TypeOf(dto.PasteWrapper{}) {
		//This beacuse response from pastebin API is an invalid XML
		result = "<pastes>" + result + "</pastes>"
	}

	xml.Unmarshal([]byte(result), class)
}

func doCall(url string, data url.Values) ([]byte, error) {

	res, err := http.PostForm(url, data)

	if err != nil {
		return nil, fmt.Errorf("Error during call to PasteBin API: %w", err)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("Error during parse response body: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(string(body))
	}

	defer res.Body.Close()

	return body, nil
}
