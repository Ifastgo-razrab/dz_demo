package main

import (
	"CLI/api"
	"CLI/bins"
	"fmt"
	"CLI/file"
	"CLI/storage"
	"flag"
)



func main() {
	create := flag.Bool("create", false, "Изменение")
	updates := flag.Bool("update", false, "Обновление")
	delet := flag.Bool("delete", false, "Удаление")
	get := flag.Bool("get", false, "Вывод")
	flagList := flag.Bool("list", false, "Вывод из лакального файла")
	name := flag.String("name", "", "")
	file := flag.String("file", "", "")
	Idstring := flag.String("id", "", "")

	flag.Parse()
	fmt.Printf("flag name = %q\n", *name)
	switch{
		case *create:
			if *name == "" || *file == "" {
				return
			}
			creates(*name, *file)
		case *updates:
			if *Idstring == "" || *file == "" {
				return
			}
			update(*file, *Idstring)
		case *delet:
			if *Idstring == "" || *file == "" {
				return
			}
			delete(*file, *Idstring)
		case *get:
			if *Idstring == "" {
				return
			}
			api.Binget(*Idstring, "https://api.jsonbin.io/v3/b/")
		case *flagList:
			list()

		
	}
}
 
func creates(name string, files string)(){
	db := file.NewDbFile(files)
	storagebin := storage.StorageBin{
		Db: db,
	}
	
	storagebin.ReadBinJson()
	
	for _ , data := range storagebin.Bins{
		if data.Name == name {
			fmt.Println("Нашел")
			api.CreatBin(data)
		}
	}
}
func update(files string, Id string){
	db := file.NewDbFile(files)
	storagebin := storage.StorageBin{
		Db: db,
	}
	
	storagebin.ReadBinJson()
	for _ , data := range storagebin.Bins{
		if data.Id == Id {
			fmt.Println(data)
			api.Put(data, Id)
		}
	}
}

func delete(files string, Id string){
	api.Delite(Id, "https://api.jsonbin.io/v3/b/")
	db := file.NewDbFile(files)
	storagebin := storage.StorageBin{
		Db: db,
	}
	noDelete:=  []bins.Bin{}
	storagebin.ReadBinJson()
	for _ , data := range storagebin.Bins{
		if data.Id != Id {
				noDelete = append(noDelete, data)
		}
	}
	storagebin.Bins = noDelete
	storagebin.NiwBinJson()
}

func list(){
	db := file.NewDbFile("data.json")
	storagebin := storage.StorageBin{
		Db: db,
	}
	
	storagebin.ReadBinJson()
	fmt.Print(storagebin.Bins)
	
}