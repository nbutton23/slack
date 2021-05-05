package slack

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

//Contains Admin methods for Enterprise Grid
//
// Admin API endpoints reach across an entire Enterprise Grid organization, not individual workspaces.
// For a token to be imbued with Admin scopes, it must be obtained from installing an app on the entire Grid org, not just a workspace within the organization.

// AdminConversatinsSetTeams sets the workspaces in an Enterprise grid org that connect to a public or private channel.
//
// https://api.slack.com/methods/admin.conversations.setTeams
//
// This Admin API method sets the workspaces connected to a channel.
//
// When used with the team_id parameter, this method sets the requested channel_id as a regular channel on the team_id workspace.
//
// When used with the target_team_ids parameter, this method sets the requested channel_id as a cross-workspace shared channel. The channel is shared to all the workspaces in target_team_ids.
// Either way, this method can be used both to add and to remove workspaces from a channel, including the workspace that originated the channel.
//
// channelID: The encoded channelID to add or remove to workspaces.
// orgChannel: True if channel has to be converted to an org channel.
// targetTeamIDs: A list of workspaces to which the channel should be shared. Not required if the channel is being shared org-wide.
// teamID: The workspace to which the channel belongs. Set to empty if the channel is a cross-workspace shared channel.
func (api *Client) AdminConversatinsSetTeams(ctx context.Context, channelID string, orgChannel bool, targetTeamIDs []string, teamID string) error {

	values := url.Values{
		"token":       {api.token},
		"channel_id":  {channelID},
		"org_channel": {strconv.FormatBool(orgChannel)},
	}

	if len(targetTeamIDs) > 0 {
		values.Add("target_team_ids", strings.Join(targetTeamIDs, ","))
	}

	if teamID != "" {
		values.Add("team_id", teamID)
	}

	resp := &SlackResponse{}

	endpoint := APIURL + "admin.conversations.setTeams"
	if err := postForm(ctx, api.httpclient, endpoint, values, resp, api); err != nil {
		return fmt.Errorf("Failed to set teams on conversation '%s': %s", channelID, err)
	}

	return nil
}
