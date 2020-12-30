package auth

// GetHeaders returns the
func GetHeaders() map[string]string {
	token := GetJSONToken()
	fullToken := "Bearer " + token.Value
	headers := map[string]string{
		"Authorization": fullToken,
		"Content-Type":  "application/json",
	}
	return headers
}

func getBaseURL() string {
	return "https://slack.com/api"
}

func getFullURL(endpoint string) string {
	return getBaseURL() + endpoint
}

func GetSetTopicEndpoint() string {
	return getFullURL("/conversations.setTopic")
}

func GetListConversationsEndpoint() string {
	return getFullURL("/conversations.list")
}

func GetJoinConversationsEndpoint() string {
	return getFullURL("/conversations.join")
}
