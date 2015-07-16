package plurk

import (
//"time"
)

// Plurk User Struct

type User struct {
	Id              int
	VerifiedAccount bool   `json:"verified_account"`
	DisplayName     string `json:"display_name"`
	Nickname        string `json:"nick_name"`
	Karma           float32
	DefaultLanguage string `json:"default_lang"`
	DateFormat      int
	BirthdayPrivacy int    `json:"bday_privacy"`
	FullName        string `json:"full_name"`
	Gender          int
	NameColor       string `json:"name_color"`
	Timezone        string
	Avatar          int
	Birthday        Time `json:"date_of_birth"`
}
