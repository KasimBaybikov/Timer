BIN = bin/
CMD = ./cmd/

NAME_BINS = timer\
			timer_continue\
			timer_pause\

PATH_BINS = $(addprefix $(BIN), $(NAME_BINS))

all: $(BIN) $(PATH_BINS)

$(BIN):
	mkdir -p $(BIN)

$(PATH_BINS):
	go build -o $(BIN) $(CMD)...
	@echo "change config file on \033[32m$(HOME)/.config/timer/.conf.timer\033[0m"

clean:
	rm -rf $(PATH_BINS)
	rm -rf $(BIN)

re: clean $(PATH_BINS)
