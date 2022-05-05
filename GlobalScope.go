package AmazonSessionAPIClient

type Country struct {
	country        string //Ulke ismi
	currency       string //Para birimi
	host           string //URL
	tag            string //Kisaltilmis ulke ismi
	exampleZipCode string //Ornek Ulke Zip kodu
	exampleCity    string //Ornek Ulke
}

func (c Country) Country() string {
	return c.country
}

func (c Country) Currency() string {
	return c.currency
}

func (c Country) Host() string {
	return c.host
}

func (c Country) Tag() string {
	return c.tag
}

func (c Country) ExampleZipCode() string {
	return c.exampleZipCode
}

func (c Country) ExampleCity() string {
	return c.exampleCity
}

//TODO: ZipCode ve city rasgele verildi ileride duzeltilmesi gerekebilir.
var AMAZON_COUNTRIES map[string]Country = map[string]Country{
	"US": AcquireCountry("United States of America", "USD", "https://www.amazon.com", "US", "10007", ""),
	"AU": AcquireCountry("Australia", "AUD", "https://www.amazon.com.au", "AU", "0872", "HALE"),
	"BR": AcquireCountry("Brazil", "BRL", "https://www.amazon.com.br", "BR", "57320-000", ""),
	"CA": AcquireCountry("Canada", "CAD", "https://www.amazon.ca", "CA", "V5Y 3E3", ""),
	//"CN": AcquireCountry("China", "CNY", "https://www.amazon.cn", "CN", "", ""),//kullanilamaz
	"FR": AcquireCountry("France", "EUR", "https://www.amazon.fr", "FR", "75001", ""),
	"DE": AcquireCountry("Germany", "EUR", "https://www.amazon.de", "DE", "10115", ""),
	//"IN": AcquireCountry("India", "INR", "https://www.amazon.in", "IN", "", ""),//kullanilamaz
	"IT": AcquireCountry("Italy", "EUR", "https://www.amazon.it", "IT", "00118", ""),
	"MX": AcquireCountry("Mexico", "MXN", "https://www.amazon.com.mx", "MX", "22056", ""),
	//"NL": AcquireCountry("Netherlands", "EUR", "https://www.amazon.nl", "NL", "", ""), //kullanilamaz
	//"SG": AcquireCountry("Singapore", "SGD", "https://www.amazon.sg", "SG", "", ""), //kullanilamaz
	//"ES": AcquireCountry("Spain", "EUR", "https://www.amazon.es", "ES", "", ""), //kullanilamaz
	//"TR": AcquireCountry("Turkey", "TRY", "https://www.amazon.com.tr", "TR", "", ""), //kullanilamaz
	//"AE": AcquireCountry("United Arab Emirates", "AED", "https://www.amazon.ae", "AE", "", ""), //kullanilamaz
	"GB": AcquireCountry("United Kingdom", "GBP", "https://www.amazon.co.uk", "GB", "GL7 1WQ", ""),
	//"JP": AcquireCountry("Japan", "JPY", "https://www.amazon.jp", "JP", "", ""), //kullanilamaz
}

func AcquireCountry(country_, currency_, host_, tag_ string, exampleZipCode_ string, exampleCity_ string) Country {
	return Country{
		country:        country_,
		currency:       currency_,
		host:           host_,
		tag:            tag_,
		exampleZipCode: exampleZipCode_,
		exampleCity:    exampleCity_,
	}
}

//Verilen country uygun olup olmadigini kontrol eder.
func CheckCountry(country string) bool {
	for key := range AMAZON_COUNTRIES {
		if key == country {
			return true
		}
	}

	return false
}
