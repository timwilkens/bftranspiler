cmake: transpiler.go hello.bf
	go run transpiler.go -c=hello.c hello.bf
	gcc -Wall -Wextra -o hello hello.c 

clean:
	rm hello hello.c
