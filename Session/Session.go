package Session

type SessionData struct {
	username string
}

type Session struct {
	data map[string]*SessionData
}
