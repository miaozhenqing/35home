# 错误码重新定义实现计划

## [x] 任务1：重新定义错误码类型

* **优先级**: P0

* **依赖**: None

* **描述**: 修改 common/error\_code.go，将错误码从 string 类型改为 int 类型，添加 ErrorCode 类型定义和常量

* **成功标准**:

  * 所有错误码都是 int 类型

  * 定义了 ErrorCode 类型

  * 包含 ErrSuccess = 0

* **测试要求**:

  * `programmatic` TR-1.1: 编译通过，无类型错误

  * `human-judgement` TR-1.2: 代码结构清晰，错误码定义完整

## [x] 任务2：调整 Response 结构体

* **优先级**: P0

* **依赖**: 任务1

* **描述**: 修改 dto/response.go，将 Response 结构体的 Code 字段从 string 类型改为 int 类型

* **成功标准**:

  * Response.Code 字段类型为 int

  * JSON 标签保持不变

* **测试要求**:

  * `programmatic` TR-2.1: 编译通过，无类型错误

  * `human-judgement` TR-2.2: Response 结构体定义符合规范

## \[x] 任务3：修改 user\_handler 使用新错误码

* **优先级**: P1

* **依赖**: 任务1、任务2

* **描述**: 修改 handler/user\_handler.go，将所有字符串错误码替换为新的 int 类型错误码常量

* **成功标准**:

  * 不再使用字符串字面量作为错误码

  * 所有错误码都使用新定义的常量

* **测试要求**:

  * `programmatic` TR-3.1: 编译通过，无类型错误

  * `human-judgement` TR-3.2: 代码风格一致，错误处理逻辑正确

## \[x] 任务4：修改 vote\_handler 使用新错误码

* **优先级**: P1

* **依赖**: 任务1、任务2

* **描述**: 修改 handler/vote\_handler.go，将所有字符串错误码替换为新的 int 类型错误码常量

* **成功标准**:

  * 不再使用字符串字面量作为错误码

  * 所有错误码都使用新定义的常量

* **测试要求**:

  * `programmatic` TR-4.1: 编译通过，无类型错误

  * `human-judgement` TR-4.2: 代码风格一致，错误处理逻辑正确

## \[x] 任务5：修改 user\_service 使用新错误码

* **优先级**: P1

* **依赖**: 任务1

* **描述**: 修改 service/user\_service.go，将所有字符串错误码替换为新的 int 类型错误码常量

* **成功标准**:

  * 不再使用字符串字面量作为错误码

  * 所有错误码都使用新定义的常量

* **测试要求**:

  * `programmatic` TR-5.1: 编译通过，无类型错误

  * `human-judgement` TR-5.2: 代码风格一致，错误处理逻辑正确

## \[x] 任务6：修改 vote\_service 使用新错误码

* **优先级**: P1

* **依赖**: 任务1

* **描述**: 修改 service/vote\_service.go，将所有字符串错误码替换为新的 int 类型错误码常量

* **成功标准**:

  * 不再使用字符串字面量作为错误码

  * 所有错误码都使用新定义的常量

* **测试要求**:

  * `programmatic` TR-6.1: 编译通过，无类型错误

  * `human-judgement` TR-6.2: 代码风格一致，错误处理逻辑正确

## [x] 任务7：更新错误信息映射并关联 Response.msg
- **优先级**: P1
- **依赖**: 任务1、任务2
- **描述**: 修改 common/error_code.go 中的 ErrorMessages 映射，将 key 从 string 改为 int，并确保 Response 结构体中的 msg 字段通过 ErrorMessages 映射获取，避免使用字面量
- **成功标准**: 
  - ErrorMessages 的 key 类型为 int
  - 所有错误码都有对应的错误信息
  - Response.msg 字段通过 ErrorMessages 映射获取，不再使用字面量
- **测试要求**:
  - `programmatic` TR-7.1: 编译通过，无类型错误
  - `human-judgement` TR-7.2: 错误信息映射完整且正确，Response.msg 不再使用字面量

