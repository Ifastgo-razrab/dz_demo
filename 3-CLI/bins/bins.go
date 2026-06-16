package bins
//import "time"

type Bin struct{
	Id string `json:"id"`
	Private bool `json:"private"`
	CreatedAT int `json:"createdat"`
	Name string	`json:"name"`
}

func BinInsal(id string , privatee bool, createdAT int , name string)(Bin){
	return Bin{
		Id : id,
		Private: privatee,
		CreatedAT: createdAT,
		Name: name,
	}
}