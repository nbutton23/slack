package slack

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestAdminConversatinsSetTeams(t *testing.T) {
	http.HandleFunc("/admin.conversations.setTeams", testUninstallAppHandler)
	once.Do(startServer)

	api := New("test-token", OptionAPIURL("http://"+serverAddr+"/"))

	err := api.AdminConversatinsSetTeams(context.TODO(), "channel_id", false, []string{"team1", "team2"}, "")

	if err != nil {
		t.Errorf("Failed, but should have succeeded")
	}
}

func testEnterpriseAdminHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(SlackResponse{Ok: true})
	w.Write(response)
}
