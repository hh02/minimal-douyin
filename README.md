# Mininal-douyin

- [x] follow (关注服务)
  - [x] CreateFollow
  - [x] DeleteFollow
  - [x] QueryFollow
  - [x] QueryFollower
  - [x] CheckFollow

服务之间的调用关系

CreateFollow -> user.AddFollowCount(+1)
DeleteFollow -> user.AddFollowCount(-1)
QueryFollow -> user.MGetUser
QueryFollower -> user.MGetUser