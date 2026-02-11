package utils

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"time"
)

func PostRequest(url string, data []byte, contenttype string) (string, error) {
	bodyReader := bytes.NewReader(data)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
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

func GetRequest(url string) (string, error) {
	return GetRequestWithUserAgent(url, "")
}

func GetRequestWithUserAgent(url string, userAgent string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	// Note: Content-Type header should not be set on GET requests
	// Some servers (like ipinfo.io) may return HTML instead of JSON if they detect certain User-Agents
	// Using the default Go HTTP client User-Agent works best for most API requests
	r.Header.Set("Accept", "application/json")

	// Only set User-Agent if explicitly provided (required for some APIs like Nominatim)
	if userAgent != "" {
		r.Header.Set("User-Agent", userAgent)
	}

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
