package AmazonSessionAPIClient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

var APISERVER_HOST string
var APISERVER_PORT string

var client *http.Client = &http.Client{}

const APIVersion string = "v1"

//WARNING:clearCart ve lockSession ayni anda cagrilamaz. => os.Exit(1)
func GetColly(targetHostCountry string, deliveryCountry string, updateSession bool, requestCount int, clearCart bool, sessionLock bool) *CollyHandler {
	sessinInfo := GetAmazonSessionReq(targetHostCountry, deliveryCountry, updateSession, requestCount, clearCart, sessionLock)

	if sessinInfo == nil {
		PrintLog(ERROR_LOGS, "AmazonSessionAPIClient'de sorun var sessionInfo bos geldi.")
		return nil
	}

	return CreateCollyHandler(sessinInfo.Proxy, AMAZON_COUNTRIES[targetHostCountry], sessinInfo.UserAgent, sessinInfo.Cookies)
}

func GetCollyX(targetHostCountry string, updateSession bool, requestCount int, clearCart bool) *CollyHandler {
	sessinInfo := GetAmazonSessionReqX(targetHostCountry, updateSession, requestCount, clearCart)

	if sessinInfo == nil {
		PrintLog(ERROR_LOGS, "AmazonSessionAPIClient'de sorun var sessionInfo bos geldi.")
		return nil
	}

	return CreateCollyHandler(sessinInfo.Proxy, AMAZON_COUNTRIES[targetHostCountry], sessinInfo.UserAgent, sessinInfo.Cookies)
}

//Random bir amazon session dondurur. Serverdan bir istek beklendigi icin bekleme olabilir. Basarili olmaya calisir.
//Sadece targetHost ile islem yapilir delivery country random gelir.
//requestCount: kac kere kullanilacak.
func GetAmazonSessionReqX(targetHostCountry string, updateSession bool, requestCount int, clearCart bool) *SessionInfo {
	for {
		apiJson, _ := getAPIData("getSession", map[string]string{"clearCart": strconv.FormatBool(clearCart), "updateSession": strconv.FormatBool(updateSession), "requestCount": fmt.Sprint(requestCount), "targetHostCountry": targetHostCountry})

		if apiJson != nil {
			var sessionInfo *SessionInfo = &SessionInfo{Code: -1}

			if err := json.Unmarshal(apiJson, sessionInfo); err == nil {
				if sessionInfo.Code == 0 {
					return sessionInfo
				}
			}
		}

		time.Sleep(10 * time.Second)
	}
}

//Random bir amazon session dondurur. Serverdan bir istek beklendigi icin bekleme olabilir. Basarili olmaya calisir.
//requestCount: kac kere kullanilacak.
//WARNING:clearCart ve lockSession ayni anda cagrilamaz. => os.Exit(1)
func GetAmazonSessionReq(targetHostCountry string, deliveryCountry string, updateSession bool, requestCount int, clearCart bool, lockSession bool) *SessionInfo {
	if clearCart && lockSession {
		PrintLog(ERROR_LOGS, "AmazonSessionAPIClient'de clearCart ve lockSession ayni anda cagrilamaz.")
		os.Exit(1)
		return nil
	}

	for {
		apiJson, _ := getAPIData("getSession", map[string]string{"clearCart": strconv.FormatBool(clearCart), "lockSession": strconv.FormatBool(lockSession), "updateSession": strconv.FormatBool(updateSession), "requestCount": fmt.Sprint(requestCount), "targetHostCountry": targetHostCountry, "deliveryCountry": deliveryCountry})

		if apiJson != nil {
			var sessionInfo *SessionInfo = &SessionInfo{Code: -1}

			if err := json.Unmarshal(apiJson, sessionInfo); err == nil {
				if sessionInfo.Code == 0 {
					return sessionInfo
				}
			}
		}

		time.Sleep(10 * time.Second)
	}
}

func GetServerInfo() *ServerInfo {
	apiJson, _ := getAPIData("getServerInfo", map[string]string{})

	if apiJson != nil {
		var serverInfo *ServerInfo = &ServerInfo{}

		if err := json.Unmarshal(apiJson, serverInfo); err == nil {
			return serverInfo
		}
	}

	return nil
}

func FeedbackClearCart(sessionInfo *SessionInfo) {
	sendErrorCount := 0

	sessionID := sessionInfo.Cookies["session-id"]

	if sessionID == "" {
		PrintLog(ERROR_LOGS, "FeedbackClearCart da session-id bulunamadi.")
		return
	}

	for sendErrorCount <= 5 {
		_, ok := getAPIData("feedBackClearCart", map[string]string{"sessionID": sessionID})

		if ok {
			break
		} else {
			sendErrorCount++
			time.Sleep(time.Second * 1)
		}
	}
}

func FeedbackUnlockSession(sessionInfo *SessionInfo) {
	sendErrorCount := 0

	sessionID := sessionInfo.Cookies["session-id"]

	if sessionID == "" {
		PrintLog(ERROR_LOGS, "FeedbackUnlockSession da session-id bulunamadi.")
		return
	}

	for sendErrorCount <= 5 {
		_, ok := getAPIData("feedBackUnlockSession", map[string]string{"sessionID": sessionID})

		if ok {
			break
		} else {
			sendErrorCount++
			time.Sleep(time.Second * 1)
		}
	}
}

//API baglantisi yapar ve istenilen verileri alir
//apiTarget => getSession, getSessions, getInfo
//apiParam => updateSession=true, quantity=31, country=US
func getAPIData(apiTarget string, apiParam map[string]string) ([]byte, bool) {
	///api/v1/
	apiURL, err := url.Parse(APISERVER_HOST + ":" + APISERVER_PORT + "/api/" + APIVersion + "/" + apiTarget)
	if err != nil {
		PrintLog(ERROR_LOGS, "Api url parse error: ", err)
		return nil, false
	}

	params := url.Values{}

	for key, value := range apiParam {
		params.Add(key, value)
	}

	apiURL.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", apiURL.String(), nil)

	if err != nil {
		return nil, false
	}

	//Do
	resp, err := client.Do(req)

	if err != nil {
		return nil, false
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, false
	}

	return body, resp.StatusCode == 200
}
