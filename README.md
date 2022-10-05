# Return Change Machine #

A classic example of the machine that returns change optimizing the amount of elements it gives for a given amount

## Contents

1. [Using go](#using-go)
1. [Problem](#problem)
1. [Testing](#testing)
1. [Credits](#credits)
1. [TODO](#todo)

## Using go

I'm a beginner in Go, so I thought of using this example as practice. I'll be trying to apply some minor optimizations/refactor than what I usually did

## Problem

You'll be given a numeric input of the amount that needs the return from.\
You should output how much of each element (5x1 dolar bills, 10x50 dolar bills, 2x1 cents, etc.)

The least return the better solution, the less code and more readable, the finer.

## Testing

Given it's a small example with predefined best solutions, it'd be nice to clearly identify the possibilities.

Test Driven Development is a nice addition to this clear exercise.

## Credits

To my teacher (Pilar) from whom I learned this exercise.
And to Chelo, a teacher that mentioned this as a technical challenge for some bussiness

## TODO

- [X] Complete the base exercise
- [X] Use a map to compute the operations
- [X] Implement an amount for the change
  - [X] Implement a smart amount for the change, meaning, if we can supply change by using lower value currency, that's still considered a good solution
- [ ] Implement a (clearly defined) test set