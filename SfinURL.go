package simplefin

import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
	"strconv"
	"time"
)

// FIXME: Add progress indicator, perhaps filesize

func SfinURL(url string, config ...URLConfig) (Accounts, error) {
	var accts Accounts

	if url == "" {
		return accts, fmt.Errorf("URL is blank!")
	}

	sfClient := http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, url + "/accounts", nil)
	if err != nil {
		return accts, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("User-Agent", "github.com/jazzboME/simplefin")
	if len(config) > 0 {
		conf := config[0]
		if conf.pending {
			req.Header.Set("pending", "1")
		}
		if conf.balancesOnly {
			req.Header.Set("balances-only", "1")
		}
		if conf.startdate > 1 {
			req.Header.Set("start-date", strconv.FormatInt(conf.startdate, 10))
		}
		if conf.enddate > 1 {
			req.Header.Set("end-date", strconv.FormatInt(conf.enddate, 10))
		}
		for _, str := range conf.accounts {
			req.Header.Set("account", str)
		}
	}

	resp, err := sfClient.Do(req)
	if err != nil {
		return accts, fmt.Errorf("failed to get response: %w", err)
	}
	if resp.Body == nil {
		return accts, fmt.Errorf("response body was nil")
	}

	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return accts, fmt.Errorf("failed to read response body: %w", err)
	}

	err = json.Unmarshal(body, &accts)
	if err!=nil {
		return accts, fmt.Errorf("failed to unmarshall response body: %w", err)
	}

	return accts, nil
}
