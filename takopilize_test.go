package takopi

import (
	"testing"
)

func TestDo(t *testing.T) {
	tests := []struct {
		msg  string
		want string
	}{
		{
			msg:  "どうだ",
			want: "どうだっピ！",
		},
		{
			msg:  "はい、いいえ、どうだ",
			want: "そうだっピ、ちがうっピ、どうだっピ！",
		},
		{
			msg:  "課題はもう終わった？",
			want: "課題はもう終わったっピ？",
		},
		{
			msg:  "共有はしてたけど",
			want: "共有はしてたけどっピ！",
		},
		{
			msg:  "もう行きます？",
			want: "もう行きますかっピ？",
		},
		{
			msg:  "一緒にご飯どうですか",
			want: "一緒にご飯どうですかっピ？",
		},
		{
			msg:  "一緒にご飯どうでしょうか",
			want: "一緒にご飯どうでしょうかっピ？",
		},
		{
			msg:  "ビルドができてればいいんじゃないんですかね",
			want: "ビルドができてればいいんじゃないんですかねっピ？",
		},
		{
			msg:  "しんどい気持ちで書いてました",
			want: "つらいっピ気持ちで書いてましたっピ！",
		},
		{
			msg:  "しんどい気持ちで書いてました。",
			want: "つらいっピ気持ちで書いてましたっピ！",
		},
		{
			msg:  "しんどい気持ちで書いてました。",
			want: "つらいっピ気持ちで書いてましたっピ！",
		},
		{
			msg:  "つらい気持ちで書いてました。",
			want: "つらいっピ気持ちで書いてましたっピ！",
		},
		{
			msg:  "辛い気持ちで書いてました。",
			want: "つらいっピ気持ちで書いてましたっピ！",
		},
		{
			msg:  "トイレじゃないの",
			want: "トイレじゃないっピ？",
		},
		{
			msg:  "ありがとう",
			want: "ありがとだっピ！",
		},
		{
			msg:  "勉強してる",
			want: "勉強してるっピ！",
		},
		{
			msg:  "これからどうする?",
			want: "これからどうするっピ？",
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
// 	tokens := kagomeTokenizer.Tokenize("一緒にご飯どうでしょうか")
// 	for _, token := range tokens {
// 		t.Logf("%s\t%v\n", token.Surface, strings.Join(token.Features(), ","))
// 	}
// 	t.Fail()
// }
