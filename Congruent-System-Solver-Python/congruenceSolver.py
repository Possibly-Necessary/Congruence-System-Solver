
""" Congruent System Solver in Python """

# Import functions from the ExtendedEuclidean script
from ExtendedEuclidean import *
from inputHandler import *
import sys

# Function that checks the moduli vector n = [n0,..,nk] for pairwise co-prime
# The greatest common divisor of each pair of the values is 1
def pairwiseCoprime(v):
    n = len(v)
    for i in range(n):
        for j in range(i+1, n):
            if gcd(v[i], v[j]) != 1:
                return False
    return True


# Function that computes system of congruence
def congruentSystemSolver(a, n):

    # Return an empty string if everything pairwise coprime of the moduli vector checks out
    err = None

    # Check the moduli array n if the values within are pairwise co-prime
    if pairwiseCoprime(n):
        N = 1 # To store the product moduli
        for y in range(len(n)):
            N *= n[y]

        # Initialize empty arrays to store the results [None]*len(n)
        Nn = []
        S = []
        T = []

        j = 0
        while j < len(n):
            Nn.append(N//n[j])
            _, Ss, Tt, strr = extdEuclid(Nn[j], n[j])
            if strr is not None: # Check if the string is not empty, that means the conditions for the
                # extdEuclid() function were not met.
                sys.exit(str)
            # append the values
            S.append(Ss)
            T.append(Tt)
            j +=1
        # Computing the soluiton
        Xx = 0
        for s in range(len(n)):
            Xx += a[s] * S[s] * Nn[s]

        # x is the unique solution
        x = Xx % N

        return Nn, S, T, x, Xx, N, err
    else:
        err = "The values in the moduli vector n = [n0,..,nk] are not pairwise coprime "
        return None, None, None, 0, 0, 0, err


# ----------- 'main' ------------------------------- #

# Call the input handler function and get the input(s)
a, n, r = userInpt()

# Call the congruent solver
Nn, S, T, x, Xx, N, err= congruentSystemSolver(a, n)

if err is not None:
    print(err)

else:
    print('\n')
    print(f'Result vectors:')
    print(Nn)
    print(S)
    print(T,"\n")

    # ▷ 1 = s j · Nj + t j · n j
    print(f'Equation form: ')

    for i in range(len(Nn)):
        res = S[i]*Nn[i] + T[i]*n[i]
        print(f'{res} = {S[i]}*{Nn[i]} + {T[i]}*{n[i]}')

    print('\n')

    print(f'Unique solution is {x} modulo to the product moduli {N}\n')

    print(f'One solution is {Xx} since {Xx} mod {N} = {x}\n')
    # check if the range is 0 to display the full set in that case
    if r == 0:
        r = N

    print(f'The set of all solutions, {Xx}+m*{N} |\'m\' is + or - integer, to this congruence system within the range {r}:\n')

    setSolution = []
    for m in range(-r+1, r):
        setSolution.append(x+m*N)

    print(setSolution)


