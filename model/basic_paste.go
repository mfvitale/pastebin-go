package model

import (
	dtomapper "github.com/dranikpg/dto-mapper"
	"github.com/mfvitale/pastebin-go/client/dto"
)

var mapper = dtomapper.Mapper{}

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

type Expiration int

const (
	NEVER Expiration = iota
	TEN_MINUTES
	ONE_HOUR
	ONE_DAY
	ONE_WEEK
	TWO_WEEK
	ONE_MONTH
	SIX_MONTH
	ONE_YEAR
)

func (exp Expiration) String() string {
	switch exp {
	case NEVER:
		return "N"
	case TEN_MINUTES:
		return "10M"
	case ONE_HOUR:
		return "1H"
	case ONE_DAY:
		return "1D"
	case ONE_WEEK:
		return "1W"
	case TWO_WEEK:
		return "2W"
	case ONE_MONTH:
		return "1M"
	case SIX_MONTH:
		return "6M"
	case ONE_YEAR:
		return "1Y"
	}
	return "unknown"
}

type BasicPaste struct {
	Text       string
	Visibility Visibility
	Name       string
	ExpireDate Expiration
	Format     string
}

func FullPaste(name string, visibility Visibility, text string, expireDate Expiration, format string) BasicPaste {
	return BasicPaste{text, visibility, name, expireDate, format}
}

func SimplePaste(text string) BasicPaste {
	return BasicPaste{text, Public, "", NEVER, ""}
}

func (paste BasicPaste) ToDTO() (*dto.PasteBinRequest, error) {

	mapper.AddConvFunc(func(visibility Visibility) string {
		return visibility.String()
	})

	mapper.AddConvFunc(func(expiration Expiration) string {
		return expiration.String()
	})

	pasteDto := dto.PasteBinRequest{}
	err := mapper.Map(&pasteDto, paste)

	if err != nil {
		return nil, err
	}

	return &pasteDto, nil
}
