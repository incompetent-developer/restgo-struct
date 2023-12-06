package line

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo"
	"github.com/incompetent-developer/restgo-struct/header"
)

// PushMessage Sends a push message to a user, group, or room at any time.
func (loa *LineOA) PushMessage(payload string) (messageResp string, err error) {

	/*
		Create url
	*/
	desURL := fmt.Sprintf(PushMessagePath, loa.APIHost, loa.APIVersion)

	/*
		New http client
	*/
	client := &http.Client{}

	/*
		Request
	*/
	req, err := http.NewRequest(http.MethodPost, desURL, strings.NewReader(payload))
	if err != nil {
		return messageResp, fmt.Errorf("Error when make request : %s", err)
	}

	/*
		Set header
	*/
	req.Header.Set(header.Authorization, fmt.Sprintf("Bearer %v", loa.ChannelAccessToken))
	req.Header.Set(header.ContentType, echo.MIMEApplicationJSON)

	/*
		Call
	*/
	resp, err := client.Do(req)
	if err != nil {
		return messageResp, fmt.Errorf("Error when request : %s", err)
	}
	defer resp.Body.Close()

	/*
		Decoder
	*/
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return messageResp, fmt.Errorf("Error when decode resp body : %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		var (
			errResp map[string]interface{}
		)

		if err := json.Unmarshal(body, &errResp); err != nil {
			return messageResp, fmt.Errorf("Error when decode err resp body : %s", err)
		}

		return messageResp, fmt.Errorf("Error from request : %v", errResp)
	}

	return string(body), nil

}

// GetFollowers Get a list of users who added your LINE Official Account as a friend
func (loa *LineOA) GetFollowers(payload LineFollowersRequest) (messageResp LineFollowersResponse, err error) {

	/*
		Create url
	*/
	desURL := fmt.Sprintf(GetAllFollowerPath, loa.APIHost, loa.APIVersion)

	/*
		Make sure session user
	*/
	requestURL, err := url.Parse(desURL)
	if err != nil {
		return messageResp, err
	}

	q := requestURL.Query()
	if payload.Limit != "" {
		q.Add("limit", payload.Limit)
	}

	if payload.Strart != "" {
		q.Add("start", payload.Limit)
	}
	requestURL.RawQuery = q.Encode()

	/*
		New http client
	*/
	client := &http.Client{}

	/*
		Request
	*/
	req, err := http.NewRequest(http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return messageResp, fmt.Errorf("Error when make request : %s", err)
	}

	/*
		Set header
	*/
	req.Header.Set(header.Authorization, fmt.Sprintf("Bearer %v", loa.ChannelAccessToken))
	req.Header.Set(header.ContentType, echo.MIMEApplicationJSON)

	/*
		Call
	*/
	resp, err := client.Do(req)
	if err != nil {
		return messageResp, fmt.Errorf("Error when request : %s", err)
	}
	defer resp.Body.Close()

	/*
		Decoder
	*/
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return messageResp, fmt.Errorf("Error when decode resp body : %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		var (
			errResp map[string]interface{}
		)

		if err := json.Unmarshal(body, &errResp); err != nil {
			return messageResp, fmt.Errorf("Error when decode err resp body : %s", err)
		}

		return messageResp, fmt.Errorf("Error from request : %v", errResp)
	}

	if err := json.Unmarshal(body, &messageResp); err != nil {
		return messageResp, fmt.Errorf("Error when decode resp body : %s", err)
	}

	return messageResp, nil

}

/*
	GetUserProfile You can get the profile information of users who meet one of two conditions:
	- Users who have added your LINE Official Account as a friend
	- Users who haven't added your LINE Official Account as a friend but have sent a message to your LINE Official Account (except users who have blocked your LINE Official Account)
*/
func (loa *LineOA) GetUserProfile(userID string) (messageResp LineUserProfileResponse, err error) {

	/*
		Create url
	*/
	desURL := fmt.Sprintf(GetUserProfilePath, loa.APIHost, loa.APIVersion, userID)

	/*
		New http client
	*/
	client := &http.Client{}

	/*
		Request
	*/
	req, err := http.NewRequest(http.MethodGet, desURL, nil)
	if err != nil {
		return messageResp, fmt.Errorf("Error when make request : %s", err)
	}

	/*
		Set header
	*/
	req.Header.Set(header.Authorization, fmt.Sprintf("Bearer %v", loa.ChannelAccessToken))
	req.Header.Set(header.ContentType, echo.MIMEApplicationJSON)

	/*
		Call
	*/
	resp, err := client.Do(req)
	if err != nil {
		return messageResp, fmt.Errorf("Error when request : %s", err)
	}
	defer resp.Body.Close()

	/*
		Decoder
	*/
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return messageResp, fmt.Errorf("Error when decode resp body : %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		var (
			errResp map[string]interface{}
		)

		if err := json.Unmarshal(body, &errResp); err != nil {
			return messageResp, fmt.Errorf("Error when decode err resp body : %s", err)
		}

		return messageResp, fmt.Errorf("Error from request : %v", errResp)
	}

	if err := json.Unmarshal(body, &messageResp); err != nil {
		return messageResp, fmt.Errorf("Error when decode resp body : %s", err)
	}

	return messageResp, nil

}
