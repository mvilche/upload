package controllers

import (
	"fmt"
	"os"

	"github.com/revel/revel"
)

type Single struct {
	App
}

func (c *Single) Upload() revel.Result {
	return c.Render()
}

func (c *Single) HandleUpload(avatar []byte) revel.Result {

	c.Validation.Required(avatar)

	// Handle errors.
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect((*Single).Upload)
	}

	f, err := os.Create("storage/" + c.Params.Files["avatar"][0].Filename)
	if err != nil {
		fmt.Println(err)
	}

	_, err2 := f.Write(avatar)

	if err2 != nil {
		fmt.Println(err)
	}

	return c.RenderJSON(FileInfo{
		ContentType: c.Params.Files["avatar"][0].Header.Get("Content-Type"),
		Filename:    c.Params.Files["avatar"][0].Filename,
		Size:        len(avatar),
		Status:      "Successfully uploaded",
	})
}
