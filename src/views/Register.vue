<template>
  <div class="register">
    <div class="container">
      <div class="register-card">
        <h1>注册</h1>
        <form @submit.prevent="handleRegister">
          <div class="form-group">
            <label for="name">姓名</label>
            <input type="text" id="name" v-model="form.name" required>
          </div>
          <div class="form-group">
            <label for="email">邮箱</label>
            <input type="email" id="email" v-model="form.email" required>
          </div>
          <div class="form-group">
            <label for="password">密码</label>
            <input type="password" id="password" v-model="form.password" required>
          </div>
          <div class="form-group">
            <label for="confirmPassword">确认密码</label>
            <input type="password" id="confirmPassword" v-model="form.confirmPassword" required>
          </div>
          <div class="form-actions">
            <button type="submit" class="register-btn">注册</button>
          </div>
          <div class="form-footer">
            <p>已有账号？<router-link to="/login">立即登录</router-link></p>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Register',
  data() {
    return {
      form: {
        name: '',
        email: '',
        password: '',
        confirmPassword: ''
      }
    }
  },
  methods: {
    handleRegister() {
      // 验证密码是否一致
      if (this.form.password !== this.form.confirmPassword) {
        alert('两次输入的密码不一致！');
        return;
      }
      
      // 注册逻辑
      console.log('注册表单提交:', this.form);
      // 模拟注册成功
      alert('注册成功！');
      this.$router.push('/login');
    }
  }
}
</script>

<style scoped>
.register {
  padding: 60px 0;
  background-color: #f5f5f5;
  min-height: 80vh;
  display: flex;
  align-items: center;
}

.register-card {
  background-color: #fff;
  padding: 40px;
  border-radius: 8px;
  box-shadow: 0 5px 20px rgba(0, 0, 0, 0.1);
  max-width: 400px;
  margin: 0 auto;
  width: 100%;
}

.register-card h1 {
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

.register-btn {
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

.register-btn:hover {
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
  .register-card {
    padding: 30px;
    margin: 0 20px;
  }
  
  .register-card h1 {
    font-size: 24px;
  }
}
</style>
