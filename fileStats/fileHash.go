//generates the hash for a given file
package fileStats

// This utility copies photos into a date-specific folder

import (
	"hash/crc32"
	"io/ioutil"
)

const form = "2006:01:02 15:04:05"
const folderFormat = "2006/01"

type pathEntry struct {
	filename string
}

// pathWalker walks the filesystem, queueing pathEntry items onto the queue.
type pathWalker struct {
	MyCounters
	queue chan pathEntry
}
type moveEntry struct {
	source   string
	dest     string
	filehash uint32

	//	dupcnt   uint32
}
type fileMover struct {
	sourcePath string
	destPath   string
	isCopy     bool
	dryRun     bool
	MyCounters
	queue chan moveEntry
}
type MyCounters struct {
	readDirCounter   int
	readFileCounter  int
	writeDirCounter  int
	writeFileCounter int
	dupCounter       int
	failCounter      int
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
