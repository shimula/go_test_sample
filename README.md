[![wercker status](https://app.wercker.com/status/eba45d9a0338240289150a8186233db8/s/master "wercker status")](https://app.wercker.com/project/bykey/eba45d9a0338240289150a8186233db8)

## run mysql

```
docker run -d -p 3306:3306 -v $(pwd)/test_data:/docker-entrypoint-initdb.d -e MYSQL_ROOT_PASSWORD=rootpassword -e MYSQL_USER=testuser -e MYSQL_PASSWORD=password  mysql:5.5 --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci
```