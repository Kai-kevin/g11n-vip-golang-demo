package bean

type QueryTranslationByCompRespData struct {
	ProductName string            `json:"productName"`
	Version     string            `json:"version"`
	Pseudo      bool              `json:"preudo"`
	Component   string            `json:"componet"`
	Messages    map[string]string `json:"messages"`
	Locale      string            `json:"locale"`
	Status      string            `json:"status"`
	Id          int               `json:"id"`
}

type QueryTranslationByCompRespEvent struct {
	Response  Response                       `json:"response"`
	Signature string                         `json:"signature"`
	Data      QueryTranslationByCompRespData `json:"data"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
