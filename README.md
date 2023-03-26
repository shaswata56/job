# job
A Go package to schedule both recurring and one time job.

## Example Usage
```go

package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/shaswata56/job"
)

func GetWeatherNow(args ...any) {
	city := ""
	if str, ok := args[0].(string); ok {
		city = str
	}
	url := "https://wttr.in/" + strings.ToLower(city) + "?format=2"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Print("Dhaka: " + string(body))
}

func main() {
	myArgs := []any{}
	myArgs = append(myArgs, "Dhaka")
	myJob := job.Job{
		Fn:   GetWeatherNow,
		Args: myArgs,
	}
	myJob.ScheduleRecurring(time.Second * 2)
	time.Sleep(time.Minute)
}
```
## Import in your project by go-get
```bash
go get github.com/shaswata56/job
```
