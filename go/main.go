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
	Img    string `json:"img"`
}

func main(){

	sdkID := os.Getenv("SDK_ID")
	baseURL := os.Getenv("BASE_URL")
	endpoint := os.Getenv("ENDPOINT")
	keyFile := os.Getenv("PEM_FILE_PATH")
	imgPath := os.Getenv("TEST_IMAGE_PATH")

	file, err := os.Open(imgPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	encoded := base64.StdEncoding.EncodeToString(content)

	estimation := &Estimation{
		Img: encoded,
	}

	jsonData, err := json.Marshal(estimation)
	if err != nil {
		fmt.Println(err)
		return
	}

	key, err := ioutil.ReadFile(keyFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Create request
	req, err := requests.SignedRequest{
		HTTPMethod: http.MethodPost,
		BaseURL:    baseURL,
		Endpoint:   "/" + endpoint,
		Headers: map[string][]string{
			"Content-Type": {"application/json"},
			"Accept":       {"application/json"},
			"X-Yoti-Auth-Id":{sdkID},
		},
		Body: jsonData,
	}.WithPemFile(key).Request()
	if err != nil {
		fmt.Println(err)
		return
	}

	//get Yoti response
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	buffer := new(strings.Builder)
	_, err = io.Copy(buffer, response.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(buffer.String())
}
