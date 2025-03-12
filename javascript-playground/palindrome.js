const isPalindrome = (word) => {
  const cleanedWord = word.toLowerCase().replace(/[^a-z0-9]/g, "")
  let startPointer = 0
  let endPointer = cleanedWord.length - 1
  while (startPointer <= endPointer) {
    if (cleanedWord[startPointer] != cleanedWord[endPointer]) {
      return false
    }
    startPointer++
    endPointer--
  }
  return true
}

module.exports = isPalindrome 
