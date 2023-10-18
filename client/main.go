package openapi

import "fmt"

func main() {
	cfg := NewConfiguration()
	apiClient := NewAPIClient(cfg)
	req := ApiPingGetRequest{}
	ret, _, err := apiClient.ExampleAPI.PingGetExecute(req)
	if err != nil {
		fmt.Printf("%v", ret)
	}
}
