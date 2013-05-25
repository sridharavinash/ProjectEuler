# If we list all the natural numbers below 10 at are multiples of 3 or 5, we
# get 3, 5, 6 and 9.  The sum of these multiples is 23.
#
# Find the sum of all the multiples of 3 or 5 below 1000.

def main(limit):
    limit = int(limit)
    sum = 0
    for number in xrange(limit):
        if (number % 3) == 0 or (number % 5 == 0):
            sum += number
    print "sum =" , sum

if __name__ == '__main__':
    limit = raw_input("Enter limit :")
    main(limit)
