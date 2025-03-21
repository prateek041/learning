const fizzBuzz = require("../fizzbuzz")

const testSample =
{
  input: 15, output: [
    1, 2, "Fizz", 4, "Buzz", "Fizz", 7, 8, "Fizz", "Buzz", 11, "Fizz", 13, 14, "FizzBuzz"
  ]
}

describe("FizzBuzz", () => {
  test("Testing FizzBuzz on sample inputs", () => {
    expect(fizzBuzz(testSample.input)).toEqual(testSample.output)
  })
})
