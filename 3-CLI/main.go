package main

import (
	"CLI/bins"
	"CLI/file"
	"CLI/storage"
)

func main() {

	bin := bins.BinInsal("", true, 1, "")

	db := file.NewDbFile("data.json")

	storagebin := storage.StorageBin{
		Db: db,
	}

	storagebin.NiwBin(bin)
	storagebin.NiwBinJson()
}