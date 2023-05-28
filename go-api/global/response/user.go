package response

import (
	"fmt"
	"time"
)

type JsonTime time.Time
func (j JsonTime) MarshalJSON() ([]byte, error) {
	var stmp = fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-21"))
	return []byte(stmp), nil
}

type UserReponse struct {
	Id int32 `json:"id"`
	Nickname string `json:"name"`
	// Birthday time.Time `json:"birthday"`
	Birthday JsonTime `json:"birthday"`
	Gender string `json:"gender"`
	Mobile string `json:"mobile"`
}