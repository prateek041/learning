# Lexing

## Why use Lexing

To convert the code in a format that makes it easily accesible

## What is Lexing

Lexing is the process of converting the source code into tokens. It is done by a lexer (also called a tokenizer)

so, if we consider this statement as an example
`let x = 5 + 5;`, lexer will convert it into an array of tokens that looks like this

```
[
LET,
IDENTIFIER("x"),
EQUAL_SIGN,
INTEGER(5),
PLUS_SIGN,
INTEGER(5),
SEMI_COLON
]
```

# Language specific things

- Identifiers: These are variable names, that are binded to an expression, or a variable.
- Keywords: let, fn.
- Numbers: Integers.
- Special Characters:

## Lexer Definition

There are a few things to be considered when creating the Lexer struct.

We want a way to go through a given string, character by character and take out the words we care about. The words that part of our language spec. It is just like the problem of finding a matching word from a string.

Therefore the struct turns out to be looking like this.

```go
type Lexer struct{
  ch byte
  readPosition int
  Position int
  input string
}
```

### Working of Lexer

We have to create methods that read the string character by character and then another method that checks the string and returns a corresponding token

As mentioned earlier, we have identifiers, numbers, special characters and keywords. Let's see one by one how they will work.

#### special Characters

Since special characters are single, we can be sure of which one they are in just one iteration of `readChar` which updates the `readPosition` by one and update the `ch` property with the current character we are dealing with.

#### Keywords

Another way of seeing it is, a token can be either a special character, a number or a string. If it is a string, it can be an identifier, which has limitless possiblities since people can name their variables whatever they want.

These are comparatively tricky, reading the first character does not give a definitive answer, we have to read characters till it starts making sense. For examplee `let` is a keyword, but `letter` can be an identifier, and we won't know the difference until we read the entire thing.

So, we want a case that reads the entire word, and checks if it is one of the known keywords (since they are limited), if it is not, it is an identifier.

Also, take care about skipping white space when necessary.

#### Numbers

These are just like characters, only difference is that the function getting you the identifier is returning a combination of digits instead of combinations of characters.

Finally, add the remaining special characters (operators) and keywords like true, false etc. Here we have three categories, single characters like `>`, `<`, `=` etc. double characters like `==`, `!=` etc. and lastly, the remaining keywords. There is nothing special here though, we already know how to deal with single characters and keywords, for double characters we create more if statements to check what the `next` character to something like `=` and decide on the basis of that what the token is.

Finally, implement a REPL
