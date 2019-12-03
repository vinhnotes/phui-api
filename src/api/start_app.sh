# !/bin/bash

# MySQLサーバーが起動するまで待機する
until mysqladmin ping -h mysql -P 3306 --silent; do
  echo 'Waiting for mysqld to be connectable...'
  sleep 2
done

echo "App is starting...!"
exec go run main.go