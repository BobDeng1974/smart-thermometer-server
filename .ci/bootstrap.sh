dep ensure

## Import local directories as vendor items
cd vendor
ln -s ../controllers controllers
ln -s ../utils utils

cd ..