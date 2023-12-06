package line

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/incompetent-developer/restgo-struct/logger"
)

func TestPushMessage(t *testing.T) {
	w := LineOA{
		APIHost:            "https://api.line.me",
		APIVersion:         "v2",
		ChannelAccessToken: "",
	}

	data := map[string]interface{}{
		"to": "",
		"messages": []interface{}{
			map[string]interface{}{
				"type": "text",
				"text": "HI",
			},
		},
	}

	payload, _ := json.Marshal(data)

	resp, err := w.PushMessage(string(payload))
	if err != nil {
		t.Errorf("Err : %v", err)
	}

	fmt.Println(resp)

}

func TestGetFollowers(t *testing.T) {
	w := LineOA{
		APIHost:            "https://api.line.me",
		APIVersion:         "v2",
		ChannelAccessToken: "",
	}

	data := LineFollowersRequest{
		Limit: "10",
	}

	resp, err := w.GetFollowers(data)
	if err != nil {
		t.Errorf("Err : %v", err)
	}
	logger.Info(resp)

}

func TestGetUserProfile(t *testing.T) {
	w := LineOA{
		APIHost:            "https://api.line.me",
		APIVersion:         "v2",
		ChannelAccessToken: "",
	}

	userID := ""

	resp, err := w.GetUserProfile(userID)
	if err != nil {
		t.Errorf("Err : %v", err)
	}
	logger.Info(resp)

}
