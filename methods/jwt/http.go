package jwt

// func ExtractRequestAccessClaims(r *http.Request) (*Claims, error) {
// 	if r == nil {
// 		return nil, fmt.Errorf("nil http request")
// 	}
// 	return ExtractClaims(r.Context(), ExtractAccessToken(r))
// }

// func ExtractRequestRefreshClaims(r *http.Request) (*Claims, error) {
// 	if r == nil {
// 		return nil, fmt.Errorf("nil http request")
// 	}
// 	return ExtractClaims(r.Context(), ExtractRefreshToken(r))
// }

// func ExtractAccessToken(r *http.Request) string {
// 	if r == nil {
// 		return ""
// 	}
// 	return r.Header.Get("")
// }

// func ExtractRefreshToken(r *http.Request) string {
// 	if r == nil {
// 		return ""
// 	}
// 	cookie, err := r.Cookie("")
// 	if err != nil || cookie == nil {
// 		return ""
// 	}
// 	return cookie.Value
// }
