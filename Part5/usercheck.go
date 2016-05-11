package finalWeb

import (
	// 	"google.golang.org/appengine"
	// 	"google.golang.org/appengine/datastore"
	// 	"google.golang.org/appengine/log"
	"io"
	// 	"io/ioutil"
	"net/http"
)

func usernameCheck(res http.ResponseWriter, req *http.Request) {
	// 	ctx := appengine.NewContext(req)
	//
	// 	// retrieve the incoming word as it is typed.
	// 	var w Word
	// 	bs, err := ioutil.ReadAll(req.Body)
	// 	//
	// 	log.Infof(ctx, "Received information: %v", string(bs))
	// 	//
	// 	if err != nil {
	// 		log.Infof(ctx, err.Error())
	// 	}
	// 	w.Name = string(bs)
	// 	log.Infof(ctx, "ENTERED wordCheck - w.Name: %v", w.Name)
	//
	// 	// check the incoming word against what is currently in the datastore
	// 	key := datastore.NewKey(ctx, "Dictionary", w.Name, 0, nil)
	// 	err = datastore.Get(ctx, key, &w)
	// 	if err != nil {
	if checkUserExists(res, req) {
		io.WriteString(res, "false")
		return
	}
	io.WriteString(res, "true")
}
