antlr := java -jar /opt/antlr-4.7/antlr-4.7.2-complete.jar

build:
	${antlr} -Dlanguage=Go -o parser -package parser Grog.g4

clean:
	rm -Rf parser