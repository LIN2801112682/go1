module Hello

go 1.17

require rsc.io/quote v1.5.2

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)

require (
	Utils v0.0.0
	github.com/bits-and-blooms/bitset v1.2.1
)

replace Utils v0.0.0 => ./Utils
