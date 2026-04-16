package main


import (
	"fmt"
	"time"
	"CLI/bins"
	"CLI/api"
	"CLI/file"
	"CLI/storage"

)


type Bin struct{
	id string
	private bool
	createdAT time.Time
	name string
}

func BinInsal(id string , privatee bool, createdAT time.Time , name string)(Bin){
	return Bin{
		id : id,
		private: privatee,
		createdAT: createdAT,
		name: name,
	}
}


func main(){
	BinList := []Bin{}
	fmt.Println(BinList)

}