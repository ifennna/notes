package utils

import (
	"fmt"
	"strconv"
)

//ParseUInt64 parses single string into uint64
func ParseUInt64(arg string) (uint64, error) {
	noteID, err := strconv.ParseInt(arg, 10, 64)
	usNoteID := uint64(noteID)
	if err != nil {
		fmt.Println("Error while parsing args")
	}
	return usNoteID, err
}

//ParseUInt64Slice parses slice of strings into slice of uint64
func ParseUInt64Slice(args []string) ([]uint64, error) {
	var usNoteIds []uint64
	var usNoteID uint64
	var err error
	for _, arg := range args {
		usNoteID, err = ParseUInt64(arg)
		if err != nil {
			return usNoteIds, err
		}
		usNoteIds = append(usNoteIds, usNoteID)
	}
	return usNoteIds, err
}
