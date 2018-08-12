package main

import (
	flag "github.com/spf13/pflag"
	"fmt"
	"os"
	"github.com/uncultured/notes/models"
	"log"
	)

type Env struct{
	db models.Datastore
}

var(
	noteContent string
)

func printFlagOptions() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(1)
}

func (env *Env) saveNote(note models.Note){
	fmt.Printf("Your note is %s\n", note.Content)
	//err := env.db.AddNote(note)
	//note, err = env.db.GetNote(1)
	//if err != nil {
	//	log.Panic(err)
	//}
	//fmt.Printf("%s\n", note.Content)
}

func main() {
	flag.Parse()

	db, err := models.NewDB("notebooks.db")
	if err != nil {
		log.Panic(err)
	}

	env := &Env{db}

	if flag.NFlag() == 0 {
		printFlagOptions()
	}

	note := models.Note{Title: "Default", Content: noteContent}

	env.saveNote(note)
}


func init() {
	flag.StringVarP(&noteContent, "jot", "j", "", "Jot down a note")
}
