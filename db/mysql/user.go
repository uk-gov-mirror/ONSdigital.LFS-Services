package mysql

import (
	"fmt"
	"services/types"
)

func (s MySQL) GetUserID(user string) (types.UserCredentials, error) {
	var creds types.UserCredentials

	col := s.DB.Collection("users")
	res := col.Find("username", user)

	if res == nil {
		return creds, fmt.Errorf("user %s not found", user)
	}

	defer func() { _ = res.Close() }()

	ok := res.Next(&creds)
	if !ok {
		return creds, fmt.Errorf("user %s not found", user)
	}
	return creds, nil
}