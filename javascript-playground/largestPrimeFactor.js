const largestPrimeFactor = (number) => {
  maxPrime = -1
  while (number % 2 === 0) {
    maxPrime = 2
    number /= 2
  }

  for (let num = 3; num * num <= number; num += 2) {
    while (number % num === 0) {
      maxPrime = num;
      number /= num
    }
  }

  if (number > 2) {
    maxPrime = number
  }

  return maxPrime
}

module.exports = largestPrimeFactor
