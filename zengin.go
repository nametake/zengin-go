package zengin

type Bank struct {
	Code     string             `json:"code"`
	Name     string             `json:"name"`
	Kana     string             `json:"kana"`
	Hira     string             `json:"hira"`
	Roma     string             `json:"roma"`
	Branches map[string]*Branch `json:"branches"`
}

type Branch struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Kana string `json:"kana"`
	Hira string `json:"hira"`
	Roma string `json:"roma"`
}
