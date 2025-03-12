const isPalindrome = require("../palindrome")

const sampleInputs = [
  { input: "racecar", output: true },
  { input: "hello", output: false },
  { input: "A man a plan a canal: Panama", output: true }
]

describe("IsPalindrome", () => {
  test("Run with sample inputs", () => {
    sampleInputs.map(({ input, output }) => {
      expect(isPalindrome(input)).toBe(output)
    })
  })
})
