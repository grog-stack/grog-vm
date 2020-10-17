antlr := java -jar /opt/antlr-4.7/antlr-4.7.2-complete.jar

build:
	${antlr} -Dlanguage=Go -package asm Grog.g4
	go build

clean:
	rm -Rf parser
	rm grog-asm