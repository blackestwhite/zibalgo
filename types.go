package zibalgo

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
