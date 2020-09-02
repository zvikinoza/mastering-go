package sharedRPC

// MFloats : exported rpc call interface
type MFloats struct {
	A1, A2 float64
}

// MInterface : exported rpc call interface
type MInterface interface {
	Multiply(args *MFloats, reply *float64) error
	Power(args *MFloats, reply *float64) error
}
