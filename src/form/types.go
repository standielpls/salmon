package form

type SpareFields = int

var FormTypes = map[string]struct{}{
	"Stan":  struct{}{},
	"Spare": struct{}{},
}

type StanFields = int

const (
	StanName StanFields = iota
	StanEmail
	StanSubject
	StanBody
)

// TODO: This makes more sense in a Config file
var StanFieldMapping = map[StanFields]string{
	StanName:    "1378387711",
	StanEmail:   "1870177209",
	StanSubject: "1070703679",
	StanBody:    "1921527425",
}
