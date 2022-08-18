NAME_TIMER = timer
SRC_TIMER = ./src/Timer

NAME_TIMER_CONTINUE = timer_continue
SRC_TIMER_CONTINUE = ./src/Timer_continue

NAME_TIMER_STOP = timer_stop
SRC_TIMER_STOP = ./src/Timer_stop


all: $(NAME_TIMER) $(NAME_TIMER_CONTINUE) $(NAME_TIMER_STOP)

$(NAME_TIMER):
	go build -o $(NAME_TIMER) $(SRC_TIMER)

$(NAME_TIMER_CONTINUE):
	go build -o $(NAME_TIMER_CONTINUE) $(SRC_TIMER_CONTINUE)

$(NAME_TIMER_STOP):
	go build -o $(NAME_TIMER_STOP) $(SRC_TIMER_STOP)

clean:
	rm -rf $(NAME_TIMER)
	rm -rf $(NAME_TIMER_CONTINUE)
	rm -rf $(NAME_TIMER_STOP)
