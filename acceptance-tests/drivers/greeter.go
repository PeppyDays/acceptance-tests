package drivers

import (
	"fmt"
	"io"
	"net/http"
)

type HTTPGreeter struct {
	Client  *http.Client
	BaseURL string
}

func (g HTTPGreeter) Greet(name string) (string, error) {
	url := fmt.Sprintf("%s/greet", g.BaseURL)
	if name != "" {
		url = fmt.Sprintf("%s?name=%s", url, name)
	}
	res, err := g.Client.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	greeting, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(greeting), nil
}

func (g HTTPGreeter) Curse(name string) (string, error) {
	url := fmt.Sprintf("%s/curse?name=%s", g.BaseURL, name)
	res, err := g.Client.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	cursing, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(cursing), nil
}
