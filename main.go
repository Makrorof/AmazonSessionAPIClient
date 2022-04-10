package AmazonSessionAPIClient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/Makrorof/AmazonSessionAPIClient/models"
)

var APISERVER_HOST string
var APISERVER_PORT string

var client *http.Client = &http.Client{}

func APIVersion() string {
	return "v1"
}

//Random bir amazon session dondurur. Serverdan bir istek beklendigi icin bekleme olabilir. Basarili olmaya calisir.
//requestCount: kac kere kullanilacak.
func GetAmazonSessionReq(targetHostCountry string, deliveryCountry string, updateSession bool, requestCount int) *models.SessionInfo {
	for {
		apiJson := getAPIData("getSession", map[string]string{"updateSession": strconv.FormatBool(updateSession), "requestCount": fmt.Sprint(requestCount), "targetHostCountry": targetHostCountry, "deliveryCountry": deliveryCountry})

		if apiJson != nil {
			var sessionInfo *models.SessionInfo = &models.SessionInfo{Code: -1}

			if err := json.Unmarshal(apiJson, sessionInfo); err == nil {
				if sessionInfo.Code == 0 { //TODO:buraya panel icin log eklenecek.
					return sessionInfo
				}
			}
		}

		time.Sleep(10 * time.Second)
	}
}

func GetServerInfo() *models.ServerInfo {
	apiJson := getAPIData("getServerInfo", map[string]string{})

	if apiJson != nil {
		var serverInfo *models.ServerInfo = &models.ServerInfo{}

		if err := json.Unmarshal(apiJson, serverInfo); err == nil {
			return serverInfo
		}
	}

	return nil
}

//Kaldirildi.//
//Random bir amazon sessionlari dondurur. Serverdan bir istek beklendigi icin bekleme olabilir.
//func GetAmazonSessionsReq(targetHostCountry string, deliveryCountry string, updateSession bool, quantity int) []*models.SessionInfo {
//	for {
//		apiJson := getAPIData("getSessions", map[string]string{"updateSession": strconv.FormatBool(updateSession), "targetHostCountry": targetHostCountry, "deliveryCountry": deliveryCountry, "quantity": strconv.Itoa(quantity)})
//
//		if apiJson != nil {
//			var sessionInfoList []*models.SessionInfo
//
//			if err := json.Unmarshal(apiJson, &sessionInfoList); err == nil {
//				if len(sessionInfoList) > 0 { //TODO:buraya panel icin log eklenecek.
//					return sessionInfoList
//				}
//			}
//		}
//
//		time.Sleep(10 * time.Second)
//	}
//}

//API baglantisi yapar ve istenilen verileri alir
//apiTarget => getSession, getSessions, getInfo
//apiParam => updateSession=true, quantity=31, country=US
func getAPIData(apiTarget string, apiParam map[string]string) []byte {
	///api/v1/
	url := APISERVER_HOST + ":" + APISERVER_PORT + "/api/" + APIVersion() + "/" + apiTarget

	if len(apiParam) != 0 {
		url += "?"

		for key, value := range apiParam {
			url += key + "=" + value
			url += "&"
		}

		url = url[0 : len(url)-1]
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil
	}

	//Do
	resp, err := client.Do(req)

	if err != nil {
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil
	}

	return body
}
