package finalWeb

import (
	"encoding/json"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/memcache"
	"net/http"
)

// retrieve user information from memcache.
func retrieveUserInformationMemcache(ctx context.Context, userId string, req *http.Request) (userInformation, error) {

	var ui userInformation

	item, err := memcache.Get(ctx, userId)
	if err != nil {
		// 		item, err = retrieveUserInformationDatastore(id, req)         // todo add data store access.
		if err != nil {
			return ui, err
		}
		setUserInformationMemcache(ctx, ui, req)
		return ui, nil
	}
	err = json.Unmarshal(item.Value, &ui)
	return ui, err
}

// save user information in memchache.
func setUserInformationMemcache(ctx context.Context, userInfo userInformation, req *http.Request) error {

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
