source export-env.sh
./lotus daemon stop
sleep 3
nohup ./lotus daemon > lotus.log &
