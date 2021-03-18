pushd ./nodejs
rm -rf node_modules
rm -rf dist 
popd

pushd ./lancher 
make clean
popd 