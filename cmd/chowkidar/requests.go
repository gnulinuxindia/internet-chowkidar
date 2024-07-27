package main

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"time"
)

func postRequest(url string, data []byte, contenttype string) (string, error) {
	bodyReader := bytes.NewReader(data)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bodyReader)
	if err != nil {
		return "", err
	}

	//UserAgent := "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/115.0"
	r.Header.Set("Content-Type", contenttype)
	//r.Header.Set("User-Agent", UserAgent)
	r.Header.Set("Accept", "application/json, text/plain, */*")
	client := &http.Client{}
	res, err0 := client.Do(r)
	if err0 != nil {
		return "", err0
	}

	defer res.Body.Close()

	body, err1 := io.ReadAll(res.Body)
	if err1 != nil {
		return "", err1
	}

	return string(body), nil
}

func getRequest(url string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	//UserAgent := "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/115.0"
	// r.Header.Set("Content-Type", "application/json")
	//r.Header.Set("User-Agent", UserAgent)

	client := &http.Client{}
	res, err0 := client.Do(r)
	if err0 != nil {
		return "", err0
	}

	defer res.Body.Close()

	body, err1 := io.ReadAll(res.Body)
	if err1 != nil {
		return "", err1
	}
	return string(body), nil
}
