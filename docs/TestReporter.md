# 测试报告

关于CI/CD部分，直接采用github上附带的Action功能，进行自动化部署和测试

## 测试模块

- [x]注册系统
- [x]登录系统
- [x]认证系统

## 测试结果

```
=== RUN   TestRegister
[GIN] 2020/08/05 - 01:02:24 | 200 |            0s |       192.0.2.1 | GET      "/users"
[GIN] 2020/08/05 - 01:02:24 | 200 |     80.8127ms |       192.0.2.1 | POST     "/register"
[GIN] 2020/08/05 - 01:02:24 | 200 |       968.4µs |       192.0.2.1 | GET      "/users"
--- PASS: TestRegister (0.17s)
=== RUN   TestLogin
[GIN] 2020/08/05 - 01:02:25 | 200 |       999.6µs |       192.0.2.1 | GET      "/users"
[GIN] 2020/08/05 - 01:02:25 | 200 |     76.8182ms |       192.0.2.1 | POST     "/register"
[GIN] 2020/08/05 - 01:02:25 | 200 |     69.7892ms |       192.0.2.1 | POST     "/login"
--- PASS: TestLogin (0.23s)
=== RUN   TestAuth
[GIN] 2020/08/05 - 01:02:25 | 200 |            0s |       192.0.2.1 | GET      "/guestPage"
[GIN] 2020/08/05 - 01:02:25 | 200 |            0s |       192.0.2.1 | GET      "/userPage"
[GIN] 2020/08/05 - 01:02:25 | 200 |     79.7583ms |       192.0.2.1 | POST     "/register"
[GIN] 2020/08/05 - 01:02:25 | 200 |     68.8702ms |       192.0.2.1 | POST     "/login"
[GIN] 2020/08/05 - 01:02:25 | 200 |            0s |       192.0.2.1 | GET      "/guestPage"
[GIN] 2020/08/05 - 01:02:25 | 200 |       942.3µs |       192.0.2.1 | GET      "/userPage"
--- PASS: TestAuth (0.23s)
PASS
```