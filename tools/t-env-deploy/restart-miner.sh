source export-env.sh
ps aux|grep lotus|awk '{print $2}'|xargs kill -9
sleep 3
nohup ./lotus-storage-miner run > lotus.log &

