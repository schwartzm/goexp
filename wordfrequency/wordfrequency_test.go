package wordfrequency

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func TestGetWordFrequency(t *testing.T) {
	input := bytes.NewBufferString(
		`hello WORLD Hello so world i at 
		World hello cats HELLO cats dogs`)

	/* NOTE: DeepEqual nuance: Can't compare [4]Word with
	 * []Word, because Array != Slice.
	 * See https://play.golang.org/p/aDPIwjDq5bJ
	 * Also, order of expect slice is important. Keep as is.
	 */
	expect := []Word{Word{"hello", 4}, Word{"world", 3},
		Word{"cats", 2}, Word{"dogs", 1}}

	actual, err := GetWordFrequency(input, SORT_DESC, 3)

	if err != nil {
		fmt.Println(err)
	}

	if !reflect.DeepEqual(expect, actual.Words) {
		t.Errorf("Expected %v; got %v\n", expect, actual.Words)
	}
}
