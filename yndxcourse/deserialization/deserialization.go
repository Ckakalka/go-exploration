package main

import (
	"encoding/json"
	"fmt"
)

const rawResp = `
{
    "header": {
        "code": 0,
        "message": ""
    },
    "data": [{
        "type": "user",
        "id": 100,
        "attributes": {
            "email": "bob@yandex.ru",
            "article_ids": [10, 11, 12]
        }
    }]
}
`

type Response struct {
	Header HeaderResponse `json:"header"`
	Data   DataResponse   `json:"data"`
}

type HeaderResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type DataResponse []DataItemResponse

type DataItemResponse struct {
	Type       string                `json:"type"`
	Id         int                   `json:"id"`
	Attributes AttributeDataResponse `json:"attributes"`
}

type AttributeDataResponse struct {
	Email      string `json:"email"`
	ArticleIds []int  `json:"article_ids"`
}

func readResponse(rawResp string) (Response, error) {
	var result Response
	err := json.Unmarshal([]byte(rawResp), &result)
	return result, err
}

func main() {
	response, err := readResponse(rawResp)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println(response)
	}
}
