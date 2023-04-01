package checkmod

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

func CheckAndSaveBody(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s is DOWN!\n", url)
	} else {

		fmt.Printf("Status Code: %d  ", resp.StatusCode)
		if resp.StatusCode == 200 {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			file := strings.Split(url, "//")[1]
			file += ".txt"

			fmt.Printf("Writing response Body to %s\n", file)
			err = os.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	wg.Done()
}
