# Go Test Playground

Практический Go-проект, который показывает инженерный уровень в формате "код, который можно запустить, проверить тестами и масштабировать".  
Это не набор разрозненных примеров, а структурированная песочница: от базовых концепций языка до конкурентности, generics, контекстов, работы с БД и unit-тестов.

## Почему этот проект сильный

- Покрывает ключевые зоны Go-разработки: `types`, `interfaces`, `errors`, `goroutines`, `mutex`, `context`, `generics`.
- Показывает умение строить модульную структуру (`internal`, `lib`, `types`, `constants`) и отделять доменную модель от учебных сценариев.
- Включает рабочую интеграцию с PostgreSQL через `gorm` и конфигурацию через `.env`.
- Поддерживает проверяемость через unit-тесты в нескольких модулях.

## Стек

- `Go 1.26.1`
- `gorm` + `postgres driver`
- `godotenv`
- `x/sync/errgroup`
- Unit tests (`go test`)

## Быстрый старт

### 1) Клонирование

```bash
git clone https://github.com/tehnickge/go-test-playground.git
cd go-test-playground
```

### 2) Запуск приложения

```bash
go run main.go
```

В консольном меню выбираются шаги-демонстрации (каждый шаг соответствует отдельному модулю в `internal/steps`).

### 3) Запуск тестов

```bash
go test ./...
```

## Что демонстрирует проект

- **Работа с типами и структурами:** методы, копирование, клонирование, указатели, сериализация в JSON.
- **Контракты и полиморфизм:** интерфейсы, pointer/value receiver, изменение состояния через интерфейс.
- **Надежная обработка ошибок:** пользовательские error-типы, обертывание ошибок, тесты на edge cases.
- **Конкурентность:** goroutines, `sync.WaitGroup`, mutex, безопасный доступ к shared state.
- **Контекст выполнения:** cancel/timeout/deadline/value, coordinated shutdown через context.
- **Generics:** обобщенные функции и контейнеры, ограничения типов, сортировка и max по ordered-типам.
- **Практика ввода/вывода:** stdin/scan/bufio сценарии для CLI.
- **DB flow:** базовые CRUD-операции пользователя через GORM.

## Полная карта проекта с описаниями файлов

> Ниже каждый файл из дерева проекта с назначением.

- `main.go` — точка входа CLI; интерактивное меню и маршрутизация по шагам.
- `go.mod` — модуль `firstapp`, версия Go и зависимости проекта.
- `go.sum` — контрольные суммы зависимостей.
- `docker-compose.yaml` — подготовка локальной инфраструктуры для БД.
- `.gitignore` — исключения для Git.
- `README.md` — документация и витрина проекта.
- `filte-tree.md` — зафиксированное дерево файлов проекта.

- `constants/index.go` — константы и демонстрационные коллекции (`iota`, map, mixed array).

- `internal/config/config.go` — загрузка конфигурации из окружения, структура `AppConfig/DBConfig`.
- `internal/config/db.go` — инициализация глобального подключения `gorm.DB`.

- `internal/models/user.go` — GORM-модель пользователя и имя таблицы.

- `internal/steps/arrayAndSlices.go` — демонстрации массивов и слайсов.
- `internal/steps/context.go` — практические сценарии `context`: cancel, timeout, deadline, value, errgroup.
- `internal/steps/db.go` — учебный CRUD по пользователям через слой `config/models`.
- `internal/steps/errors.go` — пользовательская ошибка деления и сценарии обработки.
- `internal/steps/errors_test.go` — тесты корректности обработки ошибок деления.
- `internal/steps/functions.go` — функции, multiple return values, генерация ранга с ошибкой.
- `internal/steps/generics.go` — generics: ограничения, универсальные функции, generic-структура массива.
- `internal/steps/gorutine.go` — конкурентные и последовательные вычисления для сравнения подходов.
- `internal/steps/interfaces.go` — интерфейсы `Item/Player`, реализация на `Dog/Cat`, pointer semantics.
- `internal/steps/interfaces_test.go` — тесты поведения для интерфейсных реализаций.
- `internal/steps/maps.go` — базовые операции с map (поиск, удаление, итерация).
- `internal/steps/mutex.go` — синхронизация shared state через `sync.Mutex`, включая антипример.
- `internal/steps/ptrs.go` — основы работы с указателями и изменением значения по ссылке.
- `internal/steps/readAndWrite.go` — ввод/вывод через `fmt`, `bufio`, scanner.
- `internal/steps/strings.go` — работа со строками и рунами (Unicode-safe обход).
- `internal/steps/types.go` — демонстрация пользовательского типа `User`, JSON и методы.
- `internal/steps/validations.go` — email-валидация через regexp.
- `internal/steps/validations_test.go` — тесты email-валидации.

- `lib/helpers.go` — утилиты: pointer-helper для string и безопасный `string -> int`.
- `lib/helpers_test.go` — unit-тесты helper-функций.

- `types/user.go` — прикладной тип `User` с методами (`Greet`, `IsAdmin`, `Clone`, изменение имени).

- `public/` — директория под публичные/статические ресурсы (точка расширения проекта).

## Профессиональная ценность (для CV и интервью)

- Демонстрирует не только "синтаксис Go", а системное мышление: архитектура, чистая декомпозиция, тестируемость.
- Показывает понимание production-паттернов: конфигурация, контекст, конкурентность, работа с БД.
- Легко масштабируется: можно добавлять новые `steps` как изолированные модули без ломки существующей структуры.

## Контрибьютинг

Pull requests приветствуются. Лучшая стратегия:

1. Добавить новый шаг в `internal/steps`.
2. Подключить его в `main.go`.
3. Добавить unit-тесты.
4. Обновить `README.md`.

## Контакты

- GitHub: [tehnickge](https://github.com/tehnickge)
- LinkedIn: [Nikita Sobolev](https://www.linkedin.com/in/nikita-sobolev-261bb13b7/)
