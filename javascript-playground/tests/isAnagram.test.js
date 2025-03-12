const isAnagram = require("../isAnagram")

const sampleTest = [
  {
    input: ["listen", "silent"], output: true
  },
  {
    input: ["Hello", "Olelh"], output: true
  },
  {
    input: ["Rat", "Car"], output: false
  },
  {
    input: ["Dormitory", "Dirty room"], output: false
  }
]

describe("isAnagram", () => {
  test("Testing with Sample Tests", () => {
    sampleTest.map(({ input: [wordOne, wordTwo], output }) => {
      expect(isAnagram(wordOne, wordTwo)).toBe(output)
    })
  })
})
