package utils

import (
	"fmt"
	"strconv"
)

/**
 * Parses single string into uint64
 */
func ParseUInt64(arg string) (uint64, error) {
	noteId, err := strconv.ParseInt(arg, 10, 64)
	usNoteId := uint64(noteId)
	if err != nil {
		fmt.Println("Error while parsing args")
	}
	return usNoteId, err
}

/**
 * Parses slice of strings into slice of uint64
 */
func ParseUInt64Slice(args []string) ([]uint64, error) {
	var usNoteIds []uint64
	var usNoteId uint64
	var err error
	for _, arg := range args {
		usNoteId, err = ParseUInt64(arg)
		if err != nil {
			return usNoteIds, err
		} else {
			usNoteIds = append(usNoteIds, usNoteId)
		}
	}
	return usNoteIds, err
}
