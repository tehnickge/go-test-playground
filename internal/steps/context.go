package steps

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

// MakeContext — демонстрация основных паттернов работы с context.Context
func MakeContext() {
	fmt.Printf("=== Демонстрация context в Go ===\n")

	demoBackgroundAndTODO()

	demoWithCancel()

	demoWithTimeout()

	demoWithDeadline()

	demoWithValue()

	demoErrgroupWithContext()

	fmt.Println("\n=== Конец демонстрации ===")
}

// ────────────────────────────────────────────────
// Базовые контексты (никогда не отменяются сами)
// ────────────────────────────────────────────────
func demoBackgroundAndTODO() {
	fmt.Println("1. Background и TODO")

	bg := context.Background()
	todo := context.TODO()

	fmt.Printf("  Background: %v  (всегда один и тот же объект)\n", bg)
	fmt.Printf("  TODO:       %v  (заглушка, семантически «ещё не определил»)\n", todo)
	fmt.Println("  Оба никогда не отменяются автоматически")

}

// ────────────────────────────────────────────────
// WithCancel — ручная отмена
// ────────────────────────────────────────────────
func demoWithCancel() {
	fmt.Println("\n2. context.WithCancel — ручная отмена")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// имитируем внешнее событие отмены
	go func() {
		time.Sleep(700 * time.Millisecond)
		fmt.Println("  → кто-то вызвал cancel()")
		cancel()
	}()

	waitForDone(ctx, "WithCancel", 2500*time.Millisecond)
}

// ────────────────────────────────────────────────
// WithTimeout — автоматическая отмена по времени
// ────────────────────────────────────────────────
func demoWithTimeout() {
	fmt.Println("\n3. context.WithTimeout")

	ctx, cancel := context.WithTimeout(context.Background(), 1100*time.Millisecond)
	defer cancel()

	waitForDone(ctx, "WithTimeout", 2000*time.Millisecond)
}

// ────────────────────────────────────────────────
// WithDeadline — отмена по конкретному моменту времени
// ────────────────────────────────────────────────
func demoWithDeadline() {
	fmt.Println("\n4. context.WithDeadline")

	deadline := time.Now().Add(1300 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	waitForDone(ctx, "WithDeadline", 2000*time.Millisecond)
}

// ────────────────────────────────────────────────
// WithValue — передача request-scoped значений
// ────────────────────────────────────────────────
func demoWithValue() {
	fmt.Println("\n5. context.WithValue — передача метаданных")

	type (
		userKey      struct{}
		requestIDKey struct{}
		traceIDKey   struct{}
	)

	ctx := context.Background()
	ctx = context.WithValue(ctx, userKey{}, "alice")
	ctx = context.WithValue(ctx, requestIDKey{}, "req-9f3k2m8")
	ctx = context.WithValue(ctx, traceIDKey{}, "trace-abc123")

	fmt.Println("  Пользователь  :", ctx.Value(userKey{}))
	fmt.Println("  Request ID    :", ctx.Value(requestIDKey{}))
	fmt.Println("  Trace ID      :", ctx.Value(traceIDKey{}))
	fmt.Println("  Несуществующий ключ →", ctx.Value("random-key")) // nil
}

// ────────────────────────────────────────────────
// errgroup + WithContext — параллельные задачи с отменой
// ────────────────────────────────────────────────
func demoErrgroupWithContext() {
	fmt.Println("\n6. errgroup.WithContext — параллельные задачи + отмена при ошибке")

	g, ctx := errgroup.WithContext(context.Background())

	for i := 1; i <= 5; i++ {
		i := i
		g.Go(func() error {
			dur := time.Duration(300*i) * time.Millisecond

			fmt.Printf("  Задача %d → старт (спит %v)\n", i, dur)

			select {
			case <-time.After(dur):
				if i == 3 {
					fmt.Printf("  Задача %d → ошибка!\n", i)
					return fmt.Errorf("критическая ошибка в задаче %d", i)
				}
				fmt.Printf("  Задача %d → успех\n", i)
				return nil

			case <-ctx.Done():
				fmt.Printf("  Задача %d → отменена (%v)\n", i, ctx.Err())
				return ctx.Err()
			}
		})
	}

	err := g.Wait()
	if err != nil {
		fmt.Printf("  Итог: ошибка → %v\n", err)
	} else {
		fmt.Println("  Итог: все задачи успешно завершены")
	}
}

// ────────────────────────────────────────────────
// Вспомогательная функция — ждём отмены или таймаут
// ────────────────────────────────────────────────
func waitForDone(ctx context.Context, name string, maxWait time.Duration) {
	start := time.Now()

	select {
	case <-time.After(maxWait):
		fmt.Printf("  %-12s → НЕ ОТМЕНИЛСЯ за %v\n", name, maxWait)
	case <-ctx.Done():
		duration := time.Since(start)
		fmt.Printf("  %-12s → отменён через %4d мс, причина: %v\n",
			name, duration.Milliseconds(), ctx.Err())
	}
}
