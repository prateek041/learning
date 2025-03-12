const largestPrimeFactor = require("../largestPrimeFactor")

const sampleTest = [
  { input: 15, output: 5 },
  { input: 100, output: 5 },
  { input: 29, output: 29 }
]

describe("Largest Prime Factor", () => {
  test("Finding largest Prime factor for test input", () => {
    sampleTest.map(item => {
      expect(largestPrimeFactor(item.input)).toBe(item.output)
    })
  })
})
