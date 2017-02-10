package ticket

var url = "https://api.weixin.qq.com/card/create?access_token=%s"

type CodeType string

const (
	None        CodeType = "CODE_TYPE_NONE"
	Text                 = "CODE_TYPE_TEXT"
	BarCode              = "CODE_TYPE_BARCODE"
	BarCodeOnly          = "CODE_TYPE_ONLY_BARCODE"
	QRCode               = "CODE_TYPE_QRCODE"
	QRCodeOnly           = "CODE_TYPE_ONLY_QRCODE"
)

type ColorType int

const (
/*
	Color010	ColorType = #63b359
Color020	=#2c9f67
Color030	=#509fc9
Color040	=#5885cf
Color050	=#9062c0
Color060	=#d09a45
Color070	=#e4b138
Color080	=#ee903c
Color081	=#f08500
Color082	=#a9d92d
Color090	=#dd6549
Color100	=#cc463d
Color101	=#cf3e36
Color102	=#5E6671
*/
)

type BaseTicket struct {
	LogoURL   string   `json:"logo_url"` // should be http://mmbiz.xxx
	CodeType  CodeType `json:"code_type"`
	BrandName string   `json:"brand_name"` // max 36
	Title     string   `json:"title"`      // max 27
	Color     string   `json:"color"`
}

type CreateReqeust struct {
}
