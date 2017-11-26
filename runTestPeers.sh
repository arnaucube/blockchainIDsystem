cd peer
echo "compile Go"
echo "go build"
go build

sleep 3

echo "peer server"
xterm -hold -e './peer server 3001 3002' &

sleep 2

echo "register the first block"
curl -X POST http://127.0.0.1:3002/register -d '{"address": "firstaddress"}'

sleep 2

echo "peer 1"
xterm -hold -e './peer client 3003 3004' &

sleep 2

echo "peer 2"
xterm -hold -e './peer client 3005 3006' &

sleep 2

echo "peer 3"
xterm -hold -e './peer client 3007 3008' &
