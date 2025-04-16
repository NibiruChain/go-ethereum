package vm

import "github.com/ethereum/go-ethereum/common"

func (ecrecover) Address() common.Address {
	return common.BytesToAddress([]byte{1})
}

func (sha256hash) Address() common.Address {
	return common.BytesToAddress([]byte{2})
}

func (ripemd160hash) Address() common.Address {
	return common.BytesToAddress([]byte{3})
}

func (dataCopy) Address() common.Address {
	return common.BytesToAddress([]byte{4})
}

func (bigModExp) Address() common.Address {
	return common.BytesToAddress([]byte{5})
}

func (bn256AddIstanbul) Address() common.Address {
	return common.BytesToAddress([]byte{6})
}

func (bn256AddByzantium) Address() common.Address {
	return common.BytesToAddress([]byte{6})
}

func (bn256ScalarMulIstanbul) Address() common.Address {
	return common.BytesToAddress([]byte{7})
}

func (bn256ScalarMulByzantium) Address() common.Address {
	return common.BytesToAddress([]byte{7})
}

func (bn256PairingIstanbul) Address() common.Address {
	return common.BytesToAddress([]byte{8})
}

func (bn256PairingByzantium) Address() common.Address {
	return common.BytesToAddress([]byte{8})
}

func (blake2F) Address() common.Address {
	return common.BytesToAddress([]byte{9})
}

func (bls12381G1Add) Address() common.Address {
	return common.BytesToAddress([]byte{10})
}

func (bls12381G1Mul) Address() common.Address {
	return common.BytesToAddress([]byte{11})
}

func (bls12381G1MultiExp) Address() common.Address {
	return common.BytesToAddress([]byte{12})
}

func (bls12381G2Add) Address() common.Address {
	return common.BytesToAddress([]byte{13})
}
func (bls12381G2Mul) Address() common.Address {
	return common.BytesToAddress([]byte{14})
}

func (bls12381G2MultiExp) Address() common.Address {
	return common.BytesToAddress([]byte{15})
}

func (bls12381Pairing) Address() common.Address {
	return common.BytesToAddress([]byte{16})
}

func (bls12381MapG1) Address() common.Address {
	return common.BytesToAddress([]byte{17})
}

func (bls12381MapG2) Address() common.Address {
	return common.BytesToAddress([]byte{18})
}

func (kzgPointEvaluation) Address() common.Address {
	return common.BytesToAddress([]byte{0x0a})
}
