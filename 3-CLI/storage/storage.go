package storage

import (
	"CLI/bins"
	"CLI/file"
	"encoding/json"
	"fmt"
	"os"
)

type StorageBin struct{
	Bins []bins.Bin `json:"bins"`
}

func (s *StorageBin) NiwBin(bin bins.Bin){
	s.Bins =append(s.Bins, bin)
}

func (Sbin StorageBin) NiwBinJson(files string){
	data , err := json.Marshal(Sbin)
	if err != nil{
		fmt.Println("File :", err)
	}
	
	file.Writefile(data , files)
}

func ReadBinJson(file string){
	data ,  err := os.ReadFile(file)
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
