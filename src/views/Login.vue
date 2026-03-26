<template>
  <div class="login">
    <div class="container">
      <div class="login-card">
        <h1>登录</h1>
        <form @submit.prevent="handleLogin">
          <div class="form-group">
            <label for="email">邮箱</label>
            <input type="email" id="email" v-model="form.email" required>
          </div>
          <div class="form-group">
            <label for="password">密码</label>
            <input type="password" id="password" v-model="form.password" required>
          </div>
          <div class="form-actions">
            <button type="submit" class="login-btn">登录</button>
          </div>
          <div class="form-footer">
            <p>还没有账号？<router-link to="/register">立即注册</router-link></p>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Login',
  data() {
    return {
      form: {
        email: '',
        password: ''
      }
    }
  },
  methods: {
    handleLogin() {
      // 登录逻辑
      console.log('登录表单提交:', this.form);
      
      // 调用后端API进行登录
      fetch('http://localhost:8000/api/user/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(this.form)
      })
      .then(response => response.json())
      .then(data => {
        if (data.code === 'SUCCESS') {
          // 登录成功，保存用户信息到localStorage
          localStorage.setItem('user', JSON.stringify(data.respBody.user));
          localStorage.setItem('token', data.respBody.token);
          alert('登录成功！');
          this.$router.push('/');
        } else {
          // 登录失败，显示错误信息
          alert(data.msg);
        }
      })
      .catch(error => {
        console.error('登录失败:', error);
        alert('登录失败，请稍后重试');
      });
    }
  }
}
</script>

<style scoped>
.login {
  padding: 60px 0;
  background-color: #f5f5f5;
  min-height: 80vh;
  display: flex;
  align-items: center;
}

.login-card {
  background-color: #fff;
  padding: 40px;
  border-radius: 8px;
  box-shadow: 0 5px 20px rgba(0, 0, 0, 0.1);
  max-width: 400px;
  margin: 0 auto;
  width: 100%;
}

.login-card h1 {
  font-size: 28px;
  margin-bottom: 30px;
  color: #2c3e50;
  text-align: center;
}

.form-group {
  margin-bottom: 20px;
}

.form-group label {
  display: block;
  margin-bottom: 8px;
  color: #2c3e50;
  font-weight: bold;
}

.form-group input {
  width: 100%;
  padding: 12px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 16px;
  transition: border-color 0.3s ease;
}

.form-group input:focus {
  outline: none;
  border-color: #3498db;
  box-shadow: 0 0 0 2px rgba(52, 152, 219, 0.1);
}

.form-actions {
  margin-bottom: 20px;
}

.login-btn {
  width: 100%;
  padding: 12px;
  background-color: #3498db;
  color: #fff;
  border: none;
  border-radius: 4px;
  font-size: 16px;
  font-weight: bold;
  cursor: pointer;
  transition: all 0.3s ease;
}

.login-btn:hover {
  background-color: #2980b9;
  transform: translateY(-2px);
}

.form-footer {
  text-align: center;
  color: #666;
}

.form-footer a {
  color: #3498db;
  text-decoration: none;
  font-weight: bold;
  transition: color 0.3s ease;
}

.form-footer a:hover {
  color: #2980b9;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .login-card {
    padding: 30px;
    margin: 0 20px;
  }
  
  .login-card h1 {
    font-size: 24px;
  }
}
</style>
