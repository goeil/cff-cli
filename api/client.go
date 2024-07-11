package api

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
	//"github.com/spf13/cobra"
)

type ClientType int64

const (
	LoginRequest ClientType = iota
	AppRequest
	CsvBuild
)

type Client struct {
	apiAddr string
	client  http.Client
}

func GetClient() *Client {

	clientKey := "clientPointer"

	c := viper.Get(clientKey)
	if c == nil {
		//fmt.Println("Instantiate new client pointer.")
		//rootUrl := viper.GetString("apiUrl") no reading from config file config.yml (sinon ne marche pas de n'importe où)
		rootUrl := "http://transport.opendata.ch/v1"
		//fmt.Println(rootUrl)

		c := InitClient(rootUrl)
		viper.Set(clientKey, c)
		return c
	} else {
		//fmt.Println("Found client pointer in memory.")
		return c.(*Client)
	}

}
func InitClient(apiAddr string) *Client {

	//fmt.Println("Initialisation du client.")
	return &Client{
		apiAddr: apiAddr,
		client:  *http.DefaultClient,
	}
}

type RequestParameters struct {
	Start string
	End   string
	Time  string
	Date  string
}

/**
 * Get method
 */
func (c *Client) Get(params RequestParameters) ([]byte, error) {

	var (
		request  *http.Request
		response *http.Response
		err      error
	)
	//jww.TRACE.Println("GET " + url)

	nbResults := 5

	url := fmt.Sprintf("%v/connections?limit=%v&from=%v&to=%v&time=%v&date=%v",
		c.apiAddr,
		nbResults,
		params.Start,
		params.End,
		params.Time,
		params.Date,
	)
	request, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("request building: %w", err)
	}
	// headers not necessary for GET ?
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err = c.client.Do(request)

	//fmt.Println(url)
	body, _ := io.ReadAll(response.Body)
	defer response.Body.Close()

	switch response.StatusCode {
	case 401:
		// unauthenticated

	case 200:
		return body, nil

	default:
		return nil, fmt.Errorf("request building: %w", errors.New("authentication error"))
	}

	if err != nil {
		return nil, fmt.Errorf("request building: %w", errors.New("API error"))
	}
	jww.TRACE.Println("Requête API exécutée avec succès.")
	return nil, nil

}
