package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseTag(t *testing.T) {
	contents, err := ioutil.ReadFile("tag_test_data.html")
	if err != nil {
		panic(err)
	}
	const resultSize = 19
	result := ParseTag(contents)
	if len(result.Requests) != resultSize {
		t.Errorf("book size requests result should have %d requests; but had %d", resultSize, len(result.Requests))
	}
	if len(result.Items) != resultSize {
		t.Errorf("book size item result should have %d requests; but had %d", resultSize, len(result.Requests))
	}
	bookTestData := []string{"活着", "夜晚的潜水艇", "白夜行"}
	for i, v := range bookTestData {
		if result.Items[i] != v {
			t.Errorf("except item #%d: %s; but was %s", i, v, result.Items[i])
		}
	}

}
