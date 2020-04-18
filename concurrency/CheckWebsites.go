package concurrency


type WebsiteChecker func(string) bool
type result struct{
	string
	bool
}

func CheckWebsites(wc WebsiteChecker,
	urls []string) map[string]bool{
	results := make(map[string]bool)
	resultChannel := make(chan result)
	/*
		The net result of using
		the channel is taht we've 
		parallelized the expensive 
		component (calling urls)
		but we've avoided race condition
		issues on populating the 
		non-threadsafe results hashmap

	*/
	for _,url := range urls {
		go func(u string){
			// this is a send statement.
			resultChannel <- result{u,wc(u)}
		}(url)
	}
	for i := 0; i < len(urls); i++{
		// This is a receive expression
		result := <-resultChannel
		results[result.string] = result.bool
	}
	return results
}
