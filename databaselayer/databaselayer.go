package databaselayer

import (
	"errors"
	"time"
)

const (
	MYSQL uint8 = iota
	SQLITE
	POSTGRESQL
	MONGODB
)

type DinoDBHandler interface {
	GetAvailableDynos() ([]Animal, error)
	GetDynoByNickname(string) (Animal, error)
	GetDynosByType(string) ([]Animal, error)
	AddAnimal(Animal) error
	UpdateAnimal(Animal, string) error
}

type MediaDBHandler interface {
	GetAvailableFiles() ([]Filelist, error)
	GetFileByName(string) (Filelist, error)
	GetFilesByType(string) ([]Filelist, error)
	AddFile(Filelist) error
	UpdateFile(Filelist, string) error
	//DeleteFile(Filelist, string) error
	//CreateFilelist(Filelist, string) error
}

type Animal struct {
	ID         int    `bson:"-"`
	AnimalType string `bson:"animal_type"`
	Nickname   string `bson:"nickname"`
	Zone       int    `bson:"zone"`
	Age        int    `bson:"age"`
}

type Filelist struct {
	Fileid            int       `bson:"-"`
	Filename          string    `bson:"file_name"`
	Filesuffix        string    `bson:"file_suffix"`
	Filelocation      string    `bson:"file_loc"`
	Filesize          float32   `bson:"file_siz"`
	Filehash          string    `bson:"file_hash"`
	Filedate          time.Time `bson:"file_date"`
	Rowaction         string    `bson:"row_action"`
	Rowactiondatetime time.Time `bson:"row_ts"`
}

var DBTypeNotSupported = errors.New("The Database type provided is not supported...")

//factory function
func GetDatabaseHandler(dbtype uint8, connection string) (MediaDBHandler, error) {

	switch dbtype {
	case MYSQL:
		return NewMySQLHandler(connection)
	case MONGODB:
		return NewMongodbHandler(connection)
	case SQLITE:
		return NewSQLiteHandler(connection)
	case POSTGRESQL:
		return NewPQHandler(connection)
	}
	return nil, DBTypeNotSupported
}
