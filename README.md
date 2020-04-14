# notes
A command line note taking app built with Go

[![asciicast](https://asciinema.org/a/JMZxtuxEPfrbFrOmj7uaobqKt.svg)](https://asciinema.org/a/JMZxtuxEPfrbFrOmj7uaobqKt)

## Installation
To build, make sure you have [Go](https://golang.org/dl/) installed. Run `go build` to 
create a binary or `go install` to add it to your GOPATH.

## Usage
  notes [command]
  
Available Commands:
  - `add`: Add notes
    - `notes add [notebook] "my 1st note" "my 2nd note" ..`
    - if `notebook` name is not supplied, it is added to `Default` notebook
    - if `notebook` doesn't exist, new notebook is created
  - `help`: Help about any command
    - `notes help`
  - `ls`: List stuff
    - `notes ls [notebook]`
    - if `notebook` name is not supplied, names of notebooks are displayed
    - if `notebook` name is supplied
      - if notebook by given name exists, all notes of that notebook are displayed along with their `note_id`s
      - if notebook by given name doesn't exist, only the entered notebook name is shown in output (needs to be improved)
  - `del`: Delete notes
    - `notes del notebook note_id_1 note_id_2 ..`
    - if notebook by given name exists
      - if note by given note_id exists, it is deleted; and note deletion message is displayed
      - if note by given note_id doesn't exist, nothing happens. Note deleteion message still appears (needs to be fixed)
    - if notebook by given name doesn't exist
      - command terminates unexpectedly with stacktrace (needs to be fixed)
  - `rm`: Removes a notebook
    - `notes rm notebook`
    - if `notebook` doesn't exist, nothing changes

Use "notes [command] --help" for more information about a command.

