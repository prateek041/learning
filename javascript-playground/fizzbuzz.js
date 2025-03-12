const fizzBuzz = (number) => {
  let startIndex = 1
  const output = []
  while (startIndex <= number) {
    if (multipleOfFiveAndThree(startIndex)) {
      output.push("FizzBuzz")
    } else if (multipleOfThree(startIndex)) {
      output.push("Fizz")
    } else if (multipleOfFive(startIndex)) {
      output.push("Buzz")
    } else {
      output.push(startIndex)
    }
    startIndex++
  }
  return output
}

const multipleOfFiveAndThree = (number) => {
  return multipleOfThree(number) && multipleOfFive(number)
}

const multipleOfFive = (number) => {

  return number % 5 === 0
}

const multipleOfThree = (number) => {
  return number % 3 === 0
}

module.exports = fizzBuzz
