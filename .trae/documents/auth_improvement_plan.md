# 认证功能改进实现计划

## [ ] 任务1：后端添加用户名重复校验
- **优先级**: P0
- **Depends On**: None
- **Description**: 
  - 在 UserRepository 中添加 ExistsByUsername 方法
  - 在 user_service.go 的 Register 方法中添加用户名重复校验
- **Success Criteria**:
  - 注册时用户名重复会返回适当的错误信息
- **Test Requirements**:
  - `programmatic`: 尝试使用已存在的用户名注册，应返回错误
  - `programmatic`: 使用新用户名注册应成功

## [ ] 任务2：修改注册逻辑，返回JWT token
- **优先级**: P0
- **Depends On**: 任务1
- **Description**: 
  - 修改 Register 方法，在注册成功后生成并返回 JWT token
  - 更新 RegisterResponse DTO 结构以包含 token 字段
- **Success Criteria**:
  - 注册成功后返回包含用户信息和 token 的完整响应
- **Test Requirements**:
  - `programmatic`: 注册成功响应中包含 token 字段
  - `programmatic`: 返回的 token 可以用于后续认证

## [ ] 任务3：修改前端注册页面逻辑
- **优先级**: P0
- **Depends On**: 任务2
- **Description**: 
  - 修改 Register.vue，注册成功后保存 token 和用户信息到 localStorage
  - 直接跳转到主页，不再跳转到登录页
  - 移除成功提示弹窗，只保留异常错误提示
- **Success Criteria**:
  - 注册成功后自动登录并跳转到主页
  - 没有成功提示弹窗
- **Test Requirements**:
  - `human-judgement`: 注册成功后直接跳转到主页
  - `human-judgement`: 只在异常情况下显示弹窗

## [ ] 任务4：修改登录页面支持用户名登录
- **优先级**: P1
- **Depends On**: None
- **Description**: 
  - 修改 Login.vue，默认使用用户名登录
  - 修改后端 Login 方法支持按用户名查找用户
  - 移除成功提示弹窗
- **Success Criteria**:
  - 支持使用用户名登录
  - 没有成功提示弹窗
- **Test Requirements**:
  - `programmatic`: 可以使用用户名成功登录
  - `human-judgement`: 登录成功后直接跳转到主页，无弹窗

## [ ] 任务5：添加登出功能
- **优先级**: P1
- **Depends On**: None
- **Description**: 
  - 在 Home.vue 中添加登出按钮
  - 实现登出逻辑，清除 localStorage 中的用户信息和 token
  - 登出后跳转到登录页面
- **Success Criteria**:
  - 点击登出按钮后清除登录状态并跳转
- **Test Requirements**:
  - `human-judgement`: 登出后 localStorage 为空
  - `human-judgement`: 登出后重定向到登录页面

## [ ] 任务6：统一错误处理，只显示异常错误弹窗
- **优先级**: P2
- **Depends On**: None
- **Description**: 
  - 检查所有前端组件的错误处理逻辑
  - 移除成功操作的弹窗提示
  - 只保留异常错误的弹窗提示
- **Success Criteria**:
  - 所有成功操作都没有弹窗提示
  - 只有异常错误才显示弹窗
- **Test Requirements**:
  - `human-judgement`: 成功操作无弹窗
  - `human-judgement`: 异常错误有弹窗提示