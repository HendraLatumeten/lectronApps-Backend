package middleware

import (
	"context"
	"os"
	"time"

	"github.com/HendraLatumeten/lectronApss-Backend/src/libs"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

func Cloudinary() gin.HandlerFunc {
	return func(c *gin.Context) {
		//**parse multipart/form-data
		if isErr := c.Request.ParseMultipartForm(20 << 20); isErr != nil {
			libs.New(isErr.Error(), 400, true).Send(c)
			c.Abort()
		}

		file, handlerFile, err := c.Request.FormFile("image")
		if err != nil {
			libs.New(err.Error(), 400, true)
			c.Abort()
		}
		defer file.Close()

		cntx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		//**file validation
		checkType := handlerFile.Header.Get("Content-Type") == "image/jpeg" || handlerFile.Header.Get("Content-Type") == "image/jpg" || handlerFile.Header.Get("Content-Type") == "image/png"
		if !checkType {
			libs.New(err.Error(), 400, true).Send(c)
			c.Abort()
		}

		//Format name
		name := time.Now().Format("2006-01-02_15:04:05") + "_" + handlerFile.Filename

		cld, err1 := cloudinary.NewFromParams(os.Getenv("CLOUD_NAME"), os.Getenv("CLOUD_KEY"), os.Getenv("CLOUD_SEC"))
		if err1 != nil {
			libs.New(err1.Error(), 400, true).Send(c)
			c.Abort()
		}

		upload, err2 := cld.Upload.Upload(cntx, file, uploader.UploadParams{Folder: "image", PublicID: name})
		if err2 != nil {
			libs.New(err2.Error(), 400, true).Send(c)
			c.Abort()
		}

		c.Set("image", upload.SecureURL)
		c.Next()
	}
}
