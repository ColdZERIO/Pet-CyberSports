package service

import (
	"log"
	"testing"
)

func TestHashPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		params   ArgonParams
	}{
		{name: "Easy pass", password: "123"},
		{name: "Not easy", password: "qwerty123"},
		{name: "Hard pass", password: "@dsnhjvbshb1234561"},
		{name: "Super hard pass", password: "njk421-bnjvf#5321-fbhb34b-njcu!^"},
	}

	for i, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			DefaultParams := ArgonParams{
				Memory:      64 * 1024,
				Time:        3,
				Parallelism: 2,
				SaltLength:  16,
				KeyLength:   32,
			}
			res, err := HashPassword(test.password, DefaultParams)
			if err != nil {
				log.Println(err)
			}
			log.Printf("Test №%d - %s\n", i, res)
		})
	}
}
