package finalWeb

import (
	"encoding/json"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/memcache"
	"net/http"
)

// retrieve user information from memcache.
func retrieveUserInformationMemcache(userId string, req *http.Request) (UserInformation, bool) {
	var ui UserInformation

	ctx := appengine.NewContext(req)
	item, err := memcache.Get(ctx, userId)
	if err != nil {
		return initializeUserData(userId), false
	}
	err = json.Unmarshal(item.Value, &ui)
	return ui, true
}

// save user information in memchache.
func setUserInformationMemcache(userInfo UserInformation, req *http.Request) error {
	ctx := appengine.NewContext(req)

	log.Infof(ctx, "Key value %v", userInfo)
	bs, err := json.Marshal(userInfo)
	if err != nil {
		return err
	}
	memData := memcache.Item{
		Key:   userInfo.UserId,
		Value: bs,
	}
	err = memcache.Set(ctx, &memData)
	if err != nil {
		return err
	}
	return nil
}
