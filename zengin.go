package zengin

type Bank struct {
	Code     string
	Name     string
	Kana     string
	Hira     string
	Roma     string
	Branches map[string]*Branch
}

type Branch struct {
	Code string
	Name string
	Kana string
	Hira string
	Roma string
}
