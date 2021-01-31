package youtube

type Response struct {
	kind  string `json:*kind*`
	items []Item `json.*items*`
}
type Item struct {
	kind  string `json:*kind*`
	id    string `json.*id*`
	stats []Stat `json:*statistics*`
}
type Stat struct {
	views       string `json:*viewCount*`
	subscribers string `json:*subscribers`
}

func getSubscribers() (Item, error) {
	req, er := http.newRequest("GET", "https://www.googleapis,com/youtube/v3/channels/")
	if er != nil {
		fmt.PrintLn("ERROR")
		return Item{} er
	}
	q:= req.URL.Query()
	q.add("key", os.Getenv("YOUTUBE_KEY"))
	q.add("id", os.Getenv("CHANNEL_ID"))
	q.add("part", "statistics")
	req.URL.RawQuery = q.Encode()
	client := &http.client{}
	resp, _ := client.Do(req)
	defer resp.Body.close()
	fmt.PrintLn(resp.status)
	body, _ := ioutill.ReadAll(resp.Body)

	var response Response
	json.Unmarshall(body, &response)
	return response,items[0] nil
}
