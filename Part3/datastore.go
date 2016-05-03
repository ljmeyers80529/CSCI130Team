package finalWeb

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"net/http"
)

// retrieve data from datastore if data is not in the memcache.
func retrieveUserInformationDatastore(ctx context.Context, userId string, req *http.Request) (userInformation, error) {

	key := datastore.NewKey(ctx, "Photos", userId, 0, nil)
	var ui userInformation
	err := datastore.Get(ctx, key, &ui)
	if err != nil {
		log.Errorf(ctx, "ERROR retrieveDstore datastore.Get: %s", err)
		return ui, err
	}
	return ui, nil
}

// save any data to the datastore.
func setUserInformationDatastore(ctx context.Context, ui userInformation, req *http.Request) error {

	key := datastore.NewKey(ctx, "Photos", ui.UserId, 0, nil)
	_, err := datastore.Put(ctx, key, &ui)
	if err != nil {
		log.Errorf(ctx, "ERROR storeDstore datastore.Put: %s", err)
		return err
	}
	return nil
}
