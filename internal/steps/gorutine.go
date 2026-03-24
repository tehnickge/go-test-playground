package steps

import (
	"fmt"
	"sync"
	"time"
)

func MakeGoRutine() {
	arr := []int{1, 2, 3, 4, 5}
	ch := make(chan int, len(arr))
	ch2 := make(chan int, len(arr))
	var wg sync.WaitGroup
	var wg2 sync.WaitGroup

	fmt.Println("start main:")

	for _, v := range arr {
		wg.Add(1)
		wg2.Add(1)
		go func(val int) {
			res := fakeFunc(val, &wg)

			res2 := fakeFuncBasic(val)
			defer wg2.Done()

			ch <- res
			ch2 <- res2
		}(v)
	}

	// Запускаем горутину, которая закроет канал после завершения всех задач
	go func() {
		wg.Wait()
		wg2.Wait()
		close(ch)
		close(ch2)
		fmt.Println("→ каналы закрыты, все задачи выполнены")
	}()

	// Дополнительно ожидаем
	wg.Wait()
	wg2.Wait()

	fmt.Println("end main:")

	// чтение из канала
	wg.Add(1)
	go func() {
		wg.Done()
		for result := range ch {
			fmt.Printf("Получен результат из ch: %d\n", result)
		}

		for result := range ch2 {
			fmt.Printf("Получен результат из ch2: %d\n", result)
		}
	}()

	wg.Wait()
	fmt.Println("Все горутины завершились, main тоже заканчивает работу")

}

func fakeFunc(iter int, wg *sync.WaitGroup) int {
	defer wg.Done()
	fmt.Printf("fakeFunc %d → начало\n", iter)
	time.Sleep(time.Second * 2)
	fmt.Printf("fakeFunc %d → конец\n", iter)
	return iter * 10
}
func fakeFuncBasic(iter int) int {
	fmt.Printf("fakeFuncBasic %d → начало\n", iter)
	time.Sleep(time.Second * 2)
	fmt.Printf("fakeFuncBasic %d → конец\n", iter)
	return iter * 4
}
