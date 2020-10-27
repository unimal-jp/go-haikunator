package haikunator

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

var adjectives []string = []string{
	"aged", "ancient", "autumn", "billowing", "bitter", "black", "blue", "bold", "broad", "broken", "calm", "cold", "cool", "crimson", "curly", "damp", "dark", "dawn", "delicate", "divine", "dry", "empty", "falling", "fancy", "flat", "floral", "fragrant", "frosty", "gentle", "green", "hidden", "holy", "icy", "jolly", "late", "lingering", "little", "lively", "long", "lucky", "misty", "morning", "muddy", "mute", "nameless", "noisy", "odd", "old", "orange", "patient", "plain", "polished", "proud", "purple", "quiet", "rapid", "raspy", "red", "restless", "rough", "round", "royal", "shiny", "shrill", "shy", "silent", "small", "snowy", "soft", "solitary", "sparkling", "spring", "square", "steep", "still", "summer", "super", "sweet", "throbbing", "tight", "tiny", "twilight", "wandering", "weathered", "white", "wild", "winter", "wispy", "withered", "yellow", "young",
}

var nouns []string = []string{
	"reizoko", "reitoko", "sentakuki", "oishii", "sashimi", "sushi", "onigiri", "katana", "anime", "maguro", "fuji", "musashi", "samurai", "ninja", "yoyo", "sakura", "keitai", "denwa", "jyudenki", "hikoki", "densha", "chikatetsu", "basu", "takushi", "jitensha", "eki", "basutei", "kuko", "shingo", "kosaten", "michi", "doro", "shogakko", "shogakko", "koko", "daigaku", "seito", "sensei", "senko", "gakui", "shiken", "shukudai", "ishi", "kangoshi", "bengoshi", "ryorinin", "jitsugyoka", "keisatsukan", "shoboshi", "gijutsusha", "komuin", "kaikeishi", "oya", "musume", "musuko", "sobo", "sofu", "karada", "atama", "kata", "ude", "ashi", "mune", "hara", "kao", "me", "hana", "kuchi", "mimi", "chita", "inu", "neko", "ushi", "buta", "uma", "hitsuji", "saru", "nezumi", "tora", "okami", "usagi", "ryu", "shika", "kaeru", "shishi", "kirin", "zo", "tori", "niwatori", "suzume", "karasu", "washi", "taka", "sakana", "tai", "ebi", "iwashi", "maguro", "katsuo", "sanma", "aji", "saba", "ika", "tako", "mushi", "cho", "semi", "tonbo", "kumo", "hotaru", "ka", "kai", "kaigara", "kuma", "ongaku", "rika", "sansu", "rekishi", "chiri", "taiiku", "supotsu", "shisutemu", "joho", "hitsuyo", "benkyo", "aka", "midori", "ao", "murasaki", "shiro", "kuro", "pinku", "orenji", "yu", "kori", "hi", "gasu", "kuki", "tsuchi", "kinzoku", "doro", "kemuri", "tetsu", "do", "kin", "gin", "namari", "shio", "kita", "higashi", "minami", "nishi", "taiyo", "tsuki", "hoshi", "tenki", "hare", "ame", "kumori", "yuki", "kaze", "kaminari", "taifu", "arashi", "sora", "nichiyobi", "getsuyobi", "kayobi", "suiyobi", "mokuyobi", "kinyobi", "doyobi", "nomimono", "cha", "ocha", "kohi", "gyunyu", "mizu", "biru", "wain", "sato", "shio", "shoyu", "kome", "ine", "mugi", "yasai", "kudamono", "imo", "mame", "daikon", "ninjin", "ringo", "mikan", "nashi", "kuri", "momo", "suika", "bunbogu", "inku", "pen", "ballborupen", "mannenhitsu", "enpitsu", "fude", "choku", "keshigomu", "enpitsukezuri", "jogi", "noto", "nikki", "futo", "hasami", "hotchikisu", "david",
}

var hexTokenChars = "0123456789abcdef"
var numTokenChars = "0123456789"
var alphaTokenChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func gen(tokenLength int, tokenChars string) string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	adj := adjectives[r1.Intn(len(adjectives))]
	nou := nouns[r1.Intn(len(nouns))]

	var tok bytes.Buffer
	for i := 0; i < tokenLength; i++ {
		x := string(tokenChars[r1.Intn(len(tokenChars))])
		tok.WriteString(x)
	}

	return fmt.Sprintf("%s-%s-%s", adj, nou, tok.String())
}

// Hex : Create hex string
func Hex(tokenLength int) string {
	return gen(tokenLength, hexTokenChars)
}

// Num : Create num string
func Num(tokenLength int) string {
	return gen(tokenLength, numTokenChars)
}

// Alpha : Create alpha string
func Alpha(tokenLength int) string {
	return gen(tokenLength, alphaTokenChars)
}

// Alnum : Create alpha string
func Alnum(tokenLength int) string {
	return gen(tokenLength, alphaTokenChars+numTokenChars)
}
