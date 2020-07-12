package main

import "errors"

func hash(password string, cost int) (hash string, err error) {
	err = errors.New("hash not implemented")
	return hash, err
}
