# 前端 JWT 认证实现计划

## [x] 任务1：修改登录和注册页面的成功判断
- **优先级**: P0
- **依赖**: None
- **描述**: 修改 Login.vue 和 Register.vue，将成功判断从 `data.code === 'SUCCESS'` 改为 `data.code === 0`
- **成功标准**: 
  - 登录和注册成功判断使用数字错误码
- **测试要求**:
  - `programmatic` TR-1.1: 前端页面能够正确判断登录和注册成功
  - `human-judgement` TR-1.2: 用户体验正常

## [x] 任务2：修改 Home.vue 添加 JWT 认证
- **优先级**: P0
- **依赖**: None
- **描述**: 修改 Home.vue，在投票相关API请求中添加 Authorization header
- **成功标准**: 
  - 获取投票统计时添加 JWT 令牌
  - 提交投票时添加 JWT 令牌
- **测试要求**:
  - `programmatic` TR-2.1: 能够成功获取投票统计数据
  - `programmatic` TR-2.2: 能够成功提交投票
  - `human-judgement` TR-2.3: 用户体验正常

## [x] 任务3：测试完整的前后端交互流程
- **优先级**: P1
- **依赖**: 任务1、任务2
- **描述**: 测试从登录到投票的完整流程
- **成功标准**: 
  - 用户能够正常登录
  - 登录后能够查看投票统计
  - 登录后能够提交投票
- **测试要求**:
  - `programmatic` TR-3.1: 所有API调用成功
  - `human-judgement` TR-3.2: 完整流程验证通过
