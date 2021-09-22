package tests

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/noculture/notes/utils"
)

var testArgs = []string{"1", "2", "3"}

func TestParseUInt64Slice(t *testing.T) {
	var expected []uint64

	for _, arg := range testArgs {
		intArg, err := strconv.ParseInt(arg, 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		expected = append(expected, uint64(intArg))
	}

	got, err := utils.ParseUInt64Slice(testArgs)
	if err != nil {
		t.Errorf("err got %v, want nil", err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v, expected %v \n", got, expected)
	}

}
