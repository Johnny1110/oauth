package enum

var (
	FRIZO_USER = CodeValStruct{Code: 1, Val: "FRIZO_USER"}
	FRIZO_ADMIN   = CodeValStruct{Code: 2, Val: "FRIZO_ADMIN"}
	FRIZO_RESOURCE = CodeValStruct{Code: 3, Val: "FRIZO_RESOURCE"}
)

var systemTypes = []CodeValStruct{
	FRIZO_USER,
	FRIZO_ADMIN,
	FRIZO_RESOURCE,
}

func GetSystemTypeByCode(code int) *CodeValStruct {
	for _, sc := range systemTypes {
		if sc.Code == code {
			return &sc
		}
	}
	return nil
}

func GetSystemTypeByVal(val string) *CodeValStruct {
	for _, sc := range systemTypes {
		if sc.Val == val {
			return &sc
		}
	}
	return nil
}
