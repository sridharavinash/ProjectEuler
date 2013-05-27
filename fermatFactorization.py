"""
http://en.wikipedia.org/wiki/Fermat%27s_factorization_method
"""
import math

def isSquare(n):
    sqrt = math.sqrt(n)
    return sqrt - int(sqrt) == 0

def fermatFactor(number):
    number = int(number)
    if(number %2 == 0):
        print "Fermat Factorization only works on odd numbers"
        return
    a = math.ceil(math.sqrt(number))
 
    b = (a * a) - number 
    while not isSquare(b):
        a += 1
        b = (a * a) - number 
    return a - math.sqrt(b), a + math.sqrt(b)
    

if __name__ == '__main__':
    import time
    startTime = time.clock()
    print fermatFactor(34333333)
    print time.clock() - startTime, "seconds"
