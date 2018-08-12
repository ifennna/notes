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
	note string
)

func printFlagOptions() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(1)
}

func (env *Env) saveNote(note string){
	fmt.Printf("Your note is %s\n", note)
}

func main() {
	flag.Parse()

	db, err := models.NewDB("user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Panic(err)
	}

	env := &Env{db}

	if flag.NFlag() == 0 {
		printFlagOptions()
	}

	env.saveNote(note)
}


func init() {
	flag.StringVarP(&note, "jot", "j", "", "WriteNoteDown")
}
