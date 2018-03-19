//generates the hash for a given file
package fileHash


// This utility copies photos into a date-specific folder

import (
	"hash/crc32"
	"io/ioutil"
)

type pathEntry struct {
	filename string
}

func FileHash(filename string) (uint32, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	h := crc32.NewIEEE()
	h.Write(bs)
	return h.Sum32(), nil
}
