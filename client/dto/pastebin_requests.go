package dto

type PasteBinRequest struct {
	Text       string
	Visibility string
	Name       string
	ExpireDate string
	Format     string
}

func (r *PasteBinRequest) ToMap() map[string][]string {

	return map[string][]string{
		"api_paste_code":        {r.Text},
		"api_paste_private":     {r.Visibility},
		"api_paste_name":        {r.Name},
		"api_paste_expire_date": {r.ExpireDate},
		"api_paste_format":      {r.Format},
	}
}
