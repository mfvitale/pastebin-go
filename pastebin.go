package pastebin

import (
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

func (client client) CreatePaste(paste model.Paste) string {

	request := dto.PasteBinRequest{Text: "npm run", Format: "bash"}

	req := structToValues(request)
	req.Add("api_dev_key", client.devKey)
	req.Add("api_option", "paste")

	if client.userKey != "" {
		req.Add("api_user_key", client.userKey)
	}

	res := doCall(apiPostUrl, req)

	generatePasteUrl := extractStringResponse(res)
	return generatePasteUrl
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

	b, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	return string(b)
}

func doCall(url string, data url.Values) *http.Response {

	res, err := http.PostForm(url, data)

	if err != nil {
		log.Fatal(err)
	}
	return res
}
