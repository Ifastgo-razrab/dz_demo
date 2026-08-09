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
	Db Bd
}

func (s *StorageBin) NiwBin(bin bins.Bin){
	s.Bins =append(s.Bins, bin)
}

func (Sbin *StorageBin ) NiwBinJson(){
	data , err := json.Marshal(Sbin)
	if err != nil{
		fmt.Println("File :", err)
	}
	Sbin.Db.Writefile(data)
	
}

func(a *StorageBin) ReadBinJson(){
	
	data ,  err := a.Db.ReadFile()
	if err != nil {
		fmt.Println("File:" , err)
	}
	err = json.Unmarshal(data , a)
	
	if err != nil{
		fmt.Println("File:q" , err)
	}
}
func (s *StorageBin) WriteBinJson(bins bins.Bin){
	s.ReadBinJson()
	s.NiwBin(bins)
	s.NiwBinJson()
}
