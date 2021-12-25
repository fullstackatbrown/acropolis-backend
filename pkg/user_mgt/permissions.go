package user_mgt

const (
	// UserManagementPermission grants a user access to all user management features and resources.
	UserManagementPermission = "USER_MANAGEMENT_ALL"
	// UserManagementReadPermission allows a user to view user management features and user info.
	UserManagementReadPermission = "USER_MANAGEMENT_READ"
	// UserManagementEditUserInfoPermission allows a user to edit user information.
	UserManagementEditUserInfoPermission = "USER_MANAGEMENT_EDIT_USER_INFO"
	// UserManagementCreateUserPermission allows a user to create user accounts and send account invite emails.
	UserManagementCreateUserPermission = "USER_MANAGEMENT_CREATE_USER"
	// UserManagementDisableUserPermission allows a user to disable user accounts.
	UserManagementDisableUserPermission = "USER_MANAGEMENT_DISABLE_USER"
	// UserManagementDeleteUserPermission allows a user to delete user accounts.
	UserManagementDeleteUserPermission = "USER_MANAGEMENT_DELETE_USER"
)
