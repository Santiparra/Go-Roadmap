type Comment struct {
	Id      string `json:"id"`
	UserId  string `json:"user_id"`
	Comment string `json:"comment"`
}

func createComment(url, apiKey string, commentStruct Comment) (Comment, error) {
    // encode our comment as json
	jsonData, err := json.Marshal(commentStruct)
	if err != nil {
		return Comment{}, err
	}

    // create a new request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return Comment{}, err
	}

    // set request headers
	req.Header.Set("Content-Type", "application/json")
    req.Header.Set("X-API-Key", apiKey)

    // create a new client and make the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return Comment{}, err
	}
	defer res.Body.Close()

    // decode the json data from the response
	// into a new Comment struct
	var comment Comment
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&comment)
	if err != nil {
		return Comment{}, err
	}

	return comment, nil
}