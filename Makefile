msg-1:
	echo "Levantando parte 1, recuerde que los dias no pueden ser negativos y las fechas siguen el formato YYYY-mm-dd :)enjoy"
msg-2:
	echo "Levantando parte 2, recuerde colocar el file data.csv en la ruta parte2/file/data.csv, en el directorio se encuentra un archivo de muestra, una vez que se ejeccuta se genera un file response.json, donde se encuentra la informacion generada :)enjoy"
run-1: msg-1
	go run parte1/cmd/main.go
run-2: msg-2
	go run parte2/cmd/main.go

compile-1:
	echo "generando compilados parte1, se generan en bin/parte1/"
	GOOS=freebsd GOARCH=386 go build -o bin/parte1/main-freebsd-386 parte1/cmd/main.go
	GOOS=linux GOARCH=386 go build -o bin/parte1/main-linux-386 parte1/cmd/main.go
	GOOS=windows GOARCH=386 go build -o bin/parte1/main-windows-386 parte1/cmd/main.go
compile-2:
	echo "generando compilados parte2, se generan en bin/parte2/"
	GOOS=freebsd GOARCH=386 go build -o bin/parte2/main-freebsd-386 parte2/cmd/main.go
	GOOS=linux GOARCH=386 go build -o bin/parte2/main-linux-386 parte2/cmd/main.go
	GOOS=windows GOARCH=386 go build -o bin/parte2/main-windows-386 parte2/cmd/main.go

compile-all: compile-1 compile-2