package bean

type QueryTranslationByKeyRespData struct {
	ProductName string `json:"productName"`
	Version     string `json:"version"`
	Pseudo      string `json:"pseudo"`
	Source      bool   `json:"source"`
	Translation string `json:"translation"`
	Locale      string `json:"locale"`
	Key         string `json:"key"`
	Component   string `json:"component"`
	Status      string `json:"status"`
}

type QueryTranslationByKeyRespEvent struct {
	Response  Response                      `json:"response"`
	Signature string                        `json:"signature"`
	Data      QueryTranslationByKeyRespData `json:"data"`
}
