package services

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func GetUserInfoRequest(_ context.Context, r *http.Request, t interface{}) error {
	userRequest := t.(UserRequest)
	r.URL.Path += "/user/" + strconv.Itoa(userRequest.Uid)
	return nil
}

func GetUserInfoResponse(_ context.Context, w *http.Response) (response interface{}, err error) {
	if w.StatusCode > 400 {
		return nil, errors.New("no data")
	}
	var userResponse UserResponse
	err = json.NewDecoder(w.Body).Decode(&userResponse)
	if err != nil {
		return nil, err
	}
	return userResponse, nil
}