# Congruence-System-Solver

Pseudocode & example from the book: The MoonMath Manual to zk-SNARKS, p. 12-20.

Book: https://leastauthority.com/community-matters/moonmath-manual/

Algorithm: Chinese Remainder Theorem (CRT)

	Require: , k ∈ Z, j ∈ N0 and n0, . . . , nk−1 ∈ N coprime
		procedure CONGRUENCE-SYSTEMS-SOLVER(a0, . . . , ak−1)
			N ← n0 · . . . · nk−1
			while j < k do
				Nj ← N/n j
				(_, s j,t j) ← EXT − EUCLID(Nj, n j)           ▷ 1 = s j · Nj + t j · n j
			end while
			x′ ← ∑k−1
				j=0 a j · s j · Nj (part of the summation above)
			x ← x′ mod N
			return {x + m · N | m ∈ Z}
		end procedure
	Ensure: {x + m · N | m ∈ Z}

This algorithm solves the congruence system in the form:

		x ≡ a1 ( mod n1 )
		x ≡ a2 ( mod n2 )
			· · ·
		x ≡ ak ( mod nk )

whose all possible solutions are congruent modulo to the product N=n1*n2....*nk. Example of a system of congruences:

		x ≡ 4 ( mod 7 )
		x ≡ 1 ( mod 3 )
		x ≡ 3 ( mod 5 )
		x ≡ 0 ( mod 11 )

In this example, there are 4 congruence systems, so beforehand, you'll be asked to specify the size of the vectors a and n so that their sizes are initialized accordingly. In other words, how many equations (congruence systems) do you have? In our example, 4 would be the size.

		Output:

		How many values do you want to enter for the remainder and moduli vectors 'a' and 'n' (size of vectors)?
	  	input: 4

From the systems, the remainder vector a = [4 1 3 0] and the moduli vector n = [7 3 5 11] will be the input for this congruence system solver, which should be individually entered.

		Output:

		Enter the values for the remainder vector a
		Enter element a[0] of a: 4
		Enter element a[1] of a: 1
		Enter element a[2] of a: 3
		Enter element a[3] of a: 0
		Enter the values for moduli vector n
		Enter element n[0] of n: 7
		Enter element n[1] of n: 3
		Enter element n[2] of n: 5
		Enter element n[3] of n: 11

You'll also be asked to specify the range of the congruent solutions to be displayed. In case you want the full set, enter 0 as your range.

		Output:

		Enter a range for the solution set, or 0 to display the full set.
		Range: 3

Results:

		Output:

		Result vectors:
		[165 385 231 105]
		[2 1 1 2]
		[-47 -128 -46 -19]

		Equation form:
		1 = 2 * 165 + -47 * 7
		1 = 1 * 385 + -128 * 3
		1 = 1 * 231 + -46 * 5
		1 = 2 * 105 + -19 * 11

		Unique solution is 88 modulo to the product moduli 1155
		One solution from the set is 2398 since 2398 mod 1155 = 88

Since we specified the range to be 3 in this example, it will output three solutions from the set (both positive and negative integers).

		Output:

		Set of all solutions, {2398 + m * 1155 | 'm' is + or - integer}, to this congruence system within the range 3:
		[-2222 -1067 88 1243 2398 3553]

An example where the range is chosen to be 10:

		Output:

		Set of all solutions, {2398 + m * 1155 | 'm' is + or - integer}, to this congruence system within the range 10:
		[-10307 -9152 -7997 -6842 -5687 -4532 -3377 -2222 -1067 88 1243 2398 3553 4708 5863 7018 8173 9328 10483 11638]






