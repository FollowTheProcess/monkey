# Monkey

The monkey programming language from the [Writing an Interpreter in Go] and [Writing a Compiler in Go] books, with some small experimentation thrown in!

## Tweaks

Whilst doing this I've made some tweaks to learn a bit more:

- Natively support full UTF-8 rather than simply ASCII as per the book version
- Add an `Is` method to `lexer.Token` rather than lots of helpers
- Store all the `Tokens` in the lexer package rather than a token package

[Writing an Interpreter in Go]: https://interpreterbook.com
[Writing a Compiler in Go]: https://compilerbook.com
