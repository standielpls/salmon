package form

type SpareFields = int

var FormTypes = map[string]struct{}{
	"Stan":  struct{}{},
	"Spare": struct{}{},
}

const (
	SpareName SpareFields = iota
	SpareYear
	SpareMonth
	SpareDay
	SpareDescription
	SpareHours
)

var SpareFieldMapping = map[SpareFields]string{
	SpareName:        "1455114666",
	SpareYear:        "1471338760_year",
	SpareMonth:       "1471338760_month",
	SpareDay:         "1471338760_day",
	SpareDescription: "1004500012",
	SpareHours:       "2102112842",
}

type StanFields = int

const (
	StanName StanFields = iota
	StanEmail
	StanSubject
	StanBody
)

var StanFieldMapping = map[StanFields]string{
	StanName:    "1378387711",
	StanEmail:   "1870177209",
	StanSubject: "1070703679",
	StanBody:    "1921527425",
}
