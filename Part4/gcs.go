package finalWeb

import (
	"crypto/sha1"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
	"google.golang.org/cloud/storage"
	"io"
	"net/http"
	"strings"
	"fmt"
)

// upload a file to google cloud
func uploadFile(ctx context.Context, res http.ResponseWriter, req *http.Request, ui *userInformation) string {

	ulFile, hdr, err := req.FormFile("data")												// get requested file.
	if err != nil {
		log.Errorf(ctx, "ERROR uploadPhoto req.FormFile: %s", err)
		return "Download Error"
	}
	defer ulFile.Close()
	ext := hdr.Filename[strings.LastIndex(hdr.Filename, ".")+1:]							// extract file extension.

	switch ext {																			// check for acceptable file types.
	case "jpg", "jpeg", "txt", "pdf":
		log.Infof(ctx, "GOOD FILE EXTENSION: %s", ext)
	default:
		log.Errorf(ctx, "We do not allow files of type %s. We only allow jpg, jpeg, txt or pdf extensions.", ext)
		return "Invalid file type"
	}
	h := sha1.New()																			// get sha value of file contents.
	io.Copy(h, ulFile)
	name := fmt.Sprintf("%x", h.Sum(nil)) + `.` + ext										// use result as file name + original file extension.
	gcsName := ui.UserId + "/" + name														// prefix user's uuid when being stored in GCS.
	ulFile.Seek(0, 0)

	client, err := storage.NewClient(ctx)													// create client object to write file to GCS.
	if err != nil {
		log.Errorf(ctx, "ERROR uploadPhoto storage.NewClient: %s", err)
		return "GCS storage error."
	}
	defer client.Close()
	writer := client.Bucket(gcsBucket).Object(gcsName).NewWriter(ctx)						// get GCS writer, set permissions for all users to reader.
	writer.ACL = []storage.ACLRule{
		{storage.AllUsers, storage.RoleReader},
	}
	io.Copy(writer, ulFile)																	// write file to GCS.
	err = writer.Close()
	if err != nil {
		log.Errorf(ctx, "ERROR uploadPhoto writer.Close: %s", err)
		return "Error writing to storage."
	}
	ui.Files = append(ui.Files, name)														// add file to user's list.
	// update session
	setUserInformationDatastore(ctx, *ui, req)
	setUserInformationMemcache(ctx, *ui, req)

	return ""
}
