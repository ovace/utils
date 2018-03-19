package databaselayer

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongodbHandler struct {
	*mgo.Session
}

func NewMongodbHandler(connection string) (*MongodbHandler, error) {
	s, err := mgo.Dial(connection)
	return &MongodbHandler{
		Session: s,
	}, err
}

func (handler *MongodbHandler) GetAvailableFiles() ([]Filelist, error) {
	s := handler.getFreshSession()
	defer s.Close()
	filelists := []Filelist{}
	err := s.DB("mediaDB").C("filelist").Find(nil).All(&filelists)
	return filelists, err
}

func (handler *MongodbHandler) GetFileByName(filename string) (Filelist, error) {
	s := handler.getFreshSession()
	defer s.Close()
	f := Filelist{}
	err := s.DB("mediaDB").C("filelist").Find(bson.M{"file_name": filename}).One(&f)
	return f, err
}

func (handler *MongodbHandler) GetFilesByType(fileSufix string) ([]Filelist, error) {
	s := handler.getFreshSession()
	defer s.Close()
	filelists := []Filelist{}
	err := s.DB("mediaDB").C("filelist").Find(bson.M{"file_suffix": fileSufix}).All(&filelists)
	return filelists, err
}

func (handler *MongodbHandler) AddFile(f Filelist) error {
	s := handler.getFreshSession()
	defer s.Close()
	return s.DB("mediaDB").C("filelist").Insert(f)
}

func (handler *MongodbHandler) UpdateFile(f Filelist, fname string) error {
	s := handler.getFreshSession()
	defer s.Close()
	return s.DB("mediaDB").C("filelist").Update(bson.M{"file_name": fname}, f)
}

func (handler *MongodbHandler) getFreshSession() *mgo.Session {
	return handler.Session.Copy()
}
