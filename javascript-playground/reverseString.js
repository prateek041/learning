const checkString = "hello"

const reverseStringSelfImplement = (word) => {
  const wordArray = word.split("")
  let reversedArray = []
  for (let index = wordArray.length - 1; index >= 0; index--) {
    reversedArray.push(wordArray[index])
  }
  return reversedArray.join("")
}

console.log(reverseStringSelfImplement(checkString))
