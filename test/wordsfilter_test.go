package test

import (
	"github.com/jeffcail/ginframe/server-common/utils/wordsfilter"
	"log"
	"testing"
)

func Init() {
	filepath := "../utils/wordsfilter/words_filter.txt"
	wordsfilter.SetWordsFilter(filepath)
}

func TestContentFilter(t *testing.T) {
	Init()
	//fmt.Println(words)
	if wordsfilter.ContentFilter("SB") {
		log.Fatal("SB 是敏感词")
	}
	log.Fatal("SB 不是敏感词")
}
