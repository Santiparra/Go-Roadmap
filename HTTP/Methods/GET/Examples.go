//Simple
resp, err := http.Get("https://jsonplaceholder.typicode.com/users")

//Customizable
client := &http.Client{
	Timeout: time.Second * 10,
}

req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/users", nil)
if err != nil {
	log.Fatal(err)
}

resp, err := client.Do(req)