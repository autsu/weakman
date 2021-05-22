package test

import (
	"testing"
)

func TestTopicInsertRequestJson(t *testing.T) {
	testJson := `
{
    "title": "阿斯顿撒大",
    "description": "描述test",
    "deadline": "2022-07-31T14:27:10.035542+08:00",
	"selectType": 1,
	"anonymous":  1,
	"showResult": 1,
	"password":   "",
    "option": [
        {"optionContent": "选项1"},
        {"optionContent": "选项2"}
    ]
}
`
	_ = testJson
}
