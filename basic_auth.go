package main

import (
	"crypto/subtle"
	"net/http"
)

func basicAuthMiddleware(username, password string, handler http.Handler) http.Handler {
	usernameBytes := []byte(username)
	passwordBytes := []byte(password)
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		receivedUsername, receivedPassword, _ := req.BasicAuth()
		receivedUsernameBytes := []byte(receivedUsername)
		receivedPasswordBytes := []byte(receivedPassword)

		counter := subtle.ConstantTimeCompare(usernameBytes, receivedUsernameBytes)
		counter += subtle.ConstantTimeCompare(passwordBytes, receivedPasswordBytes)

		if counter != 2 {
			writer.Header().Set("WWW-Authenticate", "Basic realm=\"Realm\", charset=\"UTF-8\"")
			writer.WriteHeader(http.StatusUnauthorized)
			writer.Write([]byte("401 Unauthorized"))
			return
		}

		handler.ServeHTTP(writer, req)
	})
}
