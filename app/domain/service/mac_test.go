package service

import (
	"ECPay/app/domain/model"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/fatih/structs"
)

func Test_GenerateCheckMacValue(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	v := &model.AioGeneralParam{
		MerchantID:        "2000132",
		MerchantTradeNo:   "ecpay20130312153023",
		MerchantTradeDate: "2013/03/12 15:30:23",
		PaymentType:       "aio",
		TotalAmount:       1000,
		TradeDesc:         "促銷方案",
		ItemName:          "Apple iphone 7 手機殼",
		ReturnURL:         "https://www.ecpay.com.tw/receive.php",
		ChoosePayment:     "ALL",
		EncryptType:       1,
	}

	m := structs.Map(v)

	s := NewService()
	mac := s.GenerateCheckMacValue(ctx, m)

	if mac == "CFA9BDE377361FBDD8F160274930E815D1A8A2E3E80CE7D404C45FC9A0A1E407" {
		fmt.Println("equal")
	} else {
		fmt.Println("err")
	}

	spew.Dump(mac)
}
