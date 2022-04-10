package models

type ErrorCode int

type APIError struct {
	Code        ErrorCode `json:"code"`
	Description string    `json:"description"`
}

//Hatalari onlemek icin iota kullanilmadi.
//Struct ve const degerler birlestirilebilir. Error Code'a erismek icin UnknownError.Code gibi kullanilabilir.
const (
	UnknownError_Code        ErrorCode = 0
	NotEnoughSession_Code    ErrorCode = 1
	NotFoundSessionInfo_Code ErrorCode = 2
	InvalidQuantity_Code     ErrorCode = 3
	InvalidRequestCount_Code ErrorCode = 4
)

var UnknownError APIError = APIError{
	Code:        UnknownError_Code,
	Description: "Bilinmeyen bir hata",
}

var NotEnoughSession APIError = APIError{
	Code:        NotEnoughSession_Code,
	Description: "Kullanilabilir sayida session bulunamadi. Tum sessionlarin kullanim limiti dolmus, yeni sessionlar uretiyor yada proxy ile ilgili sorun olmus olabilir.",
}

var NotFoundSessionInfo APIError = APIError{
	Code:        NotFoundSessionInfo_Code,
	Description: "Session bilgisine ulasilmadi. Tekrar istek cagrilabilir.",
}

var InvalidQuantity APIError = APIError{
	Code:        NotFoundSessionInfo_Code,
	Description: "Request sayisi 0 dan buyuk olmali. min:1, max:sinirsiz. Not: quantity cok buyuk verilirse ayni request degerlerin gelme sansi var.",
}

var InvalidRequestCount APIError = APIError{
	Code:        InvalidRequestCount_Code,
	Description: "Request sayisi 0 dan buyuk olmali. min:1, max:sinirsiz.",
}
