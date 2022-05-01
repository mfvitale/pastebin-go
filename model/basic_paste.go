package model

type Visibility int

const (
	Public   Visibility = 0
	Unlisted            = 1
	Private             = 2
)

func (e Visibility) String() string {
	switch e {
	case Public:
		return "0"
	case Unlisted:
		return "1"
	case Private:
		return "2"
	default:
		return ""
	}
}

type BasicPaste struct {
	text       string
	visibility Visibility
	name       string
	expireDate string
	format     string
}

func FullPaste(name string, visibility Visibility, text string, expireDate string, format string) BasicPaste {
	return BasicPaste{text, visibility, name, expireDate, format}
}

func SimplePaste(text string) BasicPaste {
	return BasicPaste{text, Public, "", "", ""}
}
