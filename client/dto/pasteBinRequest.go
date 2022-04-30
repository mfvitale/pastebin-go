package dto

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

type PasteBinRequest struct {
	Text       string
	Visibility Visibility
	Name       string
	ExpireDate string
	Format     string
}

func (r *PasteBinRequest) ToMap() map[string][]string {

	return map[string][]string{
		"api_paste_code":        {r.Text},
		"api_paste_private":     {r.Visibility.String()},
		"api_paste_name":        {r.Name},
		"api_paste_expire_date": {r.ExpireDate},
		"api_paste_format":      {r.Format},
	}
}
