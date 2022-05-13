package AmazonSessionAPIClient

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/ajg/form"
	"github.com/corpix/uarand"
	"github.com/gocolly/colly/v2"
)

type CollyHandler struct {
	collector *colly.Collector

	proxy     string
	country   Country
	userAgent string
	cookies   map[string]string

	//amazonCaptchaSolver *AmazonCaptchaReader.CollyCaptchaFixHandler //Sadece server da kullaniyor.
}

func CreateCollyHandler(proxy string, country Country, userAgent string, cookies map[string]string) *CollyHandler {
	if userAgent == "" {
		userAgent = uarand.GetRandom()
	}

	if cookies == nil {
		cookies = make(map[string]string, 0)
	}

	if collector := CreateCollyCollector(proxy, country, userAgent, cookies); collector == nil {
		log.Panicln("Collector yaratilamadi!!")
		return nil
	} else {
		return &CollyHandler{proxy: proxy, country: country, userAgent: userAgent, cookies: cookies, collector: collector}
	}
}

func CreateCollyCollector(proxy string, country Country, userAgent string, cookies map[string]string) *colly.Collector {
	if userAgent == "" {
		return nil
	}

	if cookies == nil {
		return nil
	}

	newCookies := make([]*http.Cookie, 0)

	c := colly.NewCollector(
		colly.UserAgent(userAgent),
	)

	c.SetRequestTimeout(time.Second * 60)
	c.AllowURLRevisit = true
	c.Async = false

	if proxy != "" {
		err := c.SetProxy(proxy)

		if err != nil {
			log.Println("Set proxy de sorun olustu. ", err)
			return nil
		}
	}

	if cookies != nil {
		for key, value := range cookies {
			newCookies = append(newCookies, &http.Cookie{
				Name:  key,
				Value: value,
			})
		}

		c.SetCookies(country.Host(), newCookies)
	}

	return c
}

func (handler *CollyHandler) GetCookies() map[string]string {
	cookies := handler.cookies

	for _, cookie := range handler.collector.Cookies(handler.country.Host()) {
		cookies[cookie.Name] = cookie.Value
	}

	return cookies
}

func (handler *CollyHandler) GetUserAgent() string {
	return handler.userAgent
}

//Header ile visit
func (handler *CollyHandler) visit(URL string, header map[string][]string) error {
	captchaError := false
	defer handler.OnHTMLDetach("body > div > div.a-row.a-spacing-double-large > div.a-section > div > div > form")
	handler.OnHtml("body > div > div.a-row.a-spacing-double-large > div.a-section > div > div > form", func(h *colly.HTMLElement) {
		if h.Attr("action") == "/errors/validateCaptcha" {
			captchaError = true
		}
	})

	if handler.collector.CheckHead {
		if check := handler.collector.Request("HEAD", URL, nil, nil, header); check != nil {
			return check
		}
	}

	err := handler.collector.Request("GET", URL, nil, nil, header)

	if err != nil {
		return err
	}

	if captchaError {
		return GetError(CaptchaError)
	}

	return nil
}

func (handler *CollyHandler) VisitH(url string, header map[string][]string) error {
	handler.collector.AllowURLRevisit = true
	errCount := 0

	for handler.collector != nil {
		err := handler.visit(url, header)

		if err == nil {
			break
		} else if err != nil && err.Error() != http.StatusText(http.StatusServiceUnavailable) {
			//log.Println("Colly Get Err :", err, " URL", url)
			return err
		} else {
			errCount++

			if errCount == 6 {
				log.Println("Colly Get Err :", err, " URL", url)
				return err
			}

			time.Sleep(time.Second * 3)
		}
	}

	return nil
}

//TODO:sidepanel icin geriye 404 dondugunde dogru bir islem olarak gecicek.
func (handler *CollyHandler) Visit(url string) error {
	return handler.VisitH(url, nil)
}

func (handler *CollyHandler) VisitPost(url string, postData map[string]string, jsonContentType bool, headers map[string][]string) error {
	handler.collector.AllowURLRevisit = true

	errCount := 0
	var postData_ io.Reader

	captchaError := false
	defer handler.OnHTMLDetach("body > div > div.a-row.a-spacing-double-large > div.a-section > div > div > form")
	handler.OnHtml("body > div > div.a-row.a-spacing-double-large > div.a-section > div > div > form", func(h *colly.HTMLElement) {
		if h.Attr("action") == "/errors/validateCaptcha" {
			captchaError = true
		}
	})

	if jsonContentType {
		json_data, err := json.Marshal(postData)
		if err != nil {
			log.Println("Post icin verilen data json parse edilemedi. Err:", err)
			return err
		}
		postData_ = bytes.NewReader(json_data)
	} else {
		val, err := form.EncodeToValues(postData)

		if err != nil {
			log.Println("Post icin verilen data form parse edilemedi. Err:", err)
			return err
		}

		str := val.Encode()
		str = strings.ReplaceAll(str, "%2C+%22", "%2C%22")
		str = strings.ReplaceAll(str, "%3A+", "%3A")
		str = strings.ReplaceAll(str, "submit%5C.", "submit.") //Bu bir post icin ozel olarak konuldu.

		postData_ = strings.NewReader(str)
	}

	if headers != nil {
		headers["User-Agent"] = []string{handler.userAgent}
	}

	for handler.collector != nil {
		err := handler.collector.Request("POST", url, postData_, nil, headers)

		if err == nil {
			break
		} else if err != nil && err.Error() != http.StatusText(http.StatusServiceUnavailable) {
			//log.Println("Colly Get Err :", err, " URL", url)
			return err
		} else {
			errCount++

			if errCount == 6 {
				log.Println("Colly Get Err :", err, " URL", url)
				return err
			}

			time.Sleep(time.Second * 3)
		}
	}

	if captchaError {
		return GetError(CaptchaError)
	}

	return nil
}

func (handler *CollyHandler) OnHtml(goquerySelector string, f colly.HTMLCallback) {
	handler.collector.OnHTML(goquerySelector, f)
}

func (handler *CollyHandler) OnResponse(f colly.ResponseCallback) {
	handler.collector.OnResponse(f)
}

func (handler *CollyHandler) OnHTMLDetach(goquerySelector string) {
	handler.collector.OnHTMLDetach(goquerySelector)
}

//func (handler *CollyHandler) SetCaptchaSolver() {
//	if handler.amazonCaptchaSolver == nil {
//		handler.amazonCaptchaSolver = AmazonCaptchaReader.CreateHandler(handler.collector, handler.country.Host())
//		handler.amazonCaptchaSolver.Set()
//	}
//}
//
//func (handler *CollyHandler) DisableCaptchaSolver() {
//	if handler.amazonCaptchaSolver != nil {
//		handler.amazonCaptchaSolver.Disable()
//		handler.amazonCaptchaSolver = nil
//	}
//}
//
////Eger captcha olusturulmamissa false doner.
////Captcha basarili ise true doner, eger bir sorun olusmussa false doner.
//func (handler *CollyHandler) CaptchaStatus() bool {
//	if handler.amazonCaptchaSolver != nil {
//		return handler.amazonCaptchaSolver.CaptchaStatus()
//	}
//
//	return true
//}
