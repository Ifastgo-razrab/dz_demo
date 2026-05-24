package storage

import (
	"CLI/bins"
	"encoding/json"
	"fmt"
)
type Bd interface{
	ReadFile()([]byte , error)
	Writefile([]byte)
}

type StorageBin struct{
	Bins []bins.Bin `json:"bins"`
	db Bd
}

func (s *StorageBin) NiwBin(bin bins.Bin){
	s.Bins =append(s.Bins, bin)
}

func (Sbin *StorageBin ) NiwBinJson(){
	data , err := json.Marshal(Sbin)
	if err != nil{
		fmt.Println("File :", err)
	}
	Sbin.db.Writefile(data)
	
}

func(a *StorageBin) ReadBinJson(){
	
	data ,  err := a.db.ReadFile()
	if err!= nil {
		fmt.Println("File:" , err)
	}
	var Sbin StorageBin
	err = json.Unmarshal(data , &Sbin)
	if err != nil{
		fmt.Println("File:q" , err)
	}
	for _ , data := range Sbin.Bins{
		fmt.Println(data)
	}

}
