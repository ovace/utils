// copies Source file src to destination dst.
package fileCopy

import (
	"io"
	"log"
	"os"
	"path/filepath"

	gh "github.com/ovace/utils/fileHash"
)

func CopyFile(dst string, src string) (valid bool, n int64, err error) {

	//check if target directory exists. If not create it with full path

	log.Printf("src: %v\t dst: %v\t Dest Dir: %v\n ", src, dst, filepath.Dir(dst))

	if _, err := os.Stat(filepath.Dir(dst)); os.IsNotExist(err) {
		os.MkdirAll(filepath.Dir(dst), os.ModePerm)
	}

	// calcualate the hash of the source file
	sHash, err := gh.FileHash(src)
	if err != nil {
		log.Fatalf("error calculating hash of src file: src: %v\t err: %v\n", src, err)
		return false, n, err
	}

	sf, err := os.Open(src)
	if err != nil {
		return false, 0, err
	}
	defer sf.Close()
	df, err := os.Create(dst)
	if err != nil {
		return false, 0, err
	}
	defer df.Close()
	//return io.Copy(df, sf)
	n, err = io.Copy(df, sf)
	if err == nil {
		si, err := os.Stat(src)
		if err == nil {
			err = os.Chmod(dst, si.Mode())
			if err != nil {
				log.Printf("error changing perms: %v\n", err)
			}
			dHash, err := gh.FileHash(dst)
			if err != nil {
				log.Fatalf("error calculating hash of dest file: dst: %v\t err: %v\n", dst, err)
				return false, n, err
			}
			if sHash != dHash {
				log.Fatalf("error calculating hash of dest file: %v\n", err)
				return false, n, err

			}
			return true, n, err
		}
	}
	return false, n, err
}

/*
func main() {
	fn := "copyfile.go"
	n, err := CopyFile("(copy of) "+fn, fn)
	if err != nil {
		fmt.Println(n, err)
	}
}
*/
