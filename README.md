# TickTok

2022 ByteDance Backend Summer Camp Project (inspired by TikTok)

## Contributor

Thanks to the following developers for their support of the project. (In no particular order)

[@zhihaop](https://github.com/zhihaop), [@northmachine](https://github.com/northmachine),
[@JukieChen](https://github.com/JukieChen), [@WYAOBO](https://github.com/WYAOBO), [@xjmxyt](https://github.com/xjmxyt)

## Quick Start

Go to the project directory and run the following shell commands.

```shell
go run app/main.go
```

## Project Structure

The project structure is inspired by [go-clean-arch](https://github.com/bxcodec/go-clean-arch). We divide the project
into domains according to business requirements.

```
├───app                 // bootstrap of the application
├───comment             // comment domain
│   ├───controller
│   ├───repository
│   └───service
├───core                // common utils or interfaces
│   ├───controller
│   ├───repository
│   └───service
├───entity              // entities (domains) of the application and their mocks
│   └───mocks
├───favourite           // favourite (user to video) domain
│   ├───controller
│   ├───repository
│   └───service
├───clip                // clip's domain
│   ├───controller
│   ├───repository
│   └───service
└───user                // user's domain
    ├───controller
    ├───repository
    └───service
```

## RESTFul Api Documentation

We use the RESTFul api interface provided by ByteDance, see
also [Apifox](https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18345145).