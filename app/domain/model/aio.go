package model

//
const (
	AIO = "aio"
)

// AioGeneralParam - create ecpay order base param
type AioGeneralParam struct {
	MerchantID        string                 `json:"MerchantID"`                  // 特店編號(由綠界提供)
	MerchantTradeNo   string                 `json:"MerchantTradeNo"`             // 特店交易編號(由特店提供、為唯一值)
	StoreID           *string                `json:"StoreID,omitempty"`           // 特店旗下店舖代號(分店)
	MerchantTradeDate string                 `json:"MerchantTradeDate"`           // 特店交易時間(格式: yyyy/MM/dd HH:mm:ss)
	PaymentType       string                 `json:"PaymentType"`                 // 交易類型(固定填入 `aio`)
	TotalAmount       int64                  `json:"TotalAmount"`                 // 交易金額
	TradeDesc         string                 `json:"TradeDesc"`                   // 交易描述(傳送到綠界前先做UrlEncode)
	ItemName          string                 `json:"ItemName"`                    // 商品名稱 (1. 如果商品名稱有多筆，需在金流選擇頁一行一行顯示商品名稱的話，商品名稱請以符號#分隔 2. 字數限制 400 字內，超過此限制系統將自動截斷)
	ReturnURL         string                 `json:"ReturnURL"`                   // 付款完成通知回傳網址(綠界會將付款結果參數以幕後(Server POST)回傳到該網址)    1. 請勿設定與 Client 端接收付款結果網址 OrderResultURL 相同位置，避免程式判斷錯誤。   2. 請在收到 Server 端付款結果通知後，請正確回應 1|OK 給綠界。
	ChoosePayment     ChoosePaymentEnum      `json:"ChoosePayment"`               // 選擇預設付款方式
	CheckMacValue     string                 `json:"CheckMacValue"`               // 檢查碼(特定機制產生)
	ClientBackURL     *string                `json:"ClientBackURL,omitempty"`     // Client端返回特店的按鈕連結(將頁面從綠界導回到此設定的網址(返回商店)) 1. 導回時不會帶付款結果到此網址，只是將頁面導回而已。
	ItemURL           *string                `json:"ItemURL,omitempty"`           // 商品銷售網址
	Remark            *string                `json:"Remark,omitempty"`            // 備註欄位
	ChooseSubPayment  *ChooseSubPaymentEnum  `json:"ChooseSubPayment,omitempty"`  // 付款子項目(若設定此參數，建立訂單將轉導至綠界訂單成立頁，依設定的付款方式及付款子項目帶入訂單，無法選擇其他付款子項目)
	OrderResultURL    *string                `json:"OrderResultURL,omitempty"`    // Client端回傳付款結果網址(當消費者付款完成後，綠界會將付款結果參數以幕前(Client POST)回傳到該網)  1. 若與[ClientBackURL]同時設定，將會以此參數為主  2. 銀聯卡及非即時交易(ATM、CVS、BARCODE)不支援此參數。
	NeedExtraPaidInfo *NeedExtraPaidInfoEnum `json:"NeedExtraPaidInfo,omitempty"` // 是否需要額外的付款資訊 => Y | N (若不回傳額外的付款資訊時，參數值請傳：`Ｎ`；   若要回傳額外的付款資訊時，參數值請傳：`Ｙ`，付款完成後綠界會以 Server POST 方式回傳額外付款資訊)
	DeviceSource      *string                `json:"DeviceSource,omitempty"`      // 裝置來源(請帶空值由綠界系統自動判定)
	IgnorePayment     *string                `json:"IgnorePayment,omitempty"`     // 隱藏付款 => 當付款方式 `ChoosePayment` 為 `ALL` 時，可隱藏不需要的付款方式，多筆請以井號分隔(#)  可用的參數值： -`Credit`: 信用卡 -`WebATM`: 網路 ATM  -`ATM`: 自動櫃員機  -`CVS`: 超商代碼  -`BARCODE`: 超商條碼
	PlatformID        *string                `json:"PlatformID,omitempty"`        // 特約合作平台商代號(由綠界提供) => 專案合作的平台商使用,一般特店或平台商本身介接，則參數請帶放空值。 若為專案合作平台商的特店使用時，則參數請帶平台商所綁的特店編號 `MerchantID`。
	InvoiceMark       *InvoiceMarkEnum       `json:"InvoiceMark,omitempty"`       // 電子發票開立註記(Y | N) 若要使用時，該參數須設定為「Y」，同時還要設定「電子發票介接相關參數」
	CustomField1      *string                `json:"CustomField1,omitempty"`      // 自訂名稱欄位(特殊符號只支援 `,.#()$[];%{}:/?&@<>!`)
	CustomField2      *string                `json:"CustomField2,omitempty"`      //
	CustomField3      *string                `json:"CustomField3,omitempty"`      //
	CustomField4      *string                `json:"CustomField4,omitempty"`      //
	EncryptType       EncryptTypeEnum        `json:"EncryptType"`                 // CheckMacValue 加密類型(固定填入 `1`)
	Language          *LanguageEnum          `json:"Language,omitempty"`          // 語系設定(預設繁體中文)
}

// GetTradeInfoRequest -
type GetTradeInfoRequest struct {
	MerchantID      string `json:"MerchantID"`      // 特店編號(由綠界提供)
	MerchantTradeNo string `json:"MerchantTradeNo"` // 特店交易編號(由特店提供、為唯一值)
	TimeStamp       int64  `json:"TimeStamp"`       // 時戳(3分鐘的有效期限)
	CheckMacValue   string `json:"CheckMacValue"`   // 檢查碼(特定機制產生)
}

// TradeInfo -
type TradeInfo struct {
	MerchantID           string                `json:"MerchantID" schema:"MerchantID"`                     // 特店編號(由綠界提供)
	MerchantTradeNo      string                `json:"MerchantTradeNo" schema:"MerchantTradeNo"`           // 特店交易編號(由特店提供)
	StoreID              string                `json:"StoreID" schema:"StoreID"`                           // 特店旗下店舖代號(分店)
	TradeNo              string                `json:"TradeNo" schema:"TradeNo"`                           // 綠界的交易編號(首次授權所產生的綠界交易編號)
	TradeAmt             int64                 `json:"TradeAmt" schema:"TradeAmt"`                         // 交易金額
	PaymentDate          string                `json:"PaymentDate" schema:"PaymentDate"`                   // 付款時間(格式: yyyy/MM/dd HH:mm:ss)
	PaymentType          ReturnPaymentTypeEnum `json:"PaymentType" schema:"PaymentType"`                   // 特店選擇的付款方式
	HandlingCharge       int64                 `json:"HandlingCharge"  schema:"HandlingCharge"`            // 手續費合計(履約結束後才會計算，未計算前為 0)
	PaymentTypeChargeFee float32               `json:"PaymentTypeChargeFee" schema:"PaymentTypeChargeFee"` // 通路費
	TradeDate            string                `json:"TradeDate" schema:"TradeDate"`                       // 訂單成立時間(格式: yyyy/MM/dd HH:mm:ss)
	TradeStatus          string                `json:"TradeStatus" schema:"TradeStatus"`                   // 交易狀態(若為 0 時，代表交易訂單成立未付款   若為 1 時，代表交易訂單成立已付款   若為 10200095 時，代表消費者未選擇付款方式，故交易失敗)
	ItemName             string                `json:"ItemName" schema:"ItemName"`                         // 商品名稱
	CustomField1         string                `json:"CustomField1" schema:"CustomField1"`                 // 自訂名稱欄位
	CustomField2         string                `json:"CustomField2" schema:"CustomField2"`                 //
	CustomField3         string                `json:"CustomField3" schema:"CustomField3"`                 //
	CustomField4         string                `json:"CustomField4" schema:"CustomField4"`                 //
	CheckMacValue        string                `json:"CheckMacValue" schema:"CheckMacValue"`               // 檢查碼(特定機制產生)
}

// ChoosePaymentEnum - 選擇預設付款方式
type ChoosePaymentEnum string

// ChoosePaymentEnum List
const (
	CHOOSEPAYMENTENUM_CREDIT  ChoosePaymentEnum = "Credit"  // 信用卡及銀聯卡(需申請開通)
	CHOOSEPAYMENTENUM_WEBATM  ChoosePaymentEnum = "WebATM"  // 網路 ATM(手機版時不支援)
	CHOOSEPAYMENTENUM_ATM     ChoosePaymentEnum = "ATM"     // 自動櫃員機
	CHOOSEPAYMENTENUM_CVS     ChoosePaymentEnum = "CVS"     // 超商代碼
	CHOOSEPAYMENTENUM_BARCODE ChoosePaymentEnum = "BARCODE" // 超商條碼
	CHOOSEPAYMENTENUM_ALL     ChoosePaymentEnum = "ALL"     // 不指定付款方式
)

// ChooseSubPaymentEnum - 付款子項目
type ChooseSubPaymentEnum string

// ChooseSubPaymentEnum List
const (
	CHOOSESUBPAYMENTENUM_TAISHIN    ChooseSubPaymentEnum = "TAISHIN"    // 台新(WEBATM、ATM)
	CHOOSESUBPAYMENTENUM_ESUN       ChooseSubPaymentEnum = "ESUN"       // 玉山(暫不提供)
	CHOOSESUBPAYMENTENUM_BOT        ChooseSubPaymentEnum = "BOT"        // 台灣銀行(WEBATM、ATM)
	CHOOSESUBPAYMENTENUM_FUBON      ChooseSubPaymentEnum = "FUBON"      // 台北富邦(暫不提供)
	CHOOSESUBPAYMENTENUM_CHINATRUST ChooseSubPaymentEnum = "CHINATRUST" // 中國信託(WEBATM、ATM)
	CHOOSESUBPAYMENTENUM_FIRST      ChooseSubPaymentEnum = "FIRST"      // 第一銀行(WEBATM、ATM)
	CHOOSESUBPAYMENTENUM_CATHAY     ChooseSubPaymentEnum = "CATHAY"     // 國泰世華(暫不提供)
	CHOOSESUBPAYMENTENUM_MEGA       ChooseSubPaymentEnum = "MEGA"       // 兆豐銀行(WEBATM)
	CHOOSESUBPAYMENTENUM_LAND       ChooseSubPaymentEnum = "LAND"       // 土地銀行(WEBATM、ATM)
	CHOOSESUBPAYMENTENUM_TACHONG    ChooseSubPaymentEnum = "TACHONG"    // 大眾銀行(WEBATM、ATM)
	CHOOSESUBPAYMENTENUM_SINOPAC    ChooseSubPaymentEnum = "SINOPAC"    // 永豐銀行(WEBATM)
	CHOOSESUBPAYMENTENUM_CVS        ChooseSubPaymentEnum = "CVS"        // 超商代碼繳款(CVS)
	CHOOSESUBPAYMENTENUM_OK         ChooseSubPaymentEnum = "OK"         // OK 超商代碼繳款(CVS)
	CHOOSESUBPAYMENTENUM_FAMILY     ChooseSubPaymentEnum = "FAMILY"     // 全家超商代碼繳款(CVS)
	CHOOSESUBPAYMENTENUM_HILIFE     ChooseSubPaymentEnum = "HILIFE"     // 萊爾富超商代碼繳款(CVS)
	CHOOSESUBPAYMENTENUM_IBON       ChooseSubPaymentEnum = "IBON"       // 7-11 ibon 代碼繳款(CVS)
	CHOOSESUBPAYMENTENUM_BARCODE    ChooseSubPaymentEnum = "BARCODE"    // 超商條碼繳款(BARCODE)
	CHOOSESUBPAYMENTENUM_EMPTY      ChooseSubPaymentEnum = ""           // Empty
)

// NeedExtraPaidInfoEnum - 是否需要額外的付款資訊
type NeedExtraPaidInfoEnum string

// NeedExtraPaidInfoEnum List
const (
	NEEDEXTRAPAIDINFOENUM_Y NeedExtraPaidInfoEnum = "Y"
	NEEDEXTRAPAIDINFOENUM_N NeedExtraPaidInfoEnum = "N"
)

// InvoiceMarkEnum 電子發票開立註記
type InvoiceMarkEnum string

// InvoiceMarkEnum List
const (
	INVOICEMARKENUM_Y InvoiceMarkEnum = "Y"
	INVOICEMARKENUM_N InvoiceMarkEnum = "N"
)

// EncryptTypeEnum - CheckMacValue加密類型(固定填入 `1`，使用 SHA256 加密)
type EncryptTypeEnum int

// EncryptTypeEnum List
const (
	ENCRYPTTYPEENUM_SHA256 EncryptTypeEnum = 1
)

// LanguageEnum 語系設定- 預設語系為中文
type LanguageEnum string

// LanguageEnum List
const (
	LANGUAGEENUM_ENG LanguageEnum = "ENG" // 英語
	LANGUAGEENUM_KOR LanguageEnum = "KOR" // 韓語
	LANGUAGEENUM_JPN LanguageEnum = "JPN" // 日語
	LANGUAGEENUM_CHI LanguageEnum = "CHI" // 簡體中文
)

// ReturnPaymentTypeEnum - 特店選擇的付款方式
type ReturnPaymentTypeEnum string

// ReturnPaymentTypeEnum List
const (
	RETURNPAYMENTTYPEENUM_WEB_ATM_TAISHIN    ReturnPaymentTypeEnum = "WebATM_TAISHIN"
	RETURNPAYMENTTYPEENUM_WEB_ATM_ESUN       ReturnPaymentTypeEnum = "WebATM_ESUN"
	RETURNPAYMENTTYPEENUM_WEB_ATM_BOT        ReturnPaymentTypeEnum = "WebATM_BOT"
	RETURNPAYMENTTYPEENUM_WEB_ATM_FUBON      ReturnPaymentTypeEnum = "WebATM_FUBON"
	RETURNPAYMENTTYPEENUM_WEB_ATM_CHINATRUST ReturnPaymentTypeEnum = "WebATM_CHINATRUST"
	RETURNPAYMENTTYPEENUM_WEB_ATM_FIRST      ReturnPaymentTypeEnum = "WebATM_FIRST"
	RETURNPAYMENTTYPEENUM_WEB_ATM_CATHAY     ReturnPaymentTypeEnum = "WebATM_CATHAY"
	RETURNPAYMENTTYPEENUM_WEB_ATM_MEGA       ReturnPaymentTypeEnum = "WebATM_MEGA"
	RETURNPAYMENTTYPEENUM_WEB_ATM_LAND       ReturnPaymentTypeEnum = "WebATM_LAND"
	RETURNPAYMENTTYPEENUM_WEB_ATM_TACHONG    ReturnPaymentTypeEnum = "WebATM_TACHONG"
	RETURNPAYMENTTYPEENUM_WEB_ATM_SINOPAC    ReturnPaymentTypeEnum = "WebATM_SINOPAC"
	RETURNPAYMENTTYPEENUM_ATM_TAISHIN        ReturnPaymentTypeEnum = "ATM_TAISHIN"
	RETURNPAYMENTTYPEENUM_ATM_ESUN           ReturnPaymentTypeEnum = "ATM_ESUN"
	RETURNPAYMENTTYPEENUM_ATM_BOT            ReturnPaymentTypeEnum = "ATM_BOT"
	RETURNPAYMENTTYPEENUM_ATM_FUBON          ReturnPaymentTypeEnum = "ATM_FUBON"
	RETURNPAYMENTTYPEENUM_ATM_CHINATRUST     ReturnPaymentTypeEnum = "ATM_CHINATRUST"
	RETURNPAYMENTTYPEENUM_ATM_FIRST          ReturnPaymentTypeEnum = "ATM_FIRST"
	RETURNPAYMENTTYPEENUM_ATM_LAND           ReturnPaymentTypeEnum = "ATM_LAND"
	RETURNPAYMENTTYPEENUM_ATM_CATHAY         ReturnPaymentTypeEnum = "ATM_CATHAY"
	RETURNPAYMENTTYPEENUM_ATM_TACHONG        ReturnPaymentTypeEnum = "ATM_TACHONG"
	RETURNPAYMENTTYPEENUM_CVS_CVS            ReturnPaymentTypeEnum = "CVS_CVS"
	RETURNPAYMENTTYPEENUM_CVS_OK             ReturnPaymentTypeEnum = "CVS_OK"
	RETURNPAYMENTTYPEENUM_CVS_FAMILY         ReturnPaymentTypeEnum = "CVS_FAMILY"
	RETURNPAYMENTTYPEENUM_CVS_HILIFE         ReturnPaymentTypeEnum = "CVS_HILIFE"
	RETURNPAYMENTTYPEENUM_CVS_IBON           ReturnPaymentTypeEnum = "CVS_IBON"
	RETURNPAYMENTTYPEENUM_BARCODE_BARCODE    ReturnPaymentTypeEnum = "BARCODE_BARCODE"
	RETURNPAYMENTTYPEENUM_CREDIT_CREDIT_CARD ReturnPaymentTypeEnum = "Credit_CreditCard"
)
