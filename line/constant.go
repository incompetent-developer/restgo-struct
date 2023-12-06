package line

/*
	Path
*/
const (
	/*
		For API
	*/
	PushMessagePath    string = "%s/%s/bot/message/push"  // for push message
	GetAllFollowerPath string = "%s/%s/bot/followers/ids" // for get all followers
	GetUserProfilePath string = "%s/%s/bot/profile/%s"    // for get one user frofile
)

/*
	System Notification ID
*/
const (
	NotifyBooking string = "notify-booking"
)
