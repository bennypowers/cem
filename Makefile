test:
	go test -json ./... "${@:2}" | go tool tparse -all

update:
	go test -json ./... "${@:2}" --update | go tool tparse -all

# watch:
# 	-while true; do ls ./**/*.{go,scm,ts} | entr -dcr sh -c "make test"; done

