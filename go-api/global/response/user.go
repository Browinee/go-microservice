package response

import "time"


type UserReponse struct {
	Id int32 `json:"id"`
	Nickname string `json:"name"`
	Birthday time.Time `json:"birthday"`
	Gender string `json:"gender"`
	Mobiel string `json:"mobile"`
}