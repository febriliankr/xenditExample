package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	xendit "github.com/xendit/xendit-go/v4"
	"github.com/xendit/xendit-go/v4/payment_request"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	xenditApiKey := os.Getenv("XENDIT_API_KEY")
	xenditV4 := xendit.NewClient(xenditApiKey)

	/*
	QR Code creation V4
	*/
	var amount float64 = 3000
	var orderID = "dijd0q8je0182ej"

	// idempotencyKey is important to prevent double request, enabling retry
	var idempotencyKey = fmt.Sprintf("qr-%s", uuid.New().String())

	qrCode := payment_request.QRCodeParameters{
		ChannelCode: *payment_request.NewNullableQRCodeChannelCode(payment_request.QRCODECHANNELCODE_DANA.Ptr()),
	}

	paymentMethod := payment_request.PaymentMethodParameters{
		ReferenceId: &orderID,
		Type:        payment_request.PAYMENTMETHODTYPE_QR_CODE,
		Reusability: payment_request.PAYMENTMETHODREUSABILITY_ONE_TIME_USE,
		QrCode:      *payment_request.NewNullableQRCodeParameters(&qrCode),
	}

	paymentRequestParameters := payment_request.PaymentRequestParameters{
		Amount:        &amount,
		Currency:      "IDR",
		PaymentMethod: &paymentMethod,
	}

	createPRResp, _, createPRErr := xenditV4.PaymentRequestApi.
		CreatePaymentRequest(context.Background()).
		PaymentRequestParameters(paymentRequestParameters).
		IdempotencyKey(idempotencyKey).
		Execute()

	if createPRErr != nil {
		fmt.Fprintf(os.Stdout, "Create QR Payment Request Err`: %v\n", createPRErr)
	} else {
		fmt.Println(jsonStr(createPRResp))
		// qrString := createPRResp.PaymentMethod.QrCode.Get().ChannelProperties.QrString
		// expiresAt := createPRResp.PaymentMethod.QrCode.Get().ChannelProperties.QrString
	}

}

func jsonStr(data any) string {
	j, _ := json.Marshal(data)
	return string(j)
}
