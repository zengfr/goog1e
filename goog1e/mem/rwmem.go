package mem

import (
	"fmt"
	"os"

	"github.com/avelex/procmem"
)

func RWMem() {
	if len(os.Args) >= 0 {
		return
	}
	pid := 164433
	address := 0xc00012c008
	proc, _ := procmem.NewProcess(pid)

	want := "Hello"

	got, err := proc.ReadString(uintptr(address), uint64(len(want)))
	if err != nil {
		// ...
	}

	fmt.Println(want == got) // true

	newString := "World"

	ptr := address

	if err := proc.WriteString(uintptr(ptr), newString); err != nil {
		// ...
	}

	got, err = proc.ReadString(uintptr(address), uint64(len(newString)))
	if err != nil {
		// ...
	}

	fmt.Println(newString == got) // true
}
