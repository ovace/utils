//This function reads the File System time of a fie
package fileStats

// This utility copies photos into a date-specific folder

import (
	"log"
	"time"
)
import (
	"github.com/djherbis/times"
)

func StatFsTimes(fname string) (atime, mtime, ctime time.Time, err error) {
	/*fi, err := os.Stat(name)
	if err != nil {
		return
	}
	mtime = fi.ModTime()

	stat := fi.Sys().(*syscall.Stat_t)
	atime = time.Unix(int64(stat.Atim.Sec), int64(stat.Atim.Nsec))
	ctime = time.Unix(int64(stat.Ctim.Sec), int64(stat.Ctim.Nsec))*/
	t, err := times.Stat("fname")
	if err != nil {
		log.Fatal(err.Error())
	}

	atime = t.AccessTime()
	mtime = t.ModTime()

	if t.HasChangeTime() {
		ctime = t.ChangeTime()
	}
	return
}
