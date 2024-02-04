# Xendit V4 API Example with Go

## QR Code

```go
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
```
