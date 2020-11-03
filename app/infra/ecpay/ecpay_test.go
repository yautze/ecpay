package ecpay

import (
	"ECPay/app/domain/model"
	"ECPay/app/domain/service"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/fatih/structs"
)

func Test_Create(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	v := &model.GetTradeInfoRequest{
		MerchantID:      "2000132",
		MerchantTradeNo: "A005337",
		TimeStamp:       time.Now().Unix(),
	}

	m := structs.Map(v)

	s := service.NewService()
	v.CheckMacValue = s.GenerateCheckMacValue(ctx, m)

	d := GetECPayOrder(v)

	fmt.Println(d)
}
