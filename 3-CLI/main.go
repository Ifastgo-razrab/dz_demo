package main


import (
	
)



type Bin struct{
	id string
	private bool
	createdAT int
	name string
}

func BinInsal(id string , privatee bool, createdAT int , name string)(Bin){
	return Bin{
		id : id,
		private: privatee,
		createdAT: createdAT,
		name: name,
	}
}

func main(){
	BinList := []Bin{}
	_ =BinList
}