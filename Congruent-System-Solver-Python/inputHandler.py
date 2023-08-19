
# Script to handel user input

# Function to handle user input
def userInpt():

    print("Enter the size of your moduli and remainder vectors 'a' and 'n'.")

    # Size (number of congruent equations)
    while True:
        try:
            k = int(float(input("Size: ")))
            break
        except ValueError:
            print("Invalid input. Size should be a strictly integer value.")

    a = []
    n = []

    # Enforce only integer inputs for a and n
    # In case float values are entered, only the integer part will be taken
    print("Enter the values for the remainder vector 'a'.")
    while True:
        try:
            while len(a) < k:
                a1 = int(float(input("a: ")))
                a.append(a1)
            break
        except ValueError:
            print("Invalid input. Enter a strictly integer value.")

    print("Enter the vaalues for the moduli vector 'n'.")

    while True:
        try:
            while len(n) < k:
                n1 = int(float(input("n: ")))
                n.append(n1)
            break
        except ValueError:
            print("Invalid input. Enter a strictly integer value.")


    print(f'Enter a range for the solution set, or enter 0 to display the full set of solutions.')
    while True:
        try:
            r = int(float(input("Range: ")))
            break
        except ValueError:
            print("Invalid range. Enter a strictly integer value.")

    return a, n, r




