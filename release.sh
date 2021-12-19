if [ ! -d "./release/theway2meal" ]; then
    mkdir -p release/theway2meal
fi
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o release/theway2meal/w2m main.go
cp -r data release/theway2meal/
cp -r view release/theway2meal/

echo "./w2m" > release/theway2meal/run.sh

cd release/
tar czvf m2w.tar theway2meal/*
