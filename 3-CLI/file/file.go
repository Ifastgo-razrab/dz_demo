package file

import (
	"fmt"
	"os"
	"strings"
)

type Dbfile struct{
	Name string
}
func NewDbFile(name string)*Dbfile{
	return &Dbfile{
		Name: name,
	}
}


func (db *Dbfile)Writefile(content []byte){
	data, err := os.Create(db.Name)
	if err != nil {
		fmt.Println("Файл :" , err )
	}
	defer data.Close()
	_ , err = data.Write(content)
	if err != nil {
		fmt.Println("Файл :" , err )
	}
	fmt.Println("Файл успешно записан")
}

func (db *Dbfile)ReadFile()([]byte , error){
	data , err := os.ReadFile(db.Name)
	if err != nil{
		fmt.Println("File:", err)
		
	}
	return data , err
}

func (db *Dbfile)Isjson()(bool){
	return strings.HasSuffix(db.Name , ".json")
}

