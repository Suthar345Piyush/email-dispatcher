package main

import (
	"bytes"
	"html/template"
	"sync"
)

type Recipient struct {
	Name  string
	Email string
}

func main() {

	// unbuffered channel

	recipientChannel := make(chan Recipient)

	// putting the loadRecipient into goroutine

	go func() {
		loadRecipient("./dummyEmails.csv", recipientChannel)
	}()

	// stopping program until all the emails are processed using waitgroups

	var wg sync.WaitGroup

	// email worker goroutine , ran by 5th worker

	workerCount := 5

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go emailWorker(i, recipientChannel, &wg)
	}

	wg.Wait()

}

// execution of the .tmpl file

func executeTemplate(r Recipient) (string, error) {
	t, err := template.ParseFiles("email.tmpl")

	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer

	err = t.Execute(&tpl, r)

	if err != nil {
		return "", err
	}

	return tpl.String(), nil

}
