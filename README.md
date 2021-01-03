# calchaSolve

This is my complete solution for the QuantBet Developer Challenge at
https://quantbet.com/quiz/dev. This program is able to scrape the inputs for the challenge as well as submit them using
web requests.

For context, a calcha has been defined by QuantBet to be anti-human measures on the web. This is the opposite of a
captcha which is designed to be anti-bot. To complete this challenge, the solution request must be POSTed very quickly
since it is designed to be too complex for the majority of humans.

This is my implementation of such a program.

### Installation

- To build the binary, run `go build -i ./cmd/calchasolve`
- To learn about the args run `./calchasolve -h`
- Then run the binary with `./calchasolve` with the desired flags