package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Any("/form", func(c *gin.Context) {

		log.Printf("[form:k1]=>%+v\n", c.Request.FormValue("k1"))
		form, err := c.MultipartForm()
		if err != nil {
			log.Printf("%s", err.Error())
			c.String(200, err.Error())
			return
		}

		files := form.File["files"]
		if len(files) <= 0 {
			files = form.File["files[]"]
		}
		for _, file := range files {
			log.Println(file.Filename, file.Header.Get("Content-Type"))
			if err := c.SaveUploadedFile(file, file.Filename); err != nil {
				log.Printf("[save]=>%s\n", err.Error())
			}
		}

		c.String(200, string("ok"))
	})
	r.Run(":8000")
}
