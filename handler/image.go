package handler

import (
	"github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

//Update : Update image
func Update(ctx *gin.Context) { //use more updates
	img, err := imgio.Open("/tmp/" + ctx.Param("image"))
	if err == nil {
		inverted := effect.Invert(img)
		// resized := transform.Resize(inverted, 800, 800, transform.Linear)
		// rotated := transform.Rotate(resized, 45, nil)
		if (strings.Split(ctx.Param("image"), ".")[len(strings.Split(ctx.Param("image"), "."))-1] == "jpeg") || (strings.Split(ctx.Param("image"), ".")[len(strings.Split(ctx.Param("image"), "."))-1] == "jpg") {
			if err := imgio.Save("/tmp/"+ctx.Param("image"), inverted, imgio.JPEG); err != nil {
				panic(err)
			}
		} else if strings.Split(ctx.Param("image"), ".")[len(strings.Split(ctx.Param("image"), "."))-1] == "png" {
			if err := imgio.Save("/tmp/"+ctx.Param("image"), inverted, imgio.PNG); err != nil {
				panic(err)
			}
		} else {
			if err := imgio.Save("/tmp/"+ctx.Param("image"), inverted, imgio.BMP); err != nil {
				panic(err)
			}
		}
		ctx.JSON(http.StatusOK, gin.H{
			"Name":    ctx.Query("image"),
			"Message": "Image successfully updated.",
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Data":    err,
			"Message": "Something went wrong.",
		})
	}
}

//Create : Create image
func Create(ctx *gin.Context) {
	// fmt.Println(c.Request.MultipartForm)
	file, header, err := ctx.Request.FormFile("file")
	if err == nil {
		imagename := uuid.NewV1().String() + "." + strings.Split(header.Filename, ".")[len(strings.Split(header.Filename, "."))-1]
		out, err := os.Create("/tmp/" + imagename)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"Name":    imagename,
			"Message": "Image successfully saved.",
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Data":    err,
			"Message": "Something went wrong.",
		})
	}
}

//Het : Get image
func Get(ctx *gin.Context) {
	http.ServeFile(ctx.Writer, ctx.Request, "/tmp/"+ctx.Param("image"))
}
