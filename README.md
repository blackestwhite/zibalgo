# Zibalgo

[![Run Tests](https://github.com/blackestwhite/zibalgo/actions/workflows/tests.yml/badge.svg)](https://github.com/blackestwhite/zibalgo/actions/workflows/tests.yml)

add package to your project:
```
go get github.com/blackestwhite/zibalgo
```

Example of creating and verifying payments:

```go
// Create a new ZibalClient instance
client := zibalgo.NewClient("zibal")

// Create a PaymentRequest instance
paymentRequest := PaymentRequest{
    CallbackURL: "https://example.com/callback",
    Description: "Test payment",
    Amount:      10000,
}

// Call the NewPayment method with a context
ctx := context.Background()
res, err := client.NewPayment(ctx, paymentRequest)
if err != nil {
    log.Println(err)
}

verificationRequest := VerificationRequest{
    TrackID: res.TrackID,
}
// Call the VerifyPayment method with a context
_, err = client.VerifyPayment(ctx, verificationRequest)
if err != nil {
    log.Println(err)
}
```