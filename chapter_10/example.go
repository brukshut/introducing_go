package main

import ("fmt"; "io/ioutil"; "net/http")

type HomePageSize struct {
	URL string
	Size int
}

func main() {
	urls := []string{
		"https://www.swellinfo.com",
		"https://magicseaweed.com",
		"https://www.ndbc.noaa.gov/station_page.php?station=44025",
		"https://thesurfersview.com/live-cams/new-york",
	}
	results := make(chan HomePageSize)

	// generates a function for each url and invokes it
	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()

			bs, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			results <- HomePageSize {
				URL: url,
				Size: len(bs),
			}
		}(url)
	}

	var biggest HomePageSize

	for range urls {
		result := <- results
		if result.Size > biggest.Size {
			biggest = result
		}
	}

	fmt.Println("The biggest home page:", biggest.URL)
}
