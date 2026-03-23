<template>
  <div class="app">
    <!-- 导航栏 -->
    <header class="header">
      <div class="container">
        <div class="logo">35Home</div>
        <nav class="nav">
          <router-link to="/" class="nav-item">首页</router-link>
          <router-link to="/pressure" class="nav-item">压力管理</router-link>
          <router-link to="/content" class="nav-item">内容中心</router-link>
          <router-link to="/community" class="nav-item">社区</router-link>
          <router-link to="/test" class="nav-item">压力测评</router-link>
        </nav>
        <div class="user">
          <router-link to="/login" class="user-item">登录</router-link>
          <router-link to="/register" class="user-item">注册</router-link>
        </div>
      </div>
    </header>

    <!-- 主内容区 -->
    <main class="main">
      <router-view />
    </main>

    <!-- 页脚 -->
    <footer class="footer">
      <div class="container">
        <div class="footer-content">
          <div class="footer-section">
            <h3>关于我们</h3>
            <p>35Home是一个专注于35岁左右中年男性压力管理的平台，致力于为用户提供专业的压力解决方案和支持。</p>
          </div>
          <div class="footer-section">
            <h3>快速链接</h3>
            <ul>
              <li><router-link to="/">首页</router-link></li>
              <li><router-link to="/pressure">压力管理</router-link></li>
              <li><router-link to="/content">内容中心</router-link></li>
              <li><router-link to="/community">社区</router-link></li>
            </ul>
          </div>
          <div class="footer-section">
            <h3>联系我们</h3>
            <p>邮箱：contact@35home.com</p>
            <p>电话：400-123-4567</p>
          </div>
        </div>
        <div class="footer-bottom">
          <p>&copy; 2026 35Home. 保留所有权利。</p>
        </div>
      </div>
    </footer>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: 'Microsoft YaHei', sans-serif;
  font-size: 16px;
  line-height: 1.6;
  color: #333;
  background-color: #f5f5f5;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
}

/* 导航栏 */
.header {
  background-color: #2c3e50;
  color: #fff;
  padding: 15px 0;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.logo {
  font-size: 24px;
  font-weight: bold;
  float: left;
}

.nav {
  float: left;
  margin-left: 50px;
}

.nav-item {
  display: inline-block;
  color: #fff;
  text-decoration: none;
  padding: 0 15px;
  line-height: 40px;
  transition: color 0.3s;
}

.nav-item:hover {
  color: #3498db;
}

.user {
  float: right;
}

.user-item {
  display: inline-block;
  color: #fff;
  text-decoration: none;
  padding: 0 15px;
  line-height: 40px;
  transition: color 0.3s;
}

.user-item:hover {
  color: #3498db;
}

/* 主内容区 */
.main {
  padding: 40px 0;
  min-height: 600px;
}

/* 页脚 */
.footer {
  background-color: #34495e;
  color: #fff;
  padding: 40px 0;
  margin-top: 40px;
}

.footer-content {
  display: flex;
  justify-content: space-between;
  margin-bottom: 30px;
}

.footer-section {
  flex: 1;
  margin-right: 30px;
}

.footer-section h3 {
  margin-bottom: 15px;
  color: #3498db;
}

.footer-section ul {
  list-style: none;
}

.footer-section ul li {
  margin-bottom: 10px;
}

.footer-section ul li a {
  color: #fff;
  text-decoration: none;
  transition: color 0.3s;
}

.footer-section ul li a:hover {
  color: #3498db;
}

.footer-bottom {
  text-align: center;
  padding-top: 20px;
  border-top: 1px solid #444;
}

/* 清除浮动 */
.header::after,
.footer-content::after {
  content: '';
  display: table;
  clear: both;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .nav {
    display: none;
  }
  
  .footer-content {
    flex-direction: column;
  }
  
  .footer-section {
    margin-right: 0;
    margin-bottom: 30px;
  }
}
</style>
