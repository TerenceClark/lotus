
deployto=~/tmp/lotus-deploy
mkdir -p $deployto

cd ../..
propath=`pwd`
make all

mv lotus lotus-storage-miner lotus-seal-worker $deployto
cd tools/t-env-deploy

cp restart-miner.sh restart-node.sh export-env.sh $deployto

cd ~/tmp

tar cvfz lotus-deploy.tgz lotus-deploy

scp lotus-deploy.tgz ipfsmain@ts1:~/
