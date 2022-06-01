# Follow Service

- [x] follow (关注服务)
  - [x] CreateFollow
  - [x] DeleteFollow
  - [x] QueryFollow
  - [x] QueryFollower
  - [x] CheckFollow

服务之间的调用关系

- CreateFollow -> user.AddFollowCount(+1)
- DeleteFollow -> user.AddFollowCount(-1)
- QueryFollow -> user.MGetUser
- QueryFollower -> user.MGetUser

## TODO

- CreateFollow会调用AddFollowCount，如何保证数据的一致性
- 删除冗余的服务：比如说CheckFollow好像就没有用到