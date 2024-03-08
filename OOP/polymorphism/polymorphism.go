package main

import "fmt"

// the goal of this program is to ca;culate the net income of an organization
// the net income is the total amount they made from diferent projects
// there are two types of projects, fixedbilling, and timeandmaterial

// define an interface for the income
type Income interface{
	calculate() int
	source() string
}

// the income type structs that will have mmethods to implement the incomme interface
type FixedBilling struct {
	projectName string
	biddedAmount int
}

type timeAndMaterial struct {
	projectName string
	noOfHours int
	hourlyRate int
}

// new income stream ads :)
type Advertisment struct {
	adName string
	costPerClick int
	noOfClicks int
}


// the  methods for each income type
func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm timeAndMaterial) calculate() int {
	return tm.hourlyRate * tm.noOfHours
}

func (tm timeAndMaterial) source() string {
	return tm.projectName
}

// methods for new income stream
func (ad Advertisment) calculate() int {
	return ad.noOfClicks * ad.costPerClick
}

func (ad Advertisment) source() string {
	return ad.adName
}

// next we will define the calculateNetIncome function
// we will pass a slice of incomes to the function
// this function will use a for-range loop to go through each income slice passed and use the calculate method to find the income for 'one' project.
// the netincome variable we created will then increment, adding each new income that is looped through.
// Depending on the concrete type of the Income interface, different calculate() and source() methods will be called. Thus we have achieved polymorphism in the calculateNetIncome function.
func calculateNetIncome(inc []Income) {
	var netincome int = 0
	for _, income := range inc {
		fmt.Printf("Income from %s = $%d \n", income.source(), income.calculate())
		netincome += income.calculate() // this line simply increments the netincome by adding all incomes in the slice to it.
	}
	fmt.Printf("net Income for the organization = $%d \n", netincome)
}


// in the mmain function we'll use this calculateNetIncome
func main() {
	project1 := FixedBilling{projectName: "project 1", biddedAmount: 3000}
	project2 := FixedBilling{projectName: "project 2", biddedAmount: 10000}
	project3 := timeAndMaterial{projectName: "project 3", noOfHours: 120, hourlyRate: 70}
	bannerAd := Advertisment{adName: "banner ad", costPerClick: 5, noOfClicks: 750}
	popupAd := Advertisment{adName: "popup ad", costPerClick: 3, noOfClicks: 900}

	// pass in these projects to a slice named streamsOfIncome, then call the calculateNetInncome function on that slice
	streamsOfIncome := []Income{project1, project2, project3, bannerAd, popupAd}
	calculateNetIncome(streamsOfIncome)
}

// Adding a new income stream
// Let’s say the organization has found a new income stream through advertisements. It is very simple to add this new income stream and calculate the total income without making any changes to the calculateNetIncome function. This becomes possible because of polymorphism.

// refer to line 27 to see the implementation

// You would have noticed that we did not make any changes to the calculateNetIncome function though we added a new income stream. It just worked because of polymorphism. Since the new Advertisement type also implemented the Income interface, we were able to add it to the incomeStreams slice. The calculateNetIncome function also worked without any changes as it was able to call the calculate() and source() methods of the Advertisement type.