BIN = bin/
CMD = ./cmd/

NAME_BINS = timer\
			timer_continue\
			timer_stop\

PATH_BINS = $(addprefix $(BIN), $(NAME_BINS))

all: $(BIN) $(PATH_BINS)

$(BIN):
	mkdir -p $(BIN)

$(PATH_BINS):
	go build -o $(BIN) $(CMD)...

clean:
	rm -rf $(PATH_BINS)
	rm -rf $(BIN)

re: clean $(PATH_BINS)
