# 35home项目 - 产品需求文档

## Overview
- **Summary**: 35home是一个集成用户管理和投票系统的全栈应用平台，旨在提供社区互动和数据收集功能。
- **Purpose**: 解决用户注册登录和社区投票等核心业务场景，为用户提供完整的在线服务体验。
- **Target Users**: 需要注册登录的用户群体，参与投票的社区成员。

## Goals
- 提供完整的用户认证系统（注册、登录）
- 实现社区投票功能，支持投票提交和统计查询
- 确保系统的安全性、可靠性和性能

## Non-Goals (Out of Scope)
- 不包含复杂的权限管理系统
- 不支持文件上传功能
- 不包含实时通信功能
- 不支持社交分享功能

## Background & Context
- 项目采用前后端分离架构，后端使用Go语言+Gin框架，前端使用Vue.js
- 数据库采用MySQL，使用GORM进行ORM操作
- 项目已实现基本的用户管理和投票功能，其他功能处于开发中状态
- 当前配置使用环境变量和默认值，需要支持外部配置

## Functional Requirements
- **FR-1**: 用户注册功能 - 支持新用户注册，包括用户名、邮箱、密码等基本信息收集
- **FR-2**: 用户登录功能 - 支持用户通过邮箱和密码登录，返回JWT令牌
- **FR-3**: 投票提交功能 - 支持用户提交就业状态、城市、行业等投票数据
- **FR-4**: 投票统计功能 - 支持查询投票统计数据，包括总票数、各分类统计等

## Non-Functional Requirements
- **NFR-1**: 安全性 - 用户密码必须加密存储，接口必须进行输入验证
- **NFR-2**: 性能 - API响应时间应在200ms以内（正常负载下）
- **NFR-3**: 可靠性 - 系统应具备错误处理机制，避免崩溃
- **NFR-4**: 可扩展性 - 代码结构应支持功能模块的独立扩展
- **NFR-5**: 跨域支持 - 后端API必须支持前端跨域请求

## Constraints
- **Technical**: Go 1.x, Gin框架, MySQL数据库, GORM ORM
- **Business**: 数据库连接参数通过环境变量配置，支持开发和生产环境
- **Dependencies**: 依赖第三方库包括gin-gonic/gin, jinzhu/gorm, go-sql-driver/mysql
- **Database Schema**: 所有表的主键使用long型id字段，所有时间字段使用long型时间戳

## Assumptions
- 用户邮箱是唯一标识符，不允许重复注册
- 每个用户只能投票一次，防止重复投票
- 数据库连接参数可以通过环境变量配置，默认值可用于开发环境
- JWT密钥需要在生产环境中安全配置

## Acceptance Criteria

### AC-1: 用户注册功能
- **Given**: 用户提供有效的注册信息（用户名、邮箱、密码等）
- **When**: 调用POST /api/user/register接口
- **Then**: 返回201状态码和注册成功消息，数据库中创建用户记录
- **Verification**: `programmatic`
- **Notes**: 邮箱必须唯一，密码长度至少6位

### AC-2: 用户登录功能
- **Given**: 用户提供正确的邮箱和密码
- **When**: 调用POST /api/user/login接口
- **Then**: 返回200状态码、用户信息和JWT令牌
- **Verification**: `programmatic`

### AC-3: 投票提交功能
- **Given**: 用户提供有效的投票数据（就业状态、城市、行业等）
- **When**: 调用POST /api/vote/submit接口
- **Then**: 返回201状态码和投票成功消息，数据库中创建投票记录
- **Verification**: `programmatic`
- **Notes**: 用户只能投票一次，重复投票应返回409状态码

### AC-4: 投票统计功能
- **Given**: 系统中有投票数据
- **When**: 调用GET /api/vote/stats接口
- **Then**: 返回200状态码和投票统计数据（总数、各分类统计）
- **Verification**: `programmatic`

## Open Questions
- [ ] 是否需要支持用户资料编辑功能？
- [ ] 是否需要添加用户权限管理？
- [ ] 是否需要支持数据导出功能？