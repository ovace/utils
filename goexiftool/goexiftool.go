package goexiftool

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os/exec"
)

type Image interface {
	Tags() map[string]interface{}
	Exif() EXIF
	String(name string) (string, error)
	StringSlice(name string) ([]string, error)
	AddTag(name string, value string) error
	RemoveTag(name string) error
	AddTagValue(name string, value string) error
	RemoveTagValue(name string, value string) error
}

// ImageCache holds all the metadata provided by exiftool in a map
type ImageCache struct {
	filepath string
	tags     map[string]interface{}
	exif     EXIF
}

func NewImageTags(filepath string) (Image, error) {

	/*
				-a
					Allow duplicate tag names in the output. Without this option, duplicates are suppressed.

				-f
					Force printing of tags even if their values are not found.

				-g[NUM]
					Organize output by tag group. NUM specifies the group family number, and may be 0 (general location), 1 (specific location) or 2 (category). If not specified, -g0 is assumed. Use the -listg option to list all group names for a specified family.

				-json
					Export/import tags in JSON format
				-n
					Read and write values as numbers instead of words. This option disables the print conversion that is applied when extracting values to make them more readable, and the inverse print conversion when writing.

				-S
					Very short format. The same as two -s options. Extra spaces used to column-align values are not printed.

				-U
					Extract values of unknown tags as well as unknown information from binary data blocks. This is the same as two -u options.

				-sep STR    (-separator)         Set separator string for list items
		  		-sort                            Sort output alphabetically
		  		-struct                          Enable output of structured information

				//exiftool -a -f -n -U -g -j -struct -r org > exif_out.json
	*/
	cmdOut, err := callExifTool("-j", "-struct", "-U", "-f", "-n", "-g", "-S", "-sort", filepath)
	//cmdOut, err := callExifTool("-json", "-struct", "-U", filepath)
	//cmdOut, err := callExifTool("-j", "-struct", "-U","l", filepath)
	//cmdOut, err := callExifTool("-j", filepath)

	if err != nil {
		return nil, err
	}

	var tags []map[string]interface{}

	if err := json.Unmarshal([]byte(cmdOut), &tags); err != nil {
		return nil, err
	}

	//	clean(tags[0])

	return &ImageCache{filepath: filepath, tags: tags[0]}, nil

}

func NewImageEXIF(filepath string) (ex EXIF, err error) {

	//cmdOut, err := callExifTool("-j", "-struct", "-U", filepath)
	//cmdOut, err := callExifTool("-j", "-struct", filepath)
	//cmdOut, err := callExifTool("-j", "-struct", "-U", "-f", "-n", "-g", "-S", "-sort", filepath)
	cmdOut, err := callExifTool("-j", "-struct", "-f", "-S", "-sort", filepath)
	if err != nil {
		return ex, err
	}
	log.Printf("image EXIF: %v\n ", cmdOut)

	var exf []EXIF

	if err := json.Unmarshal([]byte(cmdOut), &exf); err != nil {

		log.Printf("Encountered error while unmarshalling JSON. err: %v\n", err)
	}
	log.Printf("Unmarshalled EXIF JSON: %v\n", exf[0])
	return exf[0], nil
}

func (img *ImageCache) Tags() map[string]interface{} {
	return img.tags
}

func (img *ImageCache) Exif() EXIF {
	return img.exif
}

// String displays metadata for a given tag name
func (img *ImageCache) String(name string) (tagValue string, err error) {
	current, ok := img.tags[name]
	if !ok {
		err = errors.New("Unknown tag " + name)
	}

	if current == nil {
		return "", nil
	}

	switch v := current.(type) {
	default:
		return "", errors.New(fmt.Sprintf("unexpected tag type %T", v))
	case string:
		return current.(string), nil
	}
}

func (img *ImageCache) StringSlice(name string) ([]string, error) {
	current := img.tags[name]

	if current == nil {
		return []string{}, nil
	}

	switch v := current.(type) {
	default:
		return nil, errors.New(fmt.Sprintf("unexpected tag type %T", v))
	case string:
		return []string{current.(string)}, nil
	case []string:
		return current.([]string), nil
	}
}

func (img *ImageCache) AddTag(name string, value string) error {

	if name == "" {
		return errors.New("name required")
	}
	if value == "" {
		return errors.New("value required")
	}

	if img.tags[name] != nil {
		return errors.New(fmt.Sprintf("Tag %v already exists", name))
	}

	out, err := callExifTool(fmt.Sprintf("-%v=%v", name, value), img.filepath)

	if err != nil {
		return errors.New(fmt.Sprintf("%v: %v", err, out))
	}

	img.tags[name] = value

	return nil
}

func (img *ImageCache) RemoveTag(name string) error {

	if name == "" {
		return errors.New("name required")
	}

	if img.tags[name] == nil {
		return errors.New(fmt.Sprintf("Tag %v does not exist", name))
	}

	out, err := callExifTool(fmt.Sprintf("-%v=", name), img.filepath)

	if err != nil {
		return errors.New(fmt.Sprintf("%v: %v", err, out))
	}

	delete(img.tags, name)

	return nil
}

func (img *ImageCache) AddTagValue(name string, value string) error {

	if name == "" {
		return errors.New("name required")
	}
	if value == "" {
		return errors.New("value required")
	}

	current := img.tags[name]

	var vals []string

	if current == nil {
		vals = make([]string, 0)
	} else {

		switch v := current.(type) {
		default:
			return errors.New(fmt.Sprintf("unexpected tag type %T", v))
		case string:
			vals = []string{current.(string)}
		case []string:
			vals = current.([]string)
		}
	}

	out, err := callExifTool(fmt.Sprintf("-%v+=%v", name, value), img.filepath)

	if err != nil {
		return errors.New(fmt.Sprintf("%v: %v", err, out))
	}

	vals = append(vals, value)
	img.tags[name] = vals

	return nil
}

func (img *ImageCache) RemoveTagValue(name string, value string) error {

	if name == "" {
		return errors.New("name required")
	}
	if value == "" {
		return errors.New("value required")
	}

	current := img.tags[name]

	var vals []string

	if current == nil {
		return errors.New(fmt.Sprintf("Tag not found: %v", name))
	} else {
		switch v := current.(type) {
		default:
			return errors.New(fmt.Sprintf("unexpected tag type %T", v))
		case string:
			vals = []string{current.(string)}
		case []string:
			vals = current.([]string)
		}
	}

	out, err := callExifTool(fmt.Sprintf("-%v-=%v", name, value), img.filepath)

	if err != nil {
		return errors.New(fmt.Sprintf("%v: %v", err, out))
	}

	for i, v := range vals {
		if v == value {
			vals = append(vals[:i], vals[i+1:]...)
			break
		}
	}

	img.tags[name] = vals

	return nil
}

func clean(m map[string]interface{}) {
	for k, v := range m {
		if is, ok := v.([]interface{}); ok {
			ss := make([]string, len(is))
			for i, s := range is {
				ss[i] = s.(string)
			}
			m[k] = ss
		}
	}
}

func callExifTool(args ...string) (string, error) {
	cmdName, err := exec.LookPath("exiftool")
	if err != nil {
		return "nil", errors.New("exiftool is not installed")
	}
	//cmdOut, err := exec.Command(cmdName, args...).CombinedOutput()//combines stdout and stderr -- this causes problemn when unmarshalling JSON
	cmdOut, err := exec.Command(cmdName, args...).Output()
	//log.Printf("EXIF read2: %v\n ", cmdOut)
	//log.Printf("EXIF read3: %v\n ", string(cmdOut))
	return string(cmdOut), err
}
