package ecpay

import (
	"ECPay/app/domain/model"
	"ECPay/app/domain/service"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/fatih/structs"
)

func Test_GET(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	v := &model.GetTradeInfoRequest{
		MerchantID:      "2000132",
		MerchantTradeNo: "1326437583814660096",
		TimeStamp:       time.Now().Unix(),
	}

	m := structs.Map(v)

	s := service.NewService()
	v.CheckMacValue = s.GenerateCheckMacValue(ctx, m)

	d := GetECPayOrder(v)

	fmt.Println(d)
}

func Test_Create(t *testing.T) {
	node, _ := snowflake.NewNode(1)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	v := &model.AioGeneralParam{
		MerchantID:        "2000132",
		MerchantTradeNo:   node.Generate().String(),
		MerchantTradeDate: time.Now().Format("2006/01/02 15:04:05"),
		PaymentType:       "aio",
		TotalAmount:       1000,
		TradeDesc:         "YauTz",
		ItemName:          "Apple",
		ReturnURL:         "http://your.web.site/receive.php",
		ChoosePayment:     model.CHOOSEPAYMENTENUM_ALL,
		EncryptType:       model.ENCRYPTTYPEENUM_SHA256,
	}

	m := structs.Map(v)

	s := service.NewService()
	v.CheckMacValue = s.GenerateCheckMacValue(ctx, m)

	d := CreateECPayOrder(v)

	fmt.Println(d)
}
