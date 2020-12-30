package company

import (
	"slack-topic/slack"
	"slack-topic/utils"
)

func FilterRelevantConversations(
	tcs []slack.TeamConversation,
) []slack.TeamConversation {
	var rc []slack.TeamConversation
	lcs := GetLevelTeamsConversations()
	for _, c := range tcs {
		if utils.IsRelevantConversation(lcs, c.Name) {
			rc = append(rc, c)
		}
	}

	return rc
}
