cd ..
cd ..
cd fixtures
./cr_fix.sh
cd ..
cd application
cd service
rm service
go build
./service
