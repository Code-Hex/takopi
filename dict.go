package takopi

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
	}: "そうだっピ",
	{
		pos:   "感動詞",
		affix: "*",
		want:  "いいえ",
	}: "ちがうっピ",
	{
		pos:   "記号",
		affix: "一般",
		want:  "？",
	}: "っピ？",
	{
		pos:   "形容詞",
		affix: "自立",
		want:  "しんどい",
	}: "つらいっピ",
	{
		pos:   "形容詞",
		affix: "自立",
		want:  "つらい",
	}: "つらいっピ",
	{
		pos:   "形容詞",
		affix: "自立",
		want:  "辛い",
	}: "つらいっピ",
	{
		pos:   "感動詞",
		affix: "*",
		want:  "ありがとう",
	}: "ありがとだっピ！",
	{
		pos:   "動詞",
		affix: "非自立",
		want:  "てる",
	}: "てるっピ！",
	{
		pos:   "助詞",
		affix: "副助詞",
		want:  "か",
	}: "っピか",
}
