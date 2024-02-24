package computer

type Spec struct{
	// we want to export the fields, to do this we make the first letter of the field uppercase
	Maker string
	Price int
	model string  // not exported
}