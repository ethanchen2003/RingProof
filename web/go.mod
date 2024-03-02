module RingProof

go 1.22.0

replace github.com/phoneapp/service => ../service

replace github.com/phoneapp/security => ../security

require (
	github.com/phoneapp/security v0.0.0-00010101000000-000000000000
	github.com/phoneapp/service v0.0.0-00010101000000-000000000000
)

require github.com/google/uuid v1.6.0 // indirect
