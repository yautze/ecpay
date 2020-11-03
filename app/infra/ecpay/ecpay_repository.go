package ecpay

import (
	"ECPay/app/domain/model"
	"ECPay/config"
	"fmt"
	"net/url"

	"github.com/gorilla/schema"
	"github.com/parnurzeal/gorequest"
	"github.com/sirupsen/logrus"
)

var decoder = schema.NewDecoder()

// GetECPayOrder -
func GetECPayOrder(in *model.GetTradeInfoRequest) *model.TradeInfo {
	// gorequest init
	req := gorequest.New()
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	apiURL := fmt.Sprintf("%s%s", config.ECPayHost, config.GETECPayOrder)
	// Post method
	req.Post(apiURL)

	// send post body
	req.Type("form")
	req.Send(in)

	// response
	_, res, errs := req.End()

	// err check
	for _, err := range errs {
		if err != nil {
			logrus.Error(err)
			break
		}
	}

	// parser query
	query, _ := url.ParseQuery(res)

	// dcode query to struct
	data := new(model.TradeInfo)
	_ = decoder.Decode(data, query)

	return data
}

// CreateECPayOrder -
func CreateECPayOrder(in interface{}) error {
	// gorequest init
	req := gorequest.New()
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	apiURL := fmt.Sprintf("%s%s", config.ECPayHost, config.CreateECPayOrder)
	// Post method
	req.Post(apiURL)

	// send post body
	req.Type("form")
	req.Send(in)

	// response
	_, res, errs := req.End()

	// err check
	for _, err := range errs {
		if err != nil {
			logrus.Error(err)
			return err
		}
	}

	fmt.Println(res)

	return nil
}
