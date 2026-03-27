# Response 优化实现计划

## [x] 任务1：修改 Response 结构体的 Code 类型
- **优先级**: P0
- **依赖**: None
- **描述**: 修改 dto/response.go，将 Response 结构体的 Code 字段从 int 类型改为 common.ErrorCode 类型
- **成功标准**: 
  - Response.Code 字段类型为 common.ErrorCode
  - JSON 标签保持不变
- **测试要求**:
  - `programmatic` TR-1.1: 编译通过，无类型错误
  - `human-judgement` TR-1.2: 代码结构清晰

## [x] 任务2：添加辅助函数创建 Response
- **优先级**: P0
- **依赖**: 任务1
- **描述**: 在 dto/response.go 中添加辅助函数，根据错误码自动设置 Msg 字段
- **成功标准**: 
  - 提供创建成功响应的函数（如 NewSuccessResponse）
  - 提供创建错误响应的函数（如 NewErrorResponse）
  - 自动从 ErrorMessages 获取 Msg
- **测试要求**:
  - `programmatic` TR-2.1: 编译通过，无类型错误
  - `human-judgement` TR-2.2: 函数设计合理，使用方便

## [x] 任务3：更新 user_handler 使用新的辅助函数
- **优先级**: P1
- **依赖**: 任务1、任务2
- **描述**: 修改 handler/user_handler.go，使用新的辅助函数创建响应
- **成功标准**: 
  - 不再手动设置 Code 和 Msg
  - 使用辅助函数创建响应对象
- **测试要求**:
  - `programmatic` TR-3.1: 编译通过，无类型错误
  - `human-judgement` TR-3.2: 代码简洁，逻辑清晰

## [x] 任务4：更新 vote_handler 使用新的辅助函数
- **优先级**: P1
- **依赖**: 任务1、任务2
- **描述**: 修改 handler/vote_handler.go，使用新的辅助函数创建响应
- **成功标准**: 
  - 不再手动设置 Code 和 Msg
  - 使用辅助函数创建响应对象
- **测试要求**:
  - `programmatic` TR-4.1: 编译通过，无类型错误
  - `human-judgement` TR-4.2: 代码简洁，逻辑清晰
