package AmazonSessionAPIClient

import "errors"

type Error int

type APIError struct {
	Code        Error  `json:"code"`
	Description string `json:"description"`
}

const (
	OkError                        Error = 0
	UnknownError                   Error = 1
	NotEnoughSession               Error = 2
	NotFoundSessionInfo            Error = 3
	InvalidQuantity                Error = 4
	InvalidRequestCount            Error = 5
	UnknownCountry                 Error = 6
	UpdateSessionFunctionNotFound  Error = 7
	AmazonAddressBlock             Error = 8
	AmazonSessionIDNotFound        Error = 9
	AmazonCaptchaMaxReVisit        Error = 10
	AmazonCollyCollectorNotCreated Error = 11
	NotFound                       Error = 12
	NotSupportedCountry            Error = 13
)

var statusText = map[Error]string{
	UnknownError:                   "Bilinmeyen bir hata",
	NotEnoughSession:               "Kullanilabilir sayida session bulunamadi. Tum sessionlarin kullanim limiti dolmus, yeni sessionlar uretiyor yada proxy ile ilgili sorun olmus olabilir.",
	NotFoundSessionInfo:            "Session bilgisine ulasilmadi. Tekrar istek cagrilabilir.",
	InvalidQuantity:                "Request sayisi 0 dan buyuk olmali. min:1, max:sinirsiz. Not: quantity cok buyuk verilirse ayni request degerlerin gelme sansi var.",
	InvalidRequestCount:            "Request sayisi 0 dan buyuk olmali. min:1, max:sinirsiz.",
	UnknownCountry:                 "Verilen bolge yanlis.",
	UpdateSessionFunctionNotFound:  "Update Session Function Not Found",
	AmazonAddressBlock:             "Amazon Address Block",
	AmazonSessionIDNotFound:        "Amazon Session ID Not Found",
	AmazonCaptchaMaxReVisit:        "Amazon Captcha Max ReVisit",
	AmazonCollyCollectorNotCreated: "Amazon Colly Collector Not Created",
	NotFound:                       "Not Found",
	NotSupportedCountry:            "Verdiginiz bolge desteklenmiyor. Lutfen verdiginiz bolgeleri parametre olarak servera ekleyiniz.",
}

func GetAPIError(errType Error) APIError {
	return APIError{Code: errType, Description: statusText[errType]}
}

func GetError(errType Error) error {
	return errors.New(statusText[errType])
}

func GetErrorType(err error) Error {
	for index, status := range statusText {
		if status == err.Error() {
			return index
		}
	}

	return -1
}
