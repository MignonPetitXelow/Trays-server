SRC = $(shell find sources/ -type f -name '*.go')

CCL = go
CCW = go.exe
OUTPUT = "trays-server.exe"

all: $(NAME)

$(NAME):
	${CCW} build -o $(OUTPUT) -buildmode=exe -a -v $(SRC)

run: $(NAME)
	$(OUTPUT)

clean:
	rm $(OUTPUT)