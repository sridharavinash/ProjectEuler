"""
The series, 1^1 + 2^2 + 3^3 + ... + 10^10 = 10405071317.

Find the last ten digits of the series, 1^1 + 2^2 + 3^3 + ... + 1000^1000.
"""

sumPowers = 0
for i in xrange(1,1001):
    print i
    sumPowers += pow(i,i)

print sumPowers
strNum = str(sumPowers)

print strNum[len(strNum)-10:]
