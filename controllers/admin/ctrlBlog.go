package admin

import (
	"fmt"
	"frz/models"
	"frz/tools"
	"os"

	"github.com/gin-gonic/gin"
)

func UploadBlogInfo(ctx *gin.Context) {
	image, err := ctx.FormFile("file")
	// fmt.Printf("image: %v\n", image)
	if err != nil {
		tools.Fail(ctx, gin.H{}, "未知错误")
		return
	}
	ctx.SaveUploadedFile(image, fmt.Sprintf("./static/image/%s", image.Filename))
	// file, _ := image.Open()
	// defer file.Close()
	// var buf [4]byte
	// var content []byte
	// content = append(content, 137, 80, 78, 71)
	// file.Read(buf[:])
	// for {
	// 	n, err := file.Read(buf[:])
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Printf("err: %v\n", err)
	// 		break
	// 	}
	// 	content = append(content, buf[:n]...)
	// }
	// webpct := SaveImage(image.Filename, content)
	// SaveImage(file.Filename,file.content)
	blog := models.Blog{
		Title:    ctx.PostForm("title"),
		Text:     ctx.PostForm("text"),
		Url:      ctx.PostForm("url"),
		Avatar:   ctx.PostForm("avatar"),
		Examine:  "yes",
		Webpct:   fmt.Sprintf("%s/static/image/%s", models.DOMAIN, image.Filename),
		Category: ctx.PostForm("category"),
	}
	models.DB.Create(&blog)
}

// 保存图片
func SaveImage(name string, b []byte) string {
	file, err := os.Create(fmt.Sprintf("./static/image/%s", name))
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	_, err = file.Write(b)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return fmt.Sprintf("%s/static/image/%s", models.DOMAIN, name)
}
