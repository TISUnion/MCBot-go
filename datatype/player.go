package datatype

import "github.com/tidwall/gjson"

type Player struct {
	Username     string
	Account      string
	Password     string
	UUID         string
	AccessToken  string
	ClientToken  string
	Id_          string
	SharedSecret string
}

func (pl *Player) GetToken() bool {
	if pl.authenticate() {
		return true
	} else {
		//尝试连接5次
		for i := 5; i < 0; i-- {
			if pl.authenticate() {
				return true
			}
		}
	}
	return false
}

func (pl *Player) authenticate() bool {
	param := `{
		"agent": {
			"name": "Minecraft",
			"version": 1
		},
		"username": "` + pl.Account + `",
		"password": "` + pl.Password + `"
	}`
	if res, ok := TokenRequest("authenticate", param); ok {
		jsonresp := gjson.Parse(res)
		pl.AccessToken = jsonresp.Get("accessToken").String()
		if pl.AccessToken == "" {
			return false
		}
		pl.ClientToken = jsonresp.Get("clientToken").String()
		selectedProfile := jsonresp.Get("selectedProfile")
		pl.Id_ = selectedProfile.Get("id").String()
		pl.Username = selectedProfile.Get("name").String()
		return true
	}
	return false
}

func (pl *Player) Signout() bool {
	if pl.Account == "" {
		return false
	}
	param := `{
			"username": "` + pl.Account + `",
			"password": "` + pl.Password + `"
		}`
	if res, ok := TokenRequest("signout", param); ok && (res == "") {
		return true
	}
	return false
}

func (pl *Player) validate() bool {
	param := `{
		"accessToken": "` + pl.AccessToken + `",
		"clientToken": "` + pl.ClientToken + `"
	}`
	if res, ok := TokenRequest("validate", param); ok && (res == "") {
		return true
	}
	return false
}

func (pl *Player) Invalidate() bool {
	if pl.AccessToken == "" {
		return false
	}
	param := `{
		"accessToken": "` + pl.AccessToken + `",
		"clientToken": "` + pl.ClientToken + `"
	}`
	if res, ok := TokenRequest("invalidate", param); ok && (res == "") {
		return true
	}
	return false
}

func (pl *Player) SetSession(serverId string) bool {
	if pl.AccessToken == "" {
		pl.GetToken()
	}
	param := `{
		"accessToken": "` + pl.AccessToken + `",
		"selectedProfile": "` + pl.Id_ + `",
		"serverId": "` + serverId + `"
	}`
	if res, ok := SessionRequest(param); ok && (res == "") {
		return true
	}
	return false
}
