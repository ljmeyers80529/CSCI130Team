package finalWeb

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"io/ioutil"
	"net/http"
)

type Usernames struct {
	Name string
}

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

func checkUserExists(res http.ResponseWriter, req *http.Request) bool {
	ctx := appengine.NewContext(req)

	// retrieve the incoming word as it is typed.
	var w Usernames
	bs, err := ioutil.ReadAll(req.Body)
	//
	log.Infof(ctx, "Received information: %v", string(bs))
	//
	if err != nil {
		log.Infof(ctx, err.Error())
	}
	w.Name = string(bs)
	log.Infof(ctx, "ENTERED wordCheck - w.Name: %v", w.Name)

	// check the incoming word against what is currently in the datastore
	key := datastore.NewKey(ctx, "Dictionary", w.Name, 0, nil)
	err = datastore.Get(ctx, key, &w)
	return err != nil

}

func commitNewUsername(res http.ResponseWriter, req *http.Request, user string) {
	ctx := appengine.NewContext(req)

        var w Usernames
	w.Name = user

	log.Infof(ctx, "Username SUBMITTED: %v", w.Name)

	// save word into the datastore.
	key := datastore.NewKey(ctx, "Dictionary", w.Name, 0, nil)
	_, err := datastore.Put(ctx, key, &w)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}
