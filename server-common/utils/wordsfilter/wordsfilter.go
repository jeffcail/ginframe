package wordsfilter

import (
	"bufio"
	"github.com/syyongx/go-wordsfilter"
	"log"
	"os"
)

var (
	Wf     *wordsfilter.WordsFilter
	words  []string
	master map[string]*wordsfilter.Node
)

// SetWordsFilter 项目启动的时候，通过此方法将敏感词库加载进去
func SetWordsFilter(filepath string) {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		words = append(words, s)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	Wf = wordsfilter.New()
	master = Wf.Generate(words)
}

// ContentFilter 需要过滤的地方，调用此方法，将要过滤的内容传递进去即可
func ContentFilter(text string) bool {
	return Wf.Contains(text, master)
}
