package steps

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func MakeReadAndWrite() {
	basicReadAndWrite()
	fmt.Println("-----------")
	bufioReadAndWrite()
	fmt.Println("-----------")
	simpleScaner()
}

func basicReadAndWrite() {
	reader := strings.NewReader("test 🎉 msg 👪")
	var newString strings.Builder
	buffer := make([]byte, 4)

	for {
		numBytes, err := reader.Read(buffer)
		chunk := buffer[:numBytes]
		newString.Write(chunk)
		fmt.Printf("Read %v bytes: %c\n", numBytes, chunk)
		if err == io.EOF {
			break
		}
	}

	fmt.Printf("%v\n", newString.String())
}

func bufioReadAndWrite() {
	source := strings.NewReader("simple text for example")
	buffered := bufio.NewReader(source)
	newString, err := buffered.ReadString('\n')
	if err == io.EOF {
		fmt.Println(newString)
	} else {
		fmt.Println("went wrong")
	}
}

func simpleScaner() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0, 5)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		if len(lines) == 5 {
			break
		}
	}
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}
	fmt.Printf("line count: %v\n", len(lines))
	for i, value := range lines {
		fmt.Printf("line:%d - %v\n", i, value)
	}

}
