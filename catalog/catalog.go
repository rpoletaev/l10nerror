// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package main

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type dictionary struct {
	index []uint32
	data  string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	p, ok := messageKeyToIndex[key]
	if !ok {
		return "", false
	}
	start, end := d.index[p], d.index[p+1]
	if start == end {
		return "", false
	}
	return d.data[start:end], true
}

func init() {
	dict := map[string]catalog.Dictionary{
		"en": &dictionary{index: enIndex, data: enData},
		"ru": &dictionary{index: ruIndex, data: ruData},
	}
	fallback := language.MustParse("en")
	cat, err := catalog.NewFromMap(dict, catalog.Fallback(fallback))
	if err != nil {
		panic(err)
	}
	message.DefaultCatalog = cat
}

var messageKeyToIndex = map[string]int{
	"%.2[1]f miles traveled (%[1]f)": 4,
	"%s is visiting %s!":             3,
	"Hello":                          1,
	"Hello world!":                   0,
	"Wrong parameter: %s":            2,
}

var enIndex = []uint32{ // 6 elements
	0x00000000, 0x0000000d, 0x00000013, 0x0000002a,
	0x00000043, 0x00000062,
} // Size: 48 bytes

const enData string = "" + // Size: 98 bytes
	"\x02Hello world!\x02Hello\x02Wrong parameter: %[1]s\x02%[1]s is visiting" +
	" %[2]s!\x02%.2[1]f miles traveled (%[1]f)"

var ruIndex = []uint32{ // 6 elements
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000,
} // Size: 48 bytes

const ruData string = ""

// Total table size 194 bytes (0KiB); checksum: 19A8AC03