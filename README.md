# Gobro

It stands for "Go" and "Slowbro", which is a lovely pokemon -- but quite lazy. Since I am lazy too when it comes to making a budget, Gobro, I choose you!

## Installation

TODO

## Database

TODO

## Run the tests

```sh
go test ./...
```

## Usecase

### Create a budget and get the status

```sh
./gobro init 2000

./gobro status
Created on 2016-06-26 12:18:21.328 +0200 CEST
Initial balance 2000
Total earnings 0
Total expenses 0
Balance 2000 (0)
```

### Add an expense

```sh
./gobro add 100 "INSURANCE"

./gobro list
0 ) -100 INSURANCE 2016-06-26 12:20:43.139 +0200 CEST

./gobro status
Created on 2016-06-26 12:18:21.328 +0200 CEST
Initial balance 2000
Total earnings 0
Total expenses -100
Balance 1900 (-100)
```

### Add an earning

Use the "+" operator to add a positive value i.e. an earning.

```sh
./gobro add +25.17 "REFUND"

./gobro list
0 ) -100 INSURANCE 2016-06-26 12:20:43.139 +0200 CEST
1 ) 25.17 REFUND 2016-06-26 12:37:43.468 +0200 CEST

./gobro status
Created on 2016-06-26 12:18:21.328 +0200 CEST
Initial balance 2000
Total earnings 25.17
Total expenses -100
Balance 1925.17 (-74.83)
```

### Close the budget
``` sh
./gobro close

./gobro status
2016/06/26 12:52:47 There is not any active budget.
exit status 1

```

### Fixed expenses

Fixed expenses come every month. They are set by default when created a new budget.

``` sh
./gobro list -t fixed

./gobro add -t fixed 9.99 "SPOTIFY"

./gobro list -t fixed
0 ) -9.99 SPOTIFY 2016-06-26 12:55:21.427 +0200 CEST

./gobro init 2000
./gobro list
0 ) -9.99 SPOTIFY 2016-06-26 12:59:18.219 +0200 CEST

./gobro status
Created on 2016-06-26 12:18:21.328 +0200 CEST
Initial balance 2000
Total earnings 0
Total expenses -9.99
Balance 1990.01 (-9.99)
```

### Pristine budget

The -p or --pristine flag inits a simple budget without any fixed expense.

``` sh
./gobro list -t fixed
0 ) -9.99 SPOTIFY 2016-06-26 12:59:18.219 +0200 CEST

./gobro init -p 2000

./gobro status
Created on 2016-06-26 12:18:21.328 +0200 CEST
Initial balance 2000
Total earnings 0
Total expenses 0
Balance 2000 (0)
```

### Manually check the expenses

It's better to check frequently your bank account, in order to see if it
matches what is expected in the budget.

``` sh
./gobro check
```

For each expense, Gobro will ask if it has been traced on your bank account. You have to enter "y"
if you've got the trace. If not, then just press enter.

Running the status command gives the sum of the expenses that have not been checked yet.

### Remove / Update / Fix
TODO
