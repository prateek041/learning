const isAnagram = (wordOne, wordTwo) => {
  const wordOneMap = createWordMap(wordOne)
  const wordTwoMap = createWordMap(wordTwo)

  if (wordOneMap.size !== wordTwoMap.size) {
    return false
  }

  for (let [key, value] of wordOneMap) {
    if (wordTwoMap.get(key) !== value) {
      return false
    }
  }

  return true
}

const createWordMap = (word) => {
  word = word.toLowerCase().split("")
  let wordMap = new Map()
  word.forEach(item => {
    wordMap.set(item, (wordMap.get(item) || 0) + 1)
  })
  return wordMap
}

module.exports = isAnagram
