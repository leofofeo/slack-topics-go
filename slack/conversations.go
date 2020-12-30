package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"slack-topic/auth"
	"time"
)

type TeamConversation struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type ConversationsAPIResponse struct {
	Channels []TeamConversation `json:"channels"`
}

func getRequestBody(
	url string,
	method string,
	jsonPayload []byte,
) ([]byte, int) {
	client := &http.Client{}

	req, _ := http.NewRequest(method, url, bytes.NewBuffer(jsonPayload))

	headers := auth.GetHeaders()
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	res, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	return body, res.StatusCode
}

func GetConversations() []TeamConversation {
	url := auth.GetListConversationsEndpoint()

	body, _ := getRequestBody(url, "GET", nil)

	var jsonRes ConversationsAPIResponse
	json.Unmarshal(body, &jsonRes)

	return jsonRes.Channels
}

func RequestToJoinConversations(conversations *[]TeamConversation) {
	cs := *conversations

	url := auth.GetJoinConversationsEndpoint()

	for _, c := range cs {
		channelStr := `"channel": "` + c.Id + `"`
		var jsonStr = []byte(`{` + channelStr + `}`)
		_, statusCode := getRequestBody(url, "POST", jsonStr)

		if statusCode != 200 {
			fmt.Println(
				"Bad request, status code: " + fmt.Sprint(statusCode),
			)
		}
	}
}

func SetConversationTopics(conversations *[]TeamConversation) {
	cs := *conversations
	url := auth.GetSetTopicEndpoint()

	for _, c := range cs {
		t := time.Now()

		topic := "Testing set topic from Go for channel id " + fmt.Sprint(c.Id) + " at " + fmt.Sprint(t)

		channelStr := `"channel": "` + c.Id + `"`
		topicStr := `"topic": "` + topic + `"`

		var jsonStr = []byte(
			`{` + channelStr + `, ` + topicStr + `}`,
		)

		_, statusCode := getRequestBody(url, "POST", jsonStr)

		if statusCode != 200 {
			fmt.Println(
				"Bad request, status code: " + fmt.Sprint(statusCode),
			)
		}
	}

}
