package token

import "time"

type IssueParameter map[string]any

func GetIssueParameterValue[T any](p IssueParameter, key string) T {
	v := p[key]
	if v != nil {
		if value, ok := v.(T); ok {
			return value
		}
	}
	var zero T
	return zero
}

func (p IssueParameter) Username() string {
	return GetIssueParameterValue[string](p, "username")
}

func (p IssueParameter) Password() string {
	return GetIssueParameterValue[string](p, "password")
}

func (p IssueParameter) SetUsername(v string) {
	p["username"] = v
}

func (p IssueParameter) SetPassword(v string) {
	p["password"] = v
}

/*
private token issuer parameter
*/

func (p IssueParameter) AccessToken() string {
	return GetIssueParameterValue[string](p, "access_token")
}

func (p IssueParameter) ExpireTTL() time.Duration {
	return time.Second * time.Duration(GetIssueParameterValue[int64](p, "expired_ttl"))
}

func (p IssueParameter) SetAccessToken(v string) {
	p["access_token"] = v
}
