package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/atotto/clipboard"
	"os"
	"strconv"
	"strings"
)

type Profile struct {
	ProfileName string
	Email       string
	Phone       string
	SameBilling bool
	Shipping    ShippingAddress
	Billing     BillingAddress
	Card        CreditCard
}

type CreditCard struct {
	CardHolder  string
	Type        string
	Number      string
	ExpiryMonth string
	ExpiryYear  string
	Cvv         string
}

type ShippingAddress struct {
	FirstName  string
	LastName   string
	Address1   string
	Address2   string
	PostalCode string
	City       string
	State      string
	Country    string
}

type BillingAddress struct {
	FirstName  string
	LastName   string
	Address1   string
	Address2   string
	PostalCode string
	City       string
	State      string
	Country    string
}

func main() {
	fmt.Println("Input the directory of your file")
	fmt.Println("---------------------")
	var dir string
	fmt.Scanln(&dir)
	file, err := os.Open(dir)
	if err != nil {
		println("The file directory does not exist")
		return
	}

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		println("There was an error parsing the file")
		return
	}

	var profiles []Profile
	for col, row := range records {
		if col == 0 {
			continue
		}
		var profile Profile
		for section, cell := range row {
			switch section {
			case 0:
				profile.Email = cell
			case 1:
				profile.ProfileName = cell
			case 2:
			case 3:
				profile.Card.CardHolder = cell
			case 4:
				profile.Card.Type = cell
			case 5:
				profile.Card.Number = cell
			case 6:
				profile.Card.ExpiryMonth = cell
			case 7:
				profile.Card.ExpiryYear = cell
			case 8:
				profile.Card.Cvv = cell
			case 9:
				profile.SameBilling, _ = strconv.ParseBool(cell)
			case 10:
				fullName := strings.Split(cell, " ")
				profile.Shipping.FirstName = fullName[0]
				profile.Shipping.LastName = fullName[1]
				profile.Billing.FirstName = fullName[0]
				profile.Billing.LastName = fullName[1]
			case 11:
				profile.Phone = cell
			case 12:
				profile.Shipping.Address1 = cell
				profile.Billing.Address1 = cell
			case 13:
				profile.Shipping.Address2 = cell
				profile.Billing.Address2 = cell
			case 14:
			case 15:
				profile.Shipping.PostalCode = cell
				profile.Billing.PostalCode = cell
			case 16:
				profile.Shipping.City = cell
				profile.Billing.City = cell
			case 17:
				profile.Shipping.State = cell
				profile.Billing.State = cell
			case 18:
				profile.Shipping.Country = cell
				profile.Billing.Country = cell
			}
		}
		profiles = append(profiles, profile)
	}
	output, _ := json.Marshal(profiles)
	clipboard.WriteAll(string(output))
	clipboard.ReadAll()
	fmt.Println("Copied JSON to clipboard")
}
