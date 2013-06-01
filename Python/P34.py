"""
145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.

Find the sum of all numbers which are equal to the sum of the factorial of their digits.

Note: as 1! = 1 and 2! = 2 are not sums they are not included.
"""

from utils.fact import fact

def main():
    sumNum = 0
    numOfDigits = 7
    upperBound = fact(9) * numOfDigits
    print 'upperBound Guess  9! x %d = %d ' %(numOfDigits,upperBound)
    for num in xrange(3,upperBound):
        sumDigitsFact = 0;
        for d in str(num):
            sumDigitsFact +=fact(int(d))
        if(sumDigitsFact == num):
            sumNum += num
            print num
    print sumNum

if __name__ == '__main__':
    main()
