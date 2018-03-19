// copies Source file src to destination dst.
package fileMove

import (
	"log"
	"os"
	"path/filepath"

	gh "github.com/ovace/utils/fileHash"
)

func MoveFile(dst string, src string) (valid bool, err error) {

	//check if target directory exists. If not create it with full path

	log.Printf("src: %v\t dst: %v\t Dest Dir: %v\n ", src, dst, filepath.Dir(dst))

	if _, err := os.Stat(filepath.Dir(dst)); os.IsNotExist(err) {
		os.MkdirAll(filepath.Dir(dst), os.ModePerm)
	}

	// calcualate the hash of the source file
	sHash, err := gh.FileHash(src)
	if err != nil {
		log.Fatalf("error calculating hash of src file: src: %v\t err: %v\n", src, err)
		return false, err
	}

	err = os.Rename(src, dst)
	if err != nil {
		log.Fatalf("error moving file: src: %v\t dst: %v\t err: %v\n", src, dst, err)
		return false, err
	}

	//not needed as this is just a rename, hence original perms are preserved
	/* si, err := os.Stat(src)
	if err == nil {
		err = os.Chmod(dst, si.Mode())
		if err != nil {
			log.Error("error changing perms:"+ err)
		} */
	dHash, err := gh.FileHash(dst)
	if err != nil {
		log.Fatalf("error calculating hash of dest file: dst: %v\t err: %v\n", dst, err)
		return false, err
	}
	if sHash != dHash {
		log.Fatalf("source and destination mismatch: sHash: %v\t dHash: %v\t err: %v\n", sHash, dHash, err)
		return false, err
	}
	return true, err
	//	}
	//}
	//return false,  err
}
