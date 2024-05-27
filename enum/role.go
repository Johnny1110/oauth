package enum

var (
	ROLE_SYS_ADMIN    = CodeValStruct{Code: 1, Val: "ROLE_SYS_ADMIN"}
	ROLE_SYS_RESOURCE = CodeValStruct{Code: 2, Val: "ROLE_SYS_RESOURCE"}
	ROLE_SYS_USER_L1  = CodeValStruct{Code: 3, Val: "ROLE_SYS_USER_L1"}
	ROLE_SYS_USER_L2  = CodeValStruct{Code: 4, Val: "ROLE_SYS_USER_L2"}
	ROLE_SYS_USER_L3  = CodeValStruct{Code: 5, Val: "ROLE_SYS_USER_L3"}
)

var SysRoles = []CodeValStruct{
	ROLE_SYS_ADMIN,
	ROLE_SYS_RESOURCE,
	ROLE_SYS_USER_L1,
	ROLE_SYS_USER_L2,
	ROLE_SYS_USER_L3,
}
