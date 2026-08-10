package computer 

type Spec struct {//exported struct
	Maker string //exported field
	Price int //exported field
	model string //unexported field
}

// The above snippet creates a package computer which contains an exported struct type Spec with two exported fields Maker and Price and one unexported field model. Let’s import this package from the main package and use the Spec struct.