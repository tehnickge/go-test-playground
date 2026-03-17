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

// Инициализация map с помощью литерала
// make - встроенная функция для создания слайсов, map и каналов.
// Она выделяет память и возвращает указатель на новый объект.
var testMap = make(map[string]int)
var testMap2 = new(map[string]int)
var TestMap = map[string]int{
	"one":   1,
	"two":   2,
	"ten":   10,
	"token": 2000,
}

// Константы для статусов пользователей
// iota - это специальная константа, которая используется для создания
// последовательных числовых констант. Она автоматически увеличивается на единицу
// при каждой новой константе в блоке.
const (
	ONLINE = iota
	OFFLINE
	AWAY
	REJECTED
)
