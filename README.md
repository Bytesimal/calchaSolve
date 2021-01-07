# calchaSolve

This is my complete solution for the QuantBet Developer Challenge at https://quantbet.com/quiz/dev. This program is able
to scrape the inputs for the challenge as well as submit them using web requests.

For context, a calcha has been defined by QuantBet to be anti-human measures on the web. This is the opposite of a
captcha which is designed to be anti-bot. To complete this challenge, the solution request must be POSTed very quickly
since it is designed to be too complex for the majority of humans.

This is my implementation of such a program. It has been developed and tested on `go1.15.6` but it should be compatible
with most past and future versions.

### Installation

- To build the binary, run `go build -i ./cmd/calchasolve`
- To learn about the args run `./calchasolve -h`
- Then run the binary with `./calchasolve` with the desired flags

It is possible when specifing the `-l` flag to use time prefixes such as s, ms, us/µs and ns

Typical run examples:

- `./calchaSolve -l 1s -r -s`
- `./calchaSolve -l 50µs -r`
- `./calchaSolve -l 4m -s`
- `./calchaSolve`