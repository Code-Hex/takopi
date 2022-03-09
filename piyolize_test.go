package piyolize

import "testing"

func TestDo(t *testing.T) {
	tests := []struct {
		msg  string
		want string
	}{
		{
			msg:  "どうだ",
			want: "どうだぴよぴよ🐥",
		},
		{
			msg:  "はい、いいえ、どうだ",
			want: "そうぴよ🐥、ちがうぴよ🐥、どうだぴよぴよ🐥",
		},
		{
			msg:  "課題はもう終わった？",
			want: "課題はもう終わったぴよ🐥？",
		},
		{
			msg:  "共有はしてたけど",
			want: "共有はしてたけどぴよぴよ🐥",
		},
		{
			msg:  "もう行きます？",
			want: "もう行きますかぴよ🐥？",
		},
		{
			msg:  "一緒にご飯どうですか",
			want: "一緒にご飯どうですかぴよ🐥？",
		},
		{
			msg:  "一緒にご飯どうでしょうか",
			want: "一緒にご飯どうでしょうかぴよ🐥？",
		},
		{
			msg:  "ビルドができてればいいんじゃないんですかね",
			want: "ビルドができてればいいんじゃないんですかねぴよ🐥？",
		},
		{
			msg:  "しんどい気持ちで書いてました",
			want: "つらぴよ🐥気持ちで書いてましたぴよぴよ🐥",
		},
		{
			msg:  "しんどい気持ちで書いてました。",
			want: "つらぴよ🐥気持ちで書いてましたぴよぴよ🐥",
		},
		{
			msg:  "しんどい気持ちで書いてました。",
			want: "つらぴよ🐥気持ちで書いてましたぴよぴよ🐥",
		},
		{
			msg:  "つらい気持ちで書いてました。",
			want: "つらぴよ🐥気持ちで書いてましたぴよぴよ🐥",
		},
		{
			msg:  "辛い気持ちで書いてました。",
			want: "つらぴよ🐥気持ちで書いてましたぴよぴよ🐥",
		},
		{
			msg:  "トイレじゃないの",
			want: "トイレじゃないぴよ🐥？",
		},
	}
	for _, tt := range tests {
		t.Run(tt.msg, func(t *testing.T) {
			if got := Do(tt.msg); got != tt.want {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestT(t *testing.T) {
// 	tokens := kagomeTokenizer.Tokenize("良さそう")
// 	for _, token := range tokens {
// 		t.Logf("%s\t%v\n", token.Surface, strings.Join(token.Features(), ","))
// 	}
// 	t.Fail()
// }
