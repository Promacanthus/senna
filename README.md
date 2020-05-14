# SENNA

Sample implementation using jwt

[Medium原文链接](https://medium.com/@Raulgzm/securing-golang-api-using-json-web-token-jwt-2dc363792a48)

[github仓库地址](https://github.com/brainattica/golang-jwt-authentication-api-sample)

|原文|本文|
|---|---|
|[mux](https://github.com/gorilla/mux)|[gin](https://github.com/gin-gonic/gin)|
|[jwt-go/v3](https://github.com/dgrijalva/jwt-go)|[jwt-go/v4](https://github.com/dgrijalva/jwt-go)|
||[viper](https://github.com/spf13/viper)|

这个项目是一个典型的MVC架构。

- JWT(Json Web Tokens)相关概念点[这里](https://promacanthus.netlify.app/%E5%BC%80%E5%8F%91%E6%A1%86%E6%9E%B6/jwt/01-jwt%E7%AE%80%E4%BB%8B/)

使用模块化来组织代码，有如下好处：

1. 整个应用程序布局更易于理解。
2. 可以看到部件如何协同工作，因为必须在使用前注入模块。
3. 所有必需的功能都包含在模块内部，因此代码可重复使用。
4. 更易于测试代码。
