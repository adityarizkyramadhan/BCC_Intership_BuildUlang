package helper

type meta struct {
	Status   bool   `json:"status"`
	Response string `json:"response"`
}

type responseData struct {
	Data interface{} `json:"data"`
	Meta meta        `json:"meta"`
}

func HelperMessage(status bool, response string, data interface{}) responseData {
	messMeta := meta{Status: status, Response: response}
	messData := responseData{Meta: messMeta, Data: data}
	return messData
}
