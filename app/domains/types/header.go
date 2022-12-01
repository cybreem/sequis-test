package types

import "time"

type HeaderContent struct {
	AcceptLanguage string `header:"Accept-Language"`
	ContentType    string `header:"Content-Type"`
	Authorization  string `header:"Authorization"`
	AccessID       string `header:"Access-ID"`
}

type GeneralResponseType struct {
	ResponseStatus    bool      `json:"response_status"`
	ResponseCode      int       `json:"response_code"`
	ResponseMessage   string    `json:"response_message"`
	ResponseTimestamp time.Time `json:"response_timestamp"`
}
