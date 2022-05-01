package dto

import "encoding/xml"

type PasteWrapper struct {
	XMLName xml.Name `xml:"pastes"`
	Pastes  []Paste  `xml:"paste"`
}

type Paste struct {
	Key         string `xml:"paste_key"`
	Date        int64  `xml:"paste_date"`
	Title       string `xml:"paste_title"`
	Size        int    `xml:"paste_size"`
	ExpireDate  int64  `xml:"paste_expire_date"`
	Private     int    `xml:"paste_private"`
	FormatLong  string `xml:"paste_format_long"`
	FormatShort string `xml:"paste_format_short"`
	Url         string `xml:"paste_url"`
	Hits        int    `xml:"paste_hits"`
}
