//Get the Exif details of the media file
package fileStats

import (
	"log"
)
import (
	gef "github.com/ovace/pkg/goexiftool"
)

func StatExifTags(path string) (map[string]interface{}, error) {
	fname := path

	//Get exif tags
	x, err := gef.NewImageTags(fname)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return x.Tags(), nil
}

func StatExifStruct(path string) (gef.EXIF, error) {
	fname := path

	//Get exif Info
	x, err := gef.NewImageEXIF(fname)
	if err != nil {
		log.Println(err)
		//return , err
	}

	return x, nil
}
