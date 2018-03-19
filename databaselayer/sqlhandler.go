package databaselayer

import (
	"database/sql"
	"fmt"
	"log"
)

type SQLHandler struct {
	*sql.DB
}

func (handler *SQLHandler) GetAvailableFiles() ([]Filelist, error) {
	return handler.sendQuery("select * from filelist")
}

func (handler *SQLHandler) GetFileByName(filename string) (Filelist, error) {

	row := handler.QueryRow(fmt.Sprintf("select * from filelist where filename = '%s'", filename)) //? for mysql or sqlite and it used to be $1 for pq
	f := Filelist{}
	err := row.Scan(&f.Fileid, &f.Filename, &f.Filesuffix, &f.Filelocation, &f.Filesize, &f.Filehash, &f.Filedate, &f.Rowaction, &f.Rowactiondatetime)
	return f, err
}

func (handler *SQLHandler) GetFilesByType(fileType string) ([]Filelist, error) {
	return handler.sendQuery(fmt.Sprintf("select * from filelist where filesuffix = '%s'", fileType))
}

func (handler *SQLHandler) AddFile(f Filelist) error {
	sqlStr := fmt.Sprintf("Insert into filelist (filename,filesuffix,filelocation,filesize,filehash,filedate,rowaction) values ('%s','%s','%s',%v,'%s','%v','%s')", f.Filename, f.Filesuffix, f.Filelocation, f.Filesize, f.Filehash, f.Filedate, f.Rowaction)
	log.Println(sqlStr)
	_, err := handler.Exec(sqlStr)
	return err
}
func (handler *SQLHandler) UpdateFile(f Filelist, fname string) error {
	_, err := handler.Exec(fmt.Sprintf("Update filelist set filesuffix = '%s' ,filename = '%s',filelocation = '%s',filedate = '%v' where filename = '%s'", f.Filesuffix, f.Filename, f.Filelocation, f.Filedate, fname))
	return err
}

func (handler *SQLHandler) sendQuery(q string) ([]Filelist, error) {
	Filelists := []Filelist{}
	rows, err := handler.Query(q)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		f := Filelist{}
		err := rows.Scan(&f.Fileid, &f.Filename, &f.Filesuffix, &f.Filelocation, &f.Filesize, &f.Filehash, &f.Filedate, &f.Rowaction, &f.Rowactiondatetime)
		if err != nil {
			log.Println(err)
			continue
		}
		Filelists = append(Filelists, f)
	}
	return Filelists, rows.Err()
}

/*
	&f.Fileid,&f.Filename,&f.Filesuffix,&f.Filelocation,&f.Filesize,&f.Filehash,&f.Filedate,&f.Rowaction,&f.Rowactiondatetime
*/
