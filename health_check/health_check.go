package health_check

import (
	"fmt"
	"net/http"
	"sync"
)

type UrlStatus struct {
	url    string
	status bool
}

type HealthCheck struct {
	urls    []string
	execute func()
}

func NewHealthCheck(urls []string) *HealthCheck {
	return &HealthCheck{urls: urls}
}

func (uc *HealthCheck) checkUrl(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "is down!")
		return
	}
	fmt.Println(url, "is up and running.")
}

func (uc *HealthCheck) checkUrlWithChannel(url string, c chan UrlStatus) {
	_, err := http.Get(url)
	if err != nil {
		c <- UrlStatus{url: url, status: false}
	} else {
		c <- UrlStatus{url: url, status: true}
	}
}

func (uc *HealthCheck) WithSequential() *HealthCheck {
	uc.execute = func() {
		for _, u := range uc.urls {
			uc.checkUrl(u)
		}
	}
	return uc
}

func (uc *HealthCheck) WithWaitGroup() *HealthCheck {
	uc.execute = func() {
		var wg sync.WaitGroup

		for _, u := range uc.urls {
			wg.Add(1)
			go func(url string) {
				defer wg.Done()
				uc.checkUrl(url)
			}(u)
		}

		wg.Wait()
	}
	return uc
}

func (uc *HealthCheck) WithChannel() *HealthCheck {
	uc.execute = func() {
		c := make(chan UrlStatus)
		for _, u := range uc.urls {
			go uc.checkUrlWithChannel(u, c)
		}

		result := make([]UrlStatus, len(uc.urls))
		for i := range result {
			result[i] = <-c
			if result[i].status {
				fmt.Println(result[i].url, "is up and running.")
			} else {
				fmt.Println(result[i].url + " is down!")
			}
		}
	}
	return uc
}

func (uc *HealthCheck) Execute() {
	uc.execute()
}

func main() {
	urls := []string{
		"https://www.easyjet.com/",
		"https://www.skyscanner.de/",
		"https://www.ryanair.com",
		"https://wizzair.com/",
		"https://www.swiss.com/",
	}

	hc := NewHealthCheck(urls)

	hc.WithChannel().Execute()
	// hc.WithWaitGroup().Execute()
}
