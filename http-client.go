package go_sdk_cargo_sdek

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
	"time"
)

func (c *Client) get(method string, res, resErr interface{}) (int, error) {

	code, body, err := c.http(http.MethodGet, method, nil)
	if err != nil {
		panic(err)
	}
	return c.res(code, body, res, resErr)
}

func (c *Client) post(method string, req, res, resErr interface{}) (int, error) {

	code, body, err := c.http(http.MethodPost, method, req)
	if err != nil {
		panic(err)
	}
	return c.res(code, body, res, resErr)
}

func (c *Client) res(code int, body []byte, res, resErr interface{}) (int, error) {

	var err error
	if code == 200 || code == 202 {
		// Ok
		err = json.Unmarshal(body, res)
		if err != nil {
			panic(err)
		}
		return code, nil
	} else {
		// ERR
		err = json.Unmarshal(body, resErr)
		if err != nil {
			panic(err)
		}
		log.Println("[E] client.POST", string(body))
		return code, nil
	}
}

func (c *Client) http(httpMethod string, method string, reqJson interface{}) (int, []byte, error) {

	var bodyReq []byte
	if reqJson != nil {
		bodyReq, _ = json.MarshalIndent(reqJson, "", "    ")
	}

	u := c.endPoint + method
	req, err := http.NewRequest(httpMethod, u, bytes.NewReader(bodyReq))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json;charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+c.token)

	if c.DEBUG {
		reqDump, _ := httputil.DumpRequest(req, true)
		log.Printf("[D] REQ:\n%s\n\n", string(reqDump))
		fn, err := os.Create(LOG_DIR + time.Now().Format("2006-01-02_15-04-05.000000000"+"_"+httpMethod+".req"))
		if err != nil {
			panic(err)
		}
		defer func() { _ = fn.Close() }()
		_, _ = fn.Write(reqDump)
		_ = fn.Close()
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer func() { _ = res.Body.Close() }()

	if c.DEBUG {
		resDump, _ := httputil.DumpResponse(res, true)
		log.Printf("[D] RES:\n%s\n\n", string(resDump))
		fName := time.Now().Format("2006-01-02_15-04-05.000000000") + "_" + httpMethod + "_" + strconv.Itoa(res.StatusCode) + ".res"
		fn, err := os.Create(LOG_DIR + fName)
		if err != nil {
			panic(err)
		}
		defer func() { _ = fn.Close() }()
		_, _ = fn.Write(resDump)
		_ = fn.Close()
	}

	body, _ := io.ReadAll(res.Body)

	return res.StatusCode, body, err
}
