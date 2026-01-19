# MiniBlog API 接口文档（中文版）

## 项目概述

技术栈：

- Web 框架：Gin
- ORM：GORM
- 数据库：PostgreSQL
- 缓存/会话：Redis
- 认证方式：JWT（Redis 用于 token 黑名单及部分缓存）

当前版本：v1 基础路径：/api/v1

## 通用说明

- 认证方式：大部分接口需要携带 Authorization: Bearer <jwt-token>
- 时间格式：ISO 8601（例：2026-01-19T11:45:00+08:00）
- 返回格式：统一 JSON
- 错误返回示例：

JSON

```json
{
  "code": 400,
  "message": "参数验证失败：标题不能为空"
}
```

## 数据模型参考

| 字段       | User         | Blog     | Comment  | Like     |
| ---------- | ------------ | -------- | -------- | -------- |
| id         | 自增主键     | 自增主键 | 自增主键 | 自增主键 |
| user_id    | -            | 外键     | 外键     | 外键     |
| blog_id    | -            | -        | 外键     | 外键     |
| username   | 用户名       | -        | -        | -        |
| email      | 邮箱（唯一） | -        | -        | -        |
| password   | 加密后密码   | -        | -        | -        |
| title      | -            | 标题     | -        | -        |
| content    | -            | 文章内容 | 评论内容 | -        |
| created_at | 创建时间     | 创建时间 | 创建时间 | 创建时间 |

## 接口列表

### 1. 用户相关

#### 1.1 注册

```http
POST /api/v1/users/register
```

**请求体**

```JSON
{
  "username": "kevin123",
  "email": "kevin@example.com",
  "password": "12345678"
}
```

**成功响应** (201)

```JSON
{
  "message": "注册成功",
  "user": {
    "id": 1,
    "username": "kevin123",
    "email": "kevin@example.com"
  }
}
```

#### 1.2 登录

```http
POST /api/v1/users/login
```

**请求体**

```JSON
{
  "email": "kevin@example.com",
  "password": "12345678"
}
```

**成功响应** (200)

```JSON
{
  "message": "登录成功",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

#### 1.3 退出登录（加入黑名单）

```http
POST /api/v1/users/logout
```

需要携带 token

**成功响应** (200)

```JSON
{
  "message": "已安全退出登录"
}
```

### 2. 博客（文章）相关

#### 2.1 发布文章

```http
POST /api/v1/blogs
```

需要认证

**请求体**

```JSON
{
  "title": "今天天气真好",
  "content": "阳光明媚，心情也跟着变好了..."
}
```

**成功响应** (201)

```JSON
{
  "message": "文章发布成功",
  "blog": {
    "id": 1,
    "title": "今天天气真好",
    "content": "阳光明媚...",
    "user_id": 1,
    "created_at": "2026-01-19T11:45:00+08:00"
  }
}
```

#### 2.2 获取文章列表（分页）

```http
GET /api/v1/blogs?page=1&limit=10&sort=latest
```

**可选查询参数**

- page: 当前页（默认1）
- limit: 每页数量（默认10，最大50）
- sort: latest / hottest（最新/最热）

**成功响应** (200)

```JSON
{
  "data": [
    {
      "id": 1,
      "title": "今天天气真好",
      "content": "阳光明媚...", // 可截断显示
      "user_id": 1,
      "username": "kevin123",
      "like_count": 23,
      "comment_count": 8,
      "created_at": "2026-01-19T11:45:00+08:00"
    }
  ],
  "pagination": {
    "total": 128,
    "page": 1,
    "limit": 10,
    "pages": 13
  }
}
```

#### 2.3 获取单篇文章详情

```http
GET /api/v1/blogs/:id
```

**成功响应** (200)

```JSON
{
  "id": 1,
  "title": "今天天气真好",
  "content": "完整内容...",
  "user_id": 1,
  "username": "kevin123",
  "like_count": 23,
  "is_liked": true,          // 当前用户是否已点赞（登录后才有）
  "created_at": "2026-01-19T11:45:00+08:00"
}
```

#### 2.4 修改文章（仅作者）

```http
PUT /api/v1/blogs/:id
```

**请求体**（支持部分更新）

```JSON
{
  "title": "今天天气真的很好",
  "content": "修改后的内容..."
}
```

#### 2.5 删除文章（仅作者）

```http
DELETE /api/v1/blogs/:id
```

### 3. 评论相关

#### 3.1 添加评论

```http
POST /api/v1/blogs/:blog_id/comments
```

需要认证

**请求体**

```JSON
{
  "content": "写得真好，学习了！"
}
```

#### 3.2 获取某篇文章评论列表

```http
GET /api/v1/blogs/:blog_id/comments?page=1&limit=20
```

**成功响应示例**

```JSON
{
  "data": [
    {
      "id": 1,
      "content": "写得真好，学习了！",
      "user_id": 2,
      "username": "alice",
      "created_at": "2026-01-19T12:03:00+08:00"
    }
  ],
  "pagination": { ... }
}
```

#### 3.3 删除评论（评论作者或文章作者）

```http
DELETE /api/v1/comments/:id
```

### 4. 点赞相关

#### 4.1 点赞/取消点赞（幂等）

```http
POST /api/v1/blogs/:blog_id/like
```

需要认证，无请求体

**成功响应**

```JSON
{
  "message": "已点赞",   // 或 "已取消点赞"
  "like_count": 24
}
```

#### 4.2 获取点赞数量（可选带点赞用户列表）

```http
GET /api/v1/blogs/:blog_id/likes?with_users=false
```

**响应示例**

```JSON
{
  "like_count": 24,
  "users": [1, 3, 7, 12]   // 当 with_users=true 时返回
}
```

## 推荐的响应状态码一览

| 状态码 | 含义            | 常见场景               |
| ------ | --------------- | ---------------------- |
| 200    | 成功            | GET、PUT、POST成功     |
| 201    | 创建成功        | 注册、发文、评论       |
| 400    | 请求参数错误    | 缺少必填字段、格式错误 |
| 401    | 未授权/登录过期 | 缺少token或token失效   |
| 403    | 无权限          | 改删他人文章/评论      |
| 404    | 资源不存在      | 文章/评论/用户不存在   |
| 409    | 冲突            | 用户名/邮箱已存在      |
| 500    | 服务器内部错误  | 数据库异常等           |