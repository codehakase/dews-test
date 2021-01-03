package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"io"
	"os"
)

// ReadUserFile reads the persisted binary file and returns a list of users
// stored
func ReadUserFile() ([]user, error) {
	file, err := os.Open("users.bin")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	users, err := handleFileBuffer(file)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// handleFileBuffer here decouples the I/O operation for ReadUserFile so one can
// test this method in isolation
func handleFileBuffer(file io.Reader) ([]user, error) {
	var users []user
	s := bufio.NewScanner(file)
	for s.Scan() {
		var u user

		buf := bytes.NewBuffer(s.Bytes())
		dec := gob.NewDecoder(buf)
		err := dec.Decode(&u)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

// WriteFile writes data of type User to a binary file
func WriteFile(data user) error {
	file, err := os.OpenFile("users.bin", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	// retrive bytes
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err = e.Encode(data)
	if err != nil {
		return err
	}

	_, err = b.WriteTo(file)
	if err != nil {
		return err
	}

	_, _ = file.Write([]byte("\n"))

	return nil
}
