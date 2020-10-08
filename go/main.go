package main

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/getyoti/yoti-go-sdk/v3/requests"
	_ "github.com/joho/godotenv/autoload"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Estimation struct {
	Data    string `json:"data"`
}

func main(){

	sdkID := os.Getenv("SDK_ID")
	baseURL := os.Getenv("BASE_URL")
	keyFile := os.Getenv("PEM_FILE_PATH")
	imgPath := os.Getenv("TEST_IMAGE_PATH")

	file, _ := os.Open(imgPath)
	reader := bufio.NewReader(file)
	content, _ := ioutil.ReadAll(reader)
	encoded := base64.StdEncoding.EncodeToString(content)

	estimation := &Estimation{
		Data:encoded,
	}

	jsonData,err := json.Marshal(estimation)
	if err != nil {
		fmt.Println(err)
	}

	key, _ := ioutil.ReadFile(keyFile)
	// Create request
	req,_ := requests.SignedRequest{
		HTTPMethod: http.MethodPost,
		BaseURL:    baseURL + "/api/v1/age-verification",
		Endpoint:   "/checks",
		Headers: map[string][]string{
			"Content-Type": {"application/json"},
			"Accept":       {"application/json"},
			"X-Yoti-Auth-Id":{sdkID},
		},
		Body: jsonData,
	}.WithPemFile(key).Request()

	//get Yoti response
	response, _ := http.DefaultClient.Do(req)

	buffer := new(strings.Builder)
	_, err = io.Copy(buffer, response.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(buffer.String())
}
