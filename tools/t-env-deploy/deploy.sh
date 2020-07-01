
if [ ! $1 ];then
  echo "指定发布到哪个server"
  exit 1
fi

deployto=~/tmp/lotus-deploy
rm -rf $deployto
mkdir -p $deployto

cd ../..
propath=`pwd`
make all

mv lotus lotus-storage-miner lotus-seal-worker $deployto
cd tools/t-env-deploy

cp restart-miner.sh restart-node.sh export-env.sh $deployto

cd ~/tmp

rm lotus-deploy.tgz
tar cvfz lotus-deploy.tgz lotus-deploy

scp lotus-deploy.tgz ipfsmain@$1:~/
