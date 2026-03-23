package steps

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

var (
	counter  int
	counter2 int
)

func MakeMutex() {
	var g errgroup.Group
	var mu sync.Mutex      // для защиты counter и counter2
	var printMu sync.Mutex // для красивого вывода в консоль

	// Первый набор задач — просто увеличивают counter
	for i := 0; i < 10; i++ {
		i := i
		g.Go(func() error {
			time.Sleep(time.Duration(i%30) * 2 * time.Millisecond)
			increment(&mu)
			return nil
		})
	}

	// Второй набор задач — могут вернуть ошибку
	for i := 0; i < 10; i++ {
		i := i
		g.Go(func() error {
			time.Sleep(time.Duration(i%40+10) * 3 * time.Millisecond)

			// безопасный вывод номера задачи
			printMu.Lock()
			fmt.Printf("→ задача %4d начала\n", i)
			printMu.Unlock()

			if i == 8 {
				printMu.Lock()
				fmt.Printf("!!! Ошибка в задаче %d !!!\n", i)
				printMu.Unlock()
				return fmt.Errorf("ошибка в задаче N: %d", i)
			}

			// увеличиваем только если ошибки не было
			incrementWithError(&mu)

			printMu.Lock()
			fmt.Printf("  задача %4d → успех (counter2 = %d)\n", i, counter2)
			printMu.Unlock()

			return nil
		})
	}

	// ждём завершения всех горутин
	err := g.Wait()

	if err != nil {
		fmt.Printf("\nПроизошла ошибка: %v\n", err)
		fmt.Printf("counter   = %d\n", counter)
		fmt.Printf("counter2  = %d  ← не все задачи дошли сюда\n", counter2)
	} else {
		fmt.Println("\nВыполнено удачно (без ошибок)")
		fmt.Printf("counter   = %d\n", counter)
		fmt.Printf("counter2  = %d\n", counter2)
	}
}

func increment(mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	counter++
}

func incrementWithError(mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	counter2++
}
