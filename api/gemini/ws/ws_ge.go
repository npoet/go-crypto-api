package ws

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
	"time"
)

var key_ge = "65CqlavUlB0YVa5KFsAy"
var secret_ge = "3Ejq4PQ4sjmjw7qrAMxkDy7GnGpX"

type GeminiWSResp struct {
	Type              string `json:"type"`
	OrderID           string `json:"order_id"`
	EventID           string `json:"event_id"`
	APISession        string `json:"api_session"`
	ClientOrderID     string `json:"client_order_ID"`
	Symbol            string `json:"symbol"`
	Side              string `json:"side"`
	Behavior          string `json:"behavior"`
	OrderType         string `json:"order_type"`
	Timestamp         string `json:"timestamp"`
	IsLive            string `json:"is_live"`
	IsCancelled       string `json:"is_cancelled"`
	IsHidden          string `json:"is_hidden"`
	AvgExecutionPrice string `json:"avg_execution_price"`
	ExecutedAmount    string `json:"executed_amount"`
	RemainingAmount   string `json:"remaining_amount"`
	OriginalAmount    string `json:"original_amount"`
	Price             string `json:"price"`
	TotalSpend        string `json:"total_spend"`
}

func generate_nonce() string {
	return string(time.Nanosecond)
}

//func auth_ge(payload []byte) (http.Header) {
//	b64 := base64.StdEncoding.EncodeToString(payload)
//	sig := hmac.New(sha512.New384, []byte(secret_ge))
//	sig.Write([]byte(b64))
//	headers := make(http.Header)
//	headers.Add("Content-Length", "0")
//	headers.Add("Content-Type", "text/plain")
//	headers.Add("X-GEMINI-APIKEY", key_ge)
//	headers.Add("X-GEMINI-PAYLOAD", b64)
//	headers.Add("X-GEMINI-SIGNATURE", hex.EncodeToString(sig.Sum([]byte{})))
//	/*headers := []byte(`
//	{"X-GEMINI-APIKEY":` + string(key_ge) + `,
//	 "X-GEMINI-PAYLOAD":` + string(b64) + `,
//	 "X-GEMINI-SIGNATURE:` + string(sig) + `}`)*/
//	return headers
//}
//
//func OrderSocket() (*websocket.Conn, *http.Response){
//	payload := []byte(`{
//						"request":"/v1/order/events",
//						"nonce":` + generate_nonce() +
//						`}`)
//	dialer := new(websocket.Dialer)
//	conn, resp, err := dialer.Dial("wss://api.gemini.com/v1/order/events", auth_ge(payload))
//	//j := json.NewDecoder(resp.Body)
//	//s, _ := ioutil.ReadAll(j.Buffered())
//	if err != nil {
//		fmt.Println(resp.StatusCode)
//		panic(err)
//	}
//	return conn, resp
//}

func auth_ge(payload map[string]interface{}) http.Header {
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	payloadBase64 := base64.StdEncoding.EncodeToString(payloadJSON)
	h := hmac.New(sha512.New384, []byte(secret_ge))
	h.Write([]byte(payloadBase64))
	signature := hex.EncodeToString(h.Sum(nil))
	headers := make(http.Header)
	headers.Add("Content-Length", "0")
	headers.Add("Content-Type", "text/plain")
	headers.Add("X-GEMINI-APIKEY", key_ge)
	headers.Add("X-GEMINI-PAYLOAD", payloadBase64)
	headers.Add("X-GEMINI-SIGNATURE", signature)
	/*headers := []byte(`
	{"X-GEMINI-APIKEY":` + string(key_ge) + `,
	 "X-GEMINI-PAYLOAD":` + string(b64) + `,
	 "X-GEMINI-SIGNATURE:` + string(sig) + `}`)*/
	return headers
}

func OrderSocket() (*websocket.Conn, *http.Response) {
	payload := map[string]interface{}{
		"request": "/v1/order/events",
		"nonce":   strconv.FormatInt(time.Now().UnixNano(), 10),
	}
	dialer := new(websocket.Dialer)
	conn, resp, err := dialer.Dial("wss://api.gemini.com/v1/order/events", auth_ge(payload))
	//j := json.NewDecoder(resp.Body)
	//s, _ := ioutil.ReadAll(j.Buffered())
	if err != nil {
		fmt.Println(resp.StatusCode)
		panic(err)
	}
	return conn, resp
}
