package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseTagList(t *testing.T) {
	contents, err := ioutil.ReadFile("taglist_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseTagList(contents)
	const resultSize = 145
	exceptedTags := []string{"小说", "外国文学", "文学"}
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Requests))
	}
	if len(result.Item) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Item))
	}

	for i, v := range exceptedTags {
		if result.Item[i].(string) != v {
			t.Errorf("except item #%d: %s; but was %s", i, v, result.Item[i])
		}
	}

}
