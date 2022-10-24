package libs

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

func CloudUpload(src string) (string, error) {
	cloud := os.Getenv("CLOUD_NAME")
	keys := os.Getenv("CLOUD_KEY")
	secret := os.Getenv("CLOUD_SEC")

	ctx := context.Background()
	imageId := uuid.New()

	cld, err := cloudinary.NewFromParams(cloud, keys, secret)
	if err != nil {
		return "", err
	}

	resp, err := cld.Upload.Upload(ctx, src, uploader.UploadParams{PublicID: imageId.String(), Folder: "image"})
	if err != nil {
		return "", err
	}

	if err := os.Remove(src); err != nil {
		log.Println(err.Error())
	}

	return resp.SecureURL, nil
}

func CloudDelete(src string) (string, error) {

	cloud := os.Getenv("CLOUD_NAME")
	keys := os.Getenv("CLOUD_KEY")
	secret := os.Getenv("CLOUD_SEC")

	ctx := context.Background()

	cld, err := cloudinary.NewFromParams(cloud, keys, secret)
	if err != nil {
		return "", err
	}

	//slice url
	s := strings.SplitAfter(src, "/")
	var name []string
	for _, element := range s {
		name = append(name, element)
	}

	//note
	publicID := strings.Join(strings.Split(name[8], ".jpg"), " ")

	resp, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID:     "image/" + publicID,
		ResourceType: "image"})

	return resp.Error.Message, nil
}
