# blog_backend
costwen's blog_backend

## 文件说明

- conf：用于存储配置文件
- middleware：应用中间件
- models：应用数据库模型
- pkg：第三方包
- routers 路由逻辑处理
- runtime：应用运行时数据

## 数据格式

```json
{
    "id": "id",
    "title": "文章名称",
    "tag": ["文章标签", "xxx"],
    "content": "内容",
    "create_time": "创建时间",
    "modify_time": "修改时间",
    "page_view": "浏览量",
}
```