#### mongo
```
1. shell mkdir dir

2. docker pull mongo

3. docker run -d --name my-mongo-auth -v /Users/cc/mongo/datadb:/data/db -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=root  --privileged=true mongo

4. 
use ginframe // 选择test数据库
db.createUser(
{
user: "test",
pwd: "123456",
roles: [
{ role: "readWrite", db: "test" }
]
}
);

exit

mongo -u test -p 123456 --authenticationDatabase ginframe (以刚创建的test用户登录)
```

