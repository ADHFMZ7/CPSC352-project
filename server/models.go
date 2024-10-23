package main


type User struct {
  username string
  password string // Switch over to hashed pass eventually
}

type Session struct {
  user_ids []string // string of usernames
  symmetric_key [16]byte
}
