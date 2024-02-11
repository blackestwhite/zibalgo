package zibalgo

import "net/http"

const (
	// Result codes
	SuccessCode          = 100
	MerchantNotFoundCode = 102
	MerchantInactiveCode = 103
	MerchantInvalidCode  = 104
	AlreadyVerifiedCode  = 201
	InvalidAmountCode    = 105
	InvalidCallbackCode  = 106
	AmountExceedsLimit   = 113

	// baseURL is the base URL of the Zibal API.
	baseURL = "https://gateway.zibal.ir/"

	// requestPath is the endpoint path for requesting payments.
	requestPath = "v1/request"

	// verifyPath is the endpoint path for verifying payments.
	verifyPath = "v1/verify"
)

var (
	// Error messages
	ResultMessages = map[int]string{
		SuccessCode:          "با موفقیت تایید شد.",
		MerchantNotFoundCode: "merchant یافت نشد.",
		MerchantInactiveCode: "merchant غیر فعال",
		MerchantInvalidCode:  "merchant نامعتبر",
		AlreadyVerifiedCode:  "قبلا تایید شده.",
		InvalidAmountCode:    "amount بایستی بزرگتر از 1,000 ریال باشد.",
		InvalidCallbackCode:  "callbackUrl نامعتبر می‌باشد. (شروع با http و یا https)",
		AmountExceedsLimit:   "amount مبلغ تراکنش از سقف میزان تراکنش بیشتر است.",
	}
)

// ZibalClient represents a client for interacting with the Zibal Payment Gateway.
type ZibalClient struct {
	// httpClient is used to send HTTP requests.
	httpClient *http.Client

	// merchat is the authentication token used for API requests.
	merchant string
}

type PaymentRequest struct {
	Merchant    string `json:"merchant"`
	CallbackURL string `json:"callbackUrl"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}

type PaymentResponse struct {
	Result  int `json:"result"`
	TrackID int `json:"trackId"`
}

type VerificationRequest struct {
	Merchant string `json:"merchant"`
	TrackID  string `json:"trackId"`
}