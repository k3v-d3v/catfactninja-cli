package catfacts

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type CatFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

type CatFacts []CatFact

type Breed struct {
	Breed   string `json:"breed"`
	Country string `json:"country"`
	Origin  string `json:"origin"`
	Coat    string `json:"coat"`
	Pattern string `json:"pattern"`
}

type Breeds []Breed

type CatFactApi struct {
	client  *http.Client
	baseUrl string
}

func NewCatFactApi(baseUrl string) *CatFactApi {
	return &CatFactApi{
		client:  &http.Client{},
		baseUrl: baseUrl,
	}
}

func (c *CatFactApi) GetFacts() CatFacts {
	req, err := http.NewRequest("GET", c.baseUrl+"/facts", nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("limit", "332")
	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	catFacts := struct {
		Data CatFacts `json:"data"`
	}{}

	// Unmarshal the JSON payload into the catFacts struct
	if err := json.Unmarshal(body, &catFacts); err != nil {
		log.Fatal(err)
	}

	return catFacts.Data
}

func (c *CatFactApi) GetFact() *CatFact {
	resp, err := c.client.Get(c.baseUrl + "/fact")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	catFact := &CatFact{}

	// Unmarshal the JSON payload into the catFact struct
	if err := json.Unmarshal(body, catFact); err != nil {
		log.Fatal(err)
	}

	return catFact
}

func (c *CatFactApi) GetBreeds() Breeds {
	req, err := http.NewRequest("GET", c.baseUrl+"/breeds", nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("limit", "98")
	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	breeds := struct {
		Data []Breed `json:"data"`
	}{}

	// Unmarshal the JSON payload into the breeds struct
	if err := json.Unmarshal(body, &breeds); err != nil {
		log.Fatal(err)
	}

	return breeds.Data
}
