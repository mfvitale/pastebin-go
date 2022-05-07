package pastebin

import (
	"log"
	"os"
	"regexp"
	"testing"

	"github.com/mfvitale/pastebin-go/model"
)

func TestAnonymPasteCreation(t *testing.T) {

	anonClient := NewAnonymousClient(os.Getenv("DEV_KEY"))

	paste := model.FullPaste("Test paste bin", model.Public, "npm run", model.TEN_MINUTES, "bash")
	res, err := anonClient.CreatePaste(paste)

	want := regexp.MustCompile(`https://pastebin.com/([a-zA-Z0-9_]*)`)

	if !want.MatchString(res) || err != nil {
		t.Fatalf(`anonClient.CreatePaste(paste) = %q, %v, want match for %#q, nil`, res, err, want)
	}
}

func TestLoggedPasteCreation(t *testing.T) {

	log.Println(os.Getenv("USERNAME"), os.Getenv("PASSWORD"))
	client, err := NewClient(os.Getenv("DEV_KEY"), os.Getenv("USERNAME"), os.Getenv("PASSWORD"))

	if err != nil {
		log.Fatalln("Error creating client ", err)
	}

	paste := model.FullPaste("Test paste bin", model.Public, "npm run", model.TEN_MINUTES, "bash")
	res, err := client.CreatePaste(paste)

	want := regexp.MustCompile(`https://pastebin.com/([a-zA-Z0-9_]*)`)

	if !want.MatchString(res) || err != nil {
		t.Fatalf(`client.CreatePaste(paste) = %q, %v, want match for %#q, nil`, res, err, want)
	}
}

func TestGetPastes(t *testing.T) {

	client, err := NewClient(os.Getenv("DEV_KEY"), os.Getenv("USERNAME"), os.Getenv("PASSWORD"))

	res, err := client.GetPastes()

	if len(res) == 0 || err != nil {
		t.Fatalf(`client.GetPastes size = %q, %v, want != 0, nil`, len(res), err)
	}
}

func TestGetRawPaste(t *testing.T) {

	client, err := NewClient(os.Getenv("DEV_KEY"), os.Getenv("USERNAME"), os.Getenv("PASSWORD"))

	if err != nil {
		log.Fatalln("Error creating client ", err)
	}

	paste := model.FullPaste("This is a paste", model.Public, "npm run", model.TEN_MINUTES, "bash")
	res, err := client.CreatePaste(paste)

	if err != nil {
		log.Fatalln("Error creating paste ", err)
	}

	want := regexp.MustCompile(`https://pastebin.com/(?P<Key>[a-zA-Z0-9_]*)`)
	matches := want.FindStringSubmatch(res)
	index := want.SubexpIndex("Key")

	raw, err := client.GetRawPaste(matches[index])

	if raw != "npm run" || err != nil {
		t.Fatalf(`client.GetRawPaste(%v) = %q, %v, want %#q, nil`, matches[index], raw, err, "npm run")
	}
}
