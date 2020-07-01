ps aux|grep lotus|awk '{print $2}'|xargs kill -9
nohup ./lotus-storage-miner run > lotus.log &

