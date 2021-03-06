package finalWeb

import (
	"crypto/sha1"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
	"google.golang.org/cloud/storage"
	"io"
	"net/http"
	"strings"
)

// upload a file to google cloud
func uploadFile(ctx context.Context, res http.ResponseWriter, req *http.Request, ui *userInformation) string {

	ulFile, hdr, err := req.FormFile("data") // get requested file.
	if err != nil {
		log.Errorf(ctx, "ERROR uploadPhoto req.FormFile: %s", err)
		return "Download Error"
	}
	defer ulFile.Close()
	ext := hdr.Filename[strings.LastIndex(hdr.Filename, ".")+1:] // extract file extension.

	switch ext { // check for acceptable file types.
	case "jpg", "jpeg", "txt", "pdf":
		log.Infof(ctx, "GOOD FILE EXTENSION: %s", ext)
	default:
		log.Errorf(ctx, "We do not allow files of type %s. We only allow jpg, jpeg, txt or pdf extensions.", ext)
		return "Invalid file type"
	}
	h := sha1.New() // get sha value of file contents.
	io.Copy(h, ulFile)
	name := fmt.Sprintf("%x", h.Sum(nil)) + `.` + ext // use result as file name + original file extension.
	gcsName := ui.UserId + "/" + name                 // prefix user's uuid when being stored in GCS.
	ulFile.Seek(0, 0)

	client, err := storage.NewClient(ctx) // create client object to write file to GCS.
	if err != nil {
		log.Errorf(ctx, "ERROR uploadPhoto storage.NewClient: %s", err)
		return "GCS storage error."
	}
	defer client.Close()
	writer := client.Bucket(gcsBucket).Object(gcsName).NewWriter(ctx) // get GCS writer, set permissions for all users to reader.
	writer.ACL = []storage.ACLRule{
		{storage.AllUsers, storage.RoleReader},
	}
	io.Copy(writer, ulFile) // write file to GCS.
	err = writer.Close()
	if err != nil {
		log.Errorf(ctx, "ERROR uploadPhoto writer.Close: %s", err)
		return "Error writing to storage."
	}
	ui.Files = append(ui.Files, name) // add file to user's list.
	// update session
	setUserInformationDatastore(ctx, *ui, req)
	setUserInformationMemcache(ctx, *ui, req)

	return ""
}

// get user's file list from google cloud
func fileList(ctx context.Context, res http.ResponseWriter, req *http.Request, ui *userInformation) []dFile {
	var uFiles []dFile

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Errorf(ctx, "ERROR listBucket storage.NewClient: %v", err)
		return uFiles
	}
	defer client.Close()

	q := &storage.Query{
		Prefix: ui.UserId,
	}

	objs, err := client.Bucket(gcsBucket).List(ctx, q)
	if err != nil {
		log.Errorf(ctx, "ERROR listBucket client.Bucket: %v", err)
		return uFiles
	}

	for _, obj := range objs.Results {
		p := strings.Split(obj.Name, "/")
		var df = dFile{
			Name: p[len(p)-1],
			Link: obj.MediaLink,
			Type: obj.ContentType,
		}
		log.Infof(ctx, "File: %v    -- Type %v", df.Name, obj.ContentType)
		uFiles = append(uFiles, df)
	}
	return uFiles
}
