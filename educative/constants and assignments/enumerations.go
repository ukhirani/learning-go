package main

func main() {
	// Listing of all elements of a set is called enumeration. Constants can be used for enumerations.

	const (
		UNKNOWN = 0
		FEMALE  = 1
		MALE    = 2
	)
	// Interestingly value iota can be used to enumerate the values.
	// this is the same as using iota

	const (
		UNKNOWNN = iota
		FEMALEE  = iota
		MALEE    = iota
	)

	// The first use of iota gives 0. Whenever iota is used again on a new line, its value is incremented by 1; so UNKNOWN gets 0, FEMALE gets 1 and MALE gets 2. Remember that a new const block or declaration initializes iota back to 0.

	// he above notation can be shortened, making no difference

	const (
		UNKNOWNNN = iota
		FEMALEEE
		MALEEE
	)

	// You can give enumeration a type name.
	// For example, FEMALE, MALE and UNKNOWN are categories of Gender.

	type Gender int
	const (
		UNKNOWNNNN Gender = iota
		FEMALEEEE
		MALEEEE
	)

}
