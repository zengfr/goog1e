package asm

import (
	"fmt"

	"github.com/keystone-engine/keystone/bindings/go/keystone"
)

func AsmToX86_32(assembly string) (bs []byte, err error) {

	insn, err := AsmToByte(assembly, keystone.ARCH_X86, keystone.MODE_32)
	return insn, err
}
func AsmToX86_64(assembly string) (bs []byte, err error) {

	insn, err := AsmToByte(assembly, keystone.ARCH_X86, keystone.MODE_64)
	return insn, err
}
func AsmToArm_64(assembly string) (bs []byte, err error) {

	insn, err := AsmToByte(assembly, keystone.ARCH_ARM64, keystone.MODE_64)
	return insn, err
}
func AsmToByte(assembly string, a keystone.Architecture, m keystone.Mode) (bs []byte, err error) {
	ks, err := keystone.New(a, m)
	if err != nil {
		panic(err)
	}
	defer ks.Close()

	return _AsmToByte(assembly, ks)
}
func _AsmToByte(assembly string, ks *keystone.Keystone) (bs []byte, err error) {

	if err := ks.Option(keystone.OPT_SYNTAX, keystone.OPT_SYNTAX_INTEL); err != nil {
		panic(fmt.Errorf("Could not set syntax option to intel"))
	}
	insn, _, ok := ks.Assemble(assembly, 0)
	if !ok {
		panic(fmt.Errorf("Could not assemble instruction"))
	} else {
		fmt.Printf("%s: [%x]", assembly, insn)
	}
	return insn, err
}

func Test() {
	assembly := "add eax, ebx"
	AsmToX86_32(assembly)
}
