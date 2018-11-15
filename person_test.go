package fastfaker

import (
	"fmt"
	"testing"
)

func ExampleFaker_SSN() {
	Global.Seed(11)
	fmt.Println(Global.SSN())
	// Output: 296446360
}

func BenchmarkSSN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.SSN()
	}
}

func ExampleFaker_Gender() {
	Global.Seed(11)
	fmt.Println(Global.Gender())
	// Output: female
}

func BenchmarkGender(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.Gender()
	}
}

func ExampleFaker_Person() {
	Global.Seed(11)
	person := Global.Person()
	job := person.Job
	address := person.Address
	contact := person.Contact
	creditCard := person.CreditCard

	fmt.Println(person.FirstName)
	fmt.Println(person.LastName)
	fmt.Println(person.Gender)
	fmt.Println(person.SSN)
	fmt.Println(person.Image)

	fmt.Println(job.Company)
	fmt.Println(job.Title)
	fmt.Println(job.Descriptor)
	fmt.Println(job.Level)

	fmt.Println(address.Address)
	fmt.Println(address.Street)
	fmt.Println(address.City)
	fmt.Println(address.State)
	fmt.Println(address.Zip)
	fmt.Println(address.Country)
	fmt.Println(address.Latitude)
	fmt.Println(address.Longitude)

	fmt.Println(contact.Phone)
	fmt.Println(contact.Email)

	fmt.Println(creditCard.Type)
	fmt.Println(creditCard.Number)
	fmt.Println(creditCard.Exp)
	fmt.Println(creditCard.Cvv)

	// Output: Markus
	// Moen
	// male
	// 420776036
	// http://lorempixel.com/300/300/people
	// Lockman and Sons
	// Developer
	// Global
	// Brand
	// 5748 Streamville, Rossieview, Hawaii 60370
	// 5748 Streamville
	// Rossieview
	// Hawaii
	// 60370
	// Burkina Faso
	// -6.662594491850811
	// 23.921575244414612
	// 5822502868
	// hassanheaney@tromp.net
	// MasterCard
	// 2720991544071613
	// 07/26
	// 448
}

func BenchmarkPerson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Global.Person()
	}
}
