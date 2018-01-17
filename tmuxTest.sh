SESSION='peersTest'

tmux new-session -d -s $SESSION
tmux split-window -d -t 0 -v
tmux split-window -d -t 0 -h
tmux split-window -d -t 1 -h
tmux split-window -d -t 0 -h

tmux split-window -d -t 4 -h
tmux split-window -d -t 5 -h
tmux split-window -d -t 4 -h

tmux attach
