""" A Python script that calculates the Extended Euclidean Algorithm """

# Function that computes the Extended Euclidean Algorithm
def extdEuclid(a, b):

    k=2 # Index

    err_str = None

    # Check if conditions are satisfied, then begin the algorith
    if not (a>=0 and b>0 and a>=b):
        err_str = "Condition are not satisfied... "
        return 0, 0, 0, err_str # if conditions are not satisfied

    # Initilize the (empty) arrays r, s and t of size k
    r = [None for _ in range(k)]
    s = [None for _ in range(k)]
    t = [None for _ in range(k)]

    # Insert the integers a and b
    r[0], r[1] = a, b
    s[0], s[1] = 1, 0
    t[0], t[1] = 0, 1
    # while b (in rk-1) is not 0
    while r[k-1] != 0:
        q = r[k-2]//r[k-1]
        r.append(int(r[k-2]%r[k-1])) # encapsulate results to int(), otherwise it will append float values
        s.append(int(s[k-2]-(q*s[k-1])))
        t.append(int(t[k-2]-(q*t[k-1])))
        k += 1
    return r[len(r)-2], s[len(s)-2], t[len(t)-2], err_str

def gcd(a,b):
    if b==0:
        return a
    return gcd(b, a%b)




