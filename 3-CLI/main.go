package main

import (
	"CLI/bins"
	"CLI/storage"
)








func main(){
	bin := bins.BinInsal("", true , 1 , "")

	staragebin := storage.StorageBin{}

	staragebin.NiwBin(bin)


	
}