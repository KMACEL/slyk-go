package slyk

const (
	apiSlyk = "https://api.slyk.io"

	linkUsers      = apiSlyk + "/users"
	approve        = "/approve"
	block          = "/block"
	unblock        = "/unblock"
	changePassword = "/change-password"

	linkWallets  = apiSlyk + "/wallets"
	activity     = "/activity"
	balance      = "/balance"
	movements    = "/movements"
	transactions = "/transactions"

	linkTransactions   = apiSlyk + "/transactions"
	linkAddresses      = apiSlyk + "/addresses"
	linkAssets         = apiSlyk + "/assets"
	linkRates          = apiSlyk + "/rates"
	linkPaymentMethods = apiSlyk + "/payment-methods"
	linkMovements      = apiSlyk + "/movements"
	linkInvites        = apiSlyk + "/invites"
	linkCategories     = apiSlyk + "/categories"
	linkOrders         = apiSlyk + "/orders"
	linkProducts       = apiSlyk + "/products"
	linkQuestions      = apiSlyk + "/questions"
	linkQuestionTypes  = apiSlyk + "/question-types"
)

const (
	headerApiKey = "apikey"
)
