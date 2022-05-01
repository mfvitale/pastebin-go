package model

import "github.com/mfvitale/pastebin-go/client/dto"

type Paste struct {
	Key         string
	Date        int64
	Title       string
	Size        int
	ExpireDate  int64
	Private     int
	FormatLong  string
	FormatShort string
	Url         string
	Hits        int
}

func From(paste dto.Paste) Paste {
	return Paste{paste.Key,
		paste.Date,
		paste.Title,
		paste.Size,
		paste.ExpireDate,
		paste.Private,
		paste.FormatLong,
		paste.FormatShort,
		paste.Url,
		paste.Hits}
}
