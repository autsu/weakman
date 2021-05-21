package test

import (
	"log"
	"testing"
	service2 "vote/trash/internal/v1/service"
)

func TestAdminLogin(t *testing.T) {
	token, err := service2.AdminLogin("root", "root")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(token)
}