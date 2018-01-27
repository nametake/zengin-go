// Package internal provides ...
package internal

import zengin "github.com/nametake/zengin-go"

var Banks = map[string]*zengin.Bank{
	"0001": &zengin.Bank{
		Code: "0001",
		Name: "みずほ",
		Kana: "ミズホ",
		Hira: "みずほ",
		Roma: "mizuho",
		Branches: map[string]*zengin.Branch{
			"001": &zengin.Branch{
				Code: "001",
				Name: "東京営業部",
				Kana: "トウキヨウ",
				Hira: "とうきよう",
				Roma: "toukiyou",
			},
			"004": &zengin.Branch{
				Code: "004",
				Name: "丸の内中央",
				Kana: "マルノウチチユウオウ",
				Hira: "まるのうちちゆうおう",
				Roma: "marunouchichiyuuou",
			},
			"005": &zengin.Branch{
				Code: "005",
				Name: "丸之内",
				Kana: "マルノウチ",
				Hira: "まるのうち",
				Roma: "marunouchi",
			},
		},
	},
	"0005": &zengin.Bank{
		Code: "0005",
		Name: "三菱東京ＵＦＪ",
		Kana: "ミツビシトウキヨウＵＦＪ",
		Hira: "みつびしとうきようＵＦＪ",
		Roma: "mitsubishitoukiyouufj",
		Branches: map[string]*zengin.Branch{
			"001": &zengin.Branch{
				Code: "001",
				Name: "本店",
				Kana: "ホンテン",
				Hira: "ほんてん",
				Roma: "honten",
			},
			"002": &zengin.Branch{
				Code: "002",
				Name: "丸の内",
				Kana: "マルノウチ",
				Hira: "まるのうち",
				Roma: "marunouchi",
			},
			"003": &zengin.Branch{
				Code: "003",
				Name: "瓦町",
				Kana: "カワラマチ",
				Hira: "かわらまち",
				Roma: "kawaramachi",
			},
		},
	},
}
