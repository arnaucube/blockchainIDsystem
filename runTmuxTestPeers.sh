SESSION='peersTest'

tmux new-session -d -s $SESSION
tmux split-window -d -t 0 -v
tmux split-window -d -t 0 -h
tmux split-window -d -t 1 -h
tmux split-window -d -t 0 -h

tmux split-window -d -t 4 -h
tmux split-window -d -t 5 -h
tmux split-window -d -t 4 -h

tmux send-keys -t 0 'cd peer && go run *.go server 3001 3002' enter
sleep 2
tmux send-keys -t 1 "curl -X POST http://127.0.0.1:3002/register -d '{\"address\": \"firstaddress\"}'" enter
sleep 1
tmux send-keys -t 1 'cd peer && go run *.go client 3003 3004' enter
tmux send-keys -t 2 'cd peer && go run *.go client 3005 3006' enter
tmux send-keys -t 3 'cd peer && go run *.go client 3007 3008' enter

tmux send-keys -t 4 'cd serverCA && go run *.go' enter
tmux send-keys -t 5 'cd serverIDsigner && go run *.go' enter
tmux send-keys -t 6 'cd clientApp && go run *.go' enter
#tmux send-keys -t 7 'cd clientApp/GUI && electron .' enter

tmux attach


# webapps are at:
# serverCA: http://127.0.0.1:3080
# clientApp: http://127.0.0.1:8080
