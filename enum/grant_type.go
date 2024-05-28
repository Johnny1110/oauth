package enum

var (
	PASSWORD           = CodeValStruct{Code: 1, Val: "password"}
	REFRESH_TOKEN      = CodeValStruct{Code: 2, Val: "refresh_token"}
	AUTHORIZATION_CODE = CodeValStruct{Code: 3, Val: "authorization_code"}
	CLIENT_CREDENTIALS = CodeValStruct{Code: 4, Val: "client_credentials"}
	IMPLICIT           = CodeValStruct{Code: 5, Val: "implicit"}
)

var grantTypes = []CodeValStruct{
	PASSWORD,
	REFRESH_TOKEN,
	AUTHORIZATION_CODE,
	CLIENT_CREDENTIALS,
	IMPLICIT,
}
