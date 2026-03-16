package constants

type Inner struct {
	Saf string
	Dsa int
}

var ArrayAny = []interface{}{
	1,
	"",
	Inner{Saf: "val", Dsa: 100},
	4,
	nil,
	"test",
}

var testMap = make(map[string]int)
var TestMap = map[string]int{
	"one":   1,
	"two":   2,
	"ten":   10,
	"token": 2000,
}
