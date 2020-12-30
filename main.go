package main

import (
	"slack-topic/company"
	"slack-topic/slack"
)

func main() {
	// Get a list of all conversations and compare them to the relevant ones for teams
	allConversations := slack.GetConversations()

	relevantConversations := company.FilterRelevantConversations(
		allConversations,
	)

	// Request to join relevant conversations
	slack.RequestToJoinConversations(&relevantConversations)

	slack.SetConversationTopics(&relevantConversations)
}
