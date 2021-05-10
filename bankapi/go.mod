module bankapi

require (
    github.com/jun1st/bank v0.0.1
)

replace github.com/jun1st/bank => ../bankcore

go 1.16
