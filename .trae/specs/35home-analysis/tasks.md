# 35home项目 - 实现计划（分解和优先级任务列表）

## [x] Task 1: 实现用户注册功能
- **Priority**: P0
- **Depends On**: None
- **Description**: 
  - 实现用户注册的完整业务逻辑，包括数据验证、密码加密、用户创建
  - 确保邮箱唯一性验证，密码长度至少6位
  - 更新service层的Register方法实现
- **Acceptance Criteria Addressed**: AC-1
- **Test Requirements**:
  - `programmatic` TR-1.1: POST /api/user/register with valid data returns 201 status code
  - `programmatic` TR-1.2: POST /api/user/register with duplicate email returns 409 status code
  - `programmatic` TR-1.3: POST /api/user/register with invalid password (length < 6) returns 400 status code

## [x] Task 2: 实现用户登录功能
- **Priority**: P0
- **Depends On**: Task 1
- **Description**: 
  - 实现用户登录的完整业务逻辑，包括邮箱密码验证、JWT令牌生成
  - 更新service层的Login方法实现
  - 确保密码验证和用户认证安全
- **Acceptance Criteria Addressed**: AC-2
- **Test Requirements**:
  - `programmatic` TR-2.1: POST /api/user/login with valid credentials returns 200 status code and JWT token
  - `programmatic` TR-2.2: POST /api/user/login with invalid credentials returns 401 status code
  - `programmatic` TR-2.3: JWT token contains user ID and has valid expiration time

## [x] Task 3: 实现投票提交功能
- **Priority**: P1
- **Depends On**: Task 2
- **Description**: 
  - 实现投票提交的完整业务逻辑，包括数据验证、用户投票唯一性检查
  - 更新service层的SubmitVote方法实现
  - 确保每个用户只能投票一次
- **Acceptance Criteria Addressed**: AC-3
- **Test Requirements**:
  - `programmatic` TR-3.1: POST /api/vote/submit with valid vote data returns 201 status code
  - `programmatic` TR-3.2: POST /api/vote/submit with duplicate vote returns 409 status code
  - `programmatic` TR-3.3: POST /api/vote/submit with invalid vote data returns 400 status code

## [x] Task 4: 实现投票统计功能
- **Priority**: P1
- **Depends On**: Task 3
- **Description**: 
  - 实现投票统计的完整业务逻辑，包括总票数统计、各分类统计
  - 更新service层的GetVoteStats方法实现
  - 确保统计数据的准确性
- **Acceptance Criteria Addressed**: AC-4
- **Test Requirements**:
  - `programmatic` TR-4.1: GET /api/vote/stats returns 200 status code and valid statistics data
  - `programmatic` TR-4.2: Vote stats contain totalVotes, employedCount, unemployedCount fields
  - `programmatic` TR-4.3: Vote stats contain cityStats, industryStats, occupationStats, ageStats, genderStats maps



## [x] Task 5: 完善数据库配置和连接管理
- **Priority**: P0
- **Depends On**: None
- **Description**: 
  - 完善数据库配置加载和连接管理
  - 实现外部配置文件支持，减少硬编码
  - 优化连接池配置，提高性能
- **Acceptance Criteria Addressed**: NFR-3, NFR-4
- **Test Requirements**:
  - `programmatic` TR-5.1: Database connection can be configured via environment variables
  - `programmatic` TR-5.2: Connection pool settings are properly configured
  - `human-judgment` TR-5.3: Configuration should support both development and production environments

## [x] Task 6: 实现跨域支持和安全中间件
- **Priority**: P1
- **Depends On**: None
- **Description**: 
  - 完善跨域中间件配置
  - 添加请求验证和安全头设置
  - 确保API的安全性和可靠性
- **Acceptance Criteria Addressed**: NFR-1, NFR-5
- **Test Requirements**:
  - `programmatic` TR-6.1: API supports CORS with proper headers
  - `programmatic` TR-6.2: Input validation is performed for all API endpoints
  - `human-judgment` TR-6.3: Security headers are properly set

## [x] Task 7: 更新数据库模型规范
- **Priority**: P0
- **Depends On**: None
- **Description**: 
  - 将所有模型的主键类型从uint改为int64（long型）
  - 将所有时间字段从time.Time改为int64（时间戳）
  - 更新相关的数据库迁移和查询逻辑
- **Acceptance Criteria Addressed**: Database Schema constraint
- **Test Requirements**:
  - `programmatic` TR-7.1: All model primary keys are of type int64
  - `programmatic` TR-7.2: All time fields are stored as int64 timestamps
  - `programmatic` TR-7.3: Database migration works correctly with new schema
  - `human-judgment` TR-7.4: Code remains readable and maintainable