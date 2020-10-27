build: grog-run grog-asm

clean:
	rm grog

grog-vm:
	go build ./vm

grog-run: grog-vm
	go build -o grog-run ./run

grog-parser:
	go install ./asm/parser

grog-asm: grog-parser
	go build -o grog-asmc ./asm
	./grog-asmc ./asm/jump.asm
	./grog-asmc ./asm/increment-register.asm
	./grog-asmc ./asm/copy-register.asm
	./grog-asmc ./asm/boolean.asm
	./grog-asmc ./asm/arithmetic.asm