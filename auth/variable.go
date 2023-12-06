package auth

var (
	/*
		Token expires in
	*/
	tokenExpiresIn = int64(3600)

	/*
		Refresh token expires in
	*/
	refreshTokenExpiresIn = int64(86400)

	/*
		Confirmation token expires in
	*/
	confirmationTokenExpiresIn = int64(1800)

	/*
		Reset password token expires in
	*/
	resetPasswordTokenExpiresIn = int64(1800)
)

var (
	/*
		Grant type with expires in
	*/
	grantTypeExpiresIn = map[string]int64{
		AccessToken:        tokenExpiresIn,
		RefreshToken:       refreshTokenExpiresIn,
		ConfirmationToken:  confirmationTokenExpiresIn,
		ResetPasswordToken: resetPasswordTokenExpiresIn,
	}
)
