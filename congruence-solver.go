// Go script that solves congruence systems

package main

import (
	"fmt"
)

// Function that computes the Extended Euclidean Algorithm (for the modular inverse computation)
func ExtdEuclid(a, b int) (r, s, t int, err error) {

	var (
		k int // For the index
		q int // Temp storage for q
	)

	k = 2 //index

	if (a >= 0) && (b > 0) && (a >= b) {

		// Define the arrays/slices  r, s, t with length k
		r := make([]int, k)
		s := make([]int, k)
		t := make([]int, k)

		// Insert both a and b into the array `r'`
		r[0] = a
		r[1] = b
		s[0] = 1
		s[1] = 0
		t[0] = 0
		t[1] = 1

		// while rk-1 (which is b) is not 0
		// Go lang does not have while-loops
		// for loops are used instead
		//index := k-1

		for r[k-1] != 0 {
			q = r[k-2] / r[k-1]
			r = append(r, r[k-2]%r[k-1])
			s = append(s, s[k-2]-(q*s[k-1]))
			t = append(t, t[k-2]-(q*t[k-1]))
			k += 1
		}

		return r[len(r)-2], s[len(s)-2], t[len(t)-2], nil

	} else {

		return 0, 0, 0, fmt.Errorf("Conditions Are Not Satisfied...")

	}
}

// Fucntion used to check for coprime integers (when their gcd = 1)
func Gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

//--------------------------------------------------------------------------------------------------

// Fucntion to check if the moduli vector [n0,...,nk-1] are pairwise coprime
// It checks if each pair of numbers in the vector are coprime
// The greatest common divisor of each pair should be 1
func pairwiseCoprime(v []int) bool {
	n := len(v)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if Gcd(v[i], v[j]) != 1 { // Call the function Gcd() to check for the greatest common divisor
				return false // in case gcd(n0,n1) != 1, they are not coprime
			}
		}
	}

	return true // return true if the pairs are coprime
}

// Function to handle user input
func inputHandler() (a, n []int, r int) {

	var size int

	for true {

		// Ask the user how many values should the array consist of? (i.e, how many integers are in the a and b vector?)
		// Both vectors must be the same size

		fmt.Println("How many values do you want to enter for the remainder and  moduli vectors 'a' and 'n' (size of vectors)?")
		_, err0 := fmt.Scanln(&size)
		if err0 == nil {
			break
		}
		fmt.Println("Invalid size. The size must be an integer...")
		var dump0 string
		fmt.Scanln(&dump0)
	}

	// Create the vector(s) using the specified size
	// Both vectors are initialized to the same size, since these will be the input for-
	// -a sytem of congruent equations

	a = make([]int, size)
	n = make([]int, size)

	fmt.Println("Enter the values for the remainder vector a ")

	// Enforce only integer values input vactor a
	for true {
		for i := 0; i < size; {

			fmt.Printf("Enter element a[%d] of a: ", i)
			_, err1 := fmt.Scan(&a[i])

			// If the user enters characters
			if err1 != nil {
				fmt.Println("Invalid input. Enter a strictly integer value.")
				var dump1 string
				fmt.Scanln(&dump1) // Clear STDIN buffer
				continue
			}
			i++ // increment when the input value is valid (is an integer)
		}
		break //break once the values are entered, otherwise the program will loop
	}

	fmt.Println("Enter the values for moduli vector n ")

	// Enforce only integer values input vactor n
	for true {
		for i := 0; i < size; {

			fmt.Printf("Enter element n[%d] of n: ", i)
			_, err1 := fmt.Scan(&n[i])

			// If the user enters characters
			if err1 != nil {
				fmt.Println("Invalid input. Enter a strictly integer value.")
				var dump1 string
				fmt.Scanln(&dump1) // Clear STDIN buffer
				continue
			}
			i++ // increment when the input value is valid (is an integer)
		}
		break //break once the values are entered, otherwise the program will loop
	}

	//---------- Range of the solution set input -----------

	fmt.Println("Enter a range for the solution set, or 0 to display the full set.")

	// allow only integer value input
	for true {
		fmt.Print("Range: ")

		_, err := fmt.Scan(&r)
		if err == nil {
			break
		}
		fmt.Println("Invalid input. Enter a strictly integer value.")
		var dump string
		fmt.Scanln(&dump) // Clear STDIN buffer
	}

	return a, n, r
}

// Function that computes system of congruence
func congruentSystemSolver(a, n []int) (Nn, S, T, setSolution []int, x, Xx, N int, err error) {

	// Check n for pairwise coprime by passing the vector/list to the function
	if pairwiseCoprime(n) {
		// To store the product moduli N = n0 * n1 * ... *nk-1
		// Multiply the values in the n array
		N = 1
		for y := 0; y < len(n); y++ {
			N *= n[y] // N = N*n[0] -> N = N*n[1] -> N = N*n[2] ...
		}

		Nn = make([]int, 0)
		S = make([]int, 0)
		T = make([]int, 0) // to store the result when the extended Euclidean function is called

		j := 0
		// for j<k
		for j < len(n) {
			Nn = append(Nn, N/n[j])
			_, Ss, Tt, err1 := ExtdEuclid(Nn[j], n[j]) // Compute the modular inverse  ▷ 1 = s j · Nj + t j · n j
			if err1 != nil {                           //The ExtdEuclid (this specific implementation) returns an error if the conditions are not met
				fmt.Println(err1) // If the conditions are not met, inform the user
				return
			}
			// If the conditions are met (in the ExTEuclid Algo.) then continue the CRT algorithm
			S = append(S, Ss)
			T = append(T, Tt)
			j++
		}

		Xx = 0
		for i := 0; i < len(n); i++ { // Computing the solution
			Xx += a[i] * S[i] * Nn[i] // will give soltion between [0, N-1]
		}

		// x is the unique solution
		x := Xx % N //resul mod N (the product moduli)

		return Nn, S, T, setSolution, x, Xx, N, nil

	} else { // If is not coprime, return the error message
		return nil, nil, nil, nil, 0, 0, 0, fmt.Errorf("The vector n = [n0,...,nk] is not coprime..")
	}

}

func main() {

	// Call the function that returns the vectors a and n enetered by the user
	// Both vectors are the same size
	a, n, r := inputHandler()

	// Call the solver and feed it the input
	Nn, S, T, setSolution, x, Xx, N, err := congruentSystemSolver(a, n)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("\n")

	fmt.Println("Result vectors: ")
	fmt.Println(Nn)
	fmt.Println(S)
	fmt.Println(T)

	fmt.Print("\n")

	fmt.Println("Equation form: ")
	//▷ 1 = s j · Nj + t j · n j

	var res int
	for k := 0; k < len(Nn); k++ { // loop to write the returned result into an equation
		res = S[k]*Nn[k] + T[k]*n[k]
		fmt.Printf("%d = %d * %d + %d * %d", res, S[k], Nn[k], T[k], n[k])
		fmt.Print("\n")
	}

	fmt.Println("\n")
	fmt.Printf("Unique solution is %d modulo to the product moduli %d", x, N)

	fmt.Println("\n")

	fmt.Printf("One solution from the set is %d since %d mod %d = %d ", Xx, Xx, N, x)
	fmt.Println("\n")

	// Check if the range is 0, display the full set
	if r == 0 {
		r = N // the range of the product of moduli N
	}

	// Choose the range of the colution set
	// There could be infinitely many solutions, so it's best to rturn the solution set within a range

	fmt.Printf("Set of all solutions, {%d + m * %d | 'm' is + or - integer}, to this congruence system within the range %d: ", Xx, N, r) // {x + m · N | m ∈ Z}
	fmt.Println("\n")

	// Choose m values that cover [0, N-1] for the full solution set -- This covers only the positive solutions
	// To include the netagive solutions, the range should be from m=-N+1; m <= N
	for m := -r + 1; m <= r; m++ {
		setSolution = append(setSolution, x+m*N)
	}

	fmt.Println(setSolution)

}
