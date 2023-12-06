package header

// Header constants
const (
	AuthorizationType string = "Bearer"
	ContentType       string = "Content-Type"
	ContentLength     string = "Content-Length"
	Authorization     string = "Authorization"
	CompanyID         string = "X-Consumer-Custom-ID"
	SchemaName        string = "X-Consumer-Username"
	Authenticator     string = "X-Authenticated-Userid"
	UserScope         string = "X-Authenticated-Scope"
	XAPIKey           string = "X-API-KEY"
	ClientID          string = "Client-ID"
	UserIP            string = "User-Ip"
	UserAgent         string = "User-Agent"
	XForwardedFor     string = "X-Forwarded-For"
	XForwardedHost    string = "X-Forwarded-Host"
	XForwardedPort    string = "X-Forwarded-Port"
	XForwardedProto   string = "X-Forwarded-Proto"
	XRealIP           string = "X-Real-Ip"
	XRequestedWith    string = "X-Requested-With"
	Origin            string = "Origin"
	Referer           string = "Referer"
	CacheControl      string = "Cache-Control"
)

// Rate Limit
const (
	XRateLimitLimit     string = "X-RateLimit-Limit"
	XRateLimitRemaining string = "X-RateLimit-Remaining"
	XRateLimitReset     string = "X-RateLimit-Reset"
)

// RequestedWith constants
const (
	RequestedWithXMLHttpRequest string = "XMLHttpRequest"
)
