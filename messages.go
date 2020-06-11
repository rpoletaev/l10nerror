package main

import "github.com/nicksnyder/go-i18n/v2/i18n"

var messages = []i18n.Message{
	{
		ID:          "invoices",
		Description: "The number of invoices a person has",
		One:         "You can {{.Count}} invoice",
		Other:       "You have {{.Count}} invoices",
	},
	{
		ID:          "ErrAlreadyExists",
		Description: "Error already exists",
		One:         "already exists: {{.Entity}}",
		Other:       "already exists: {{.Entity}}",
	},
}
