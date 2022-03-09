package piyolize

import (
	"fmt"
	"strings"
	"sync"

	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

var (
	once            sync.Once
	kagomeTokenizer *tokenizer.Tokenizer
)

func init() {
	once.Do(func() {
		tok, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
		if err != nil {
			panic(fmt.Errorf("unexpected error: %w", err))
		}
		kagomeTokenizer = tok
	})
}

func Do(msg string) string {
	tokens := tokenize(msg)

	var (
		buf          strings.Builder
		isLastSymbol bool
	)

	for _, token := range tokens {
		features := token.Features()
		key := key{
			pos:   features[0],
			affix: features[1],
			want:  token.Surface,
		}
		if replace, ok := dict[key]; ok {
			buf.WriteString(replace)
		} else {
			buf.WriteString(token.Surface)
		}
		isLastSymbol = features[0] == "記号"
	}
	if isLastSymbol {
		return buf.String()
	}
	return buf.String() + "ぴよぴよ🐥"
}

func tokenize(msg string) []tokenizer.Token {
	normalized := normalize(msg)
	return kagomeTokenizer.Tokenize(normalized)
}

func normalize(msg string) string {
	cases := []struct {
		old string
		new string
	}{
		{
			old: "。",
			new: "",
		},
		{
			old: "す？",
			new: "すか？",
		},
		{
			old: "ですか",
			new: "ですか？",
		},
		{
			old: "でしょうか",
			new: "でしょうか？",
		},
		{
			old: "ですかね",
			new: "ですかね？",
		},
		{
			old: "じゃないの",
			new: "じゃない？",
		},
	}
	for _, c := range cases {
		msg = normalizeLastPhrase(msg, c.old, c.new)
	}
	return msg
}

func normalizeLastPhrase(msg, old, new string) string {
	if !strings.HasSuffix(msg, old) {
		return msg
	}
	idx := strings.LastIndex(msg, old)
	if idx < 0 {
		return msg
	}
	return msg[:idx] + strings.Replace(msg[idx:], old, new, 1)
}
