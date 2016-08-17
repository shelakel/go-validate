package validate

//go:generate validate -type=bool -output=%s_generated.go

//go:generate validate -type=int -output=%s_generated.go
//go:generate validate -type=int8 -output=%s_generated.go
//go:generate validate -type=int16 -output=%s_generated.go
//go:generate validate -type=int32 -output=%s_generated.go
//go:generate validate -type=int64 -output=%s_generated.go
//go:generate validate -type=byte -output=%s_generated.go
//go:generate validate -type=uint -output=%s_generated.go
//go:generate validate -type=uint8 -output=%s_generated.go
//go:generate validate -type=uint16 -output=%s_generated.go
//go:generate validate -type=uint32 -output=%s_generated.go
//go:generate validate -type=uint64 -output=%s_generated.go

//go:generate validate -type=float32 -output=%s_generated.go
//go:generate validate -type=float64 -output=%s_generated.go
//go:generate validate -type=complex64 -output=%s_generated.go
//go:generate validate -type=complex128 -output=%s_generated.go

//go:generate validate -type=string -output=%s_generated.go
//go:generate validate -type=time.Time -name=Time -imports=time -output=%s_generated.go
