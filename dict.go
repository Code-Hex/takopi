package piyolize

type key struct {
	pos   string // 名詞や感動詞など
	affix string // 接辞 ~ 接尾 など
	want  string // 期待するワード
}

var dict = map[key]string{
	{
		pos:   "感動詞",
		affix: "*",
		want:  "はい",
	}: "そうぴよ🐥",
	{
		pos:   "感動詞",
		affix: "*",
		want:  "いいえ",
	}: "ちがうぴよ🐥",
	{
		pos:   "記号",
		affix: "一般",
		want:  "？",
	}: "ぴよ🐥？",
	{
		pos:   "名詞",
		affix: "一般",
		want:  "人",
	}: "🐥",
	{
		pos:   "形容詞",
		affix: "自立",
		want:  "しんどい",
	}: "つらぴよ🐥",
	{
		pos:   "形容詞",
		affix: "自立",
		want:  "つらい",
	}: "つらぴよ🐥",
	{
		pos:   "形容詞",
		affix: "自立",
		want:  "辛い",
	}: "つらぴよ🐥",
}
