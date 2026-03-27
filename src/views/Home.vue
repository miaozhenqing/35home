<template>
  <div class="home">
    <!-- 用户信息和登出 -->
    <div class="user-info" v-if="isLoggedIn">
      <span class="welcome-text">欢迎，{{ user.username }}</span>
      <button class="logout-btn" @click="logout">登出</button>
    </div>
    
    <!-- 投票区 -->
    <section class="hero">
      <div class="container">
        <h1>平台</h1>
        <p>数据</p>
        <div class="poll-container">
          <h2>投票</h2>
          <p class="poll-question">状态是？</p>
          <div class="poll-vs">
            <div 
              class="poll-circle-container"
              :class="{ active: selectedOption === 'employed' }"
            >
              <div class="poll-circle" @click="vote('employed')">
                <span class="circle-icon">⚪</span>
              </div>
              <p class="poll-label">工作</p>
              <p class="poll-count">{{ pollResults.employed }}票</p>
            </div>
            <div class="poll-vs-text">vs</div>
            <div 
              class="poll-circle-container"
              :class="{ active: selectedOption === 'unemployed' }"
            >
              <div class="poll-circle" @click="vote('unemployed')">
                <span class="circle-icon">⚪</span>
              </div>
              <p class="poll-label">失业</p>
              <p class="poll-count">{{ pollResults.unemployed }}票</p>
            </div>
          </div>
          <p class="poll-note" v-if="voted">感谢你的投票！</p>
          
          <!-- 统计按钮 -->
          <div class="stats-section">
            <button class="stats-btn" @click="toggleStats">
              <span>失业统计</span>
              <span class="stats-arrow" :class="{ rotated: showStats }">&#9660;</span>
            </button>
            <div class="stats-dropdown" v-if="showStats">
              <h4>失业统计分析</h4>
              
              <!-- 城市统计 -->
              <div class="stats-item">
                <h5>按城市</h5>
                <div class="chart-container">
                  <canvas ref="cityChart"></canvas>
                </div>
              </div>
              
              <!-- 行业统计 -->
              <div class="stats-item">
                <h5>按行业</h5>
                <div class="chart-container">
                  <canvas ref="industryChart"></canvas>
                </div>
              </div>
              
              <!-- 职业统计 -->
              <div class="stats-item">
                <h5>按职业</h5>
                <div class="chart-container">
                  <canvas ref="occupationChart"></canvas>
                </div>
              </div>
              
              <!-- 年龄统计 -->
              <div class="stats-item">
                <h5>按年龄</h5>
                <div class="chart-container">
                  <canvas ref="ageChart"></canvas>
                </div>
              </div>
              
              <!-- 性别统计 -->
              <div class="stats-item">
                <h5>按性别</h5>
                <div class="chart-container">
                  <canvas ref="genderChart"></canvas>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
import Chart from 'chart.js/auto';
import { buildApiUrl, API_ENDPOINTS } from '../api/config';

export default {
  name: 'Home',
  data() {
    return {
      pollResults: {
        employed: 0,
        unemployed: 0
      },
      selectedOption: null,
      voted: false,
      showStats: false,
      charts: {},
      // 失业统计数据
      unemploymentData: {
        cities: {
          labels: [],
          data: []
        },
        industries: {
          labels: [],
          data: []
        },
        occupations: {
          labels: [],
          data: []
        },
        ages: {
          labels: [],
          data: []
        },
        genders: {
          labels: [],
          data: []
        }
      },
      // 登录状态
      isLoggedIn: false,
      // 用户信息
      user: null
    }
  },
  computed: {
    totalVotes() {
      return this.pollResults.employed + this.pollResults.unemployed;
    },
    employedPercentage() {
      if (this.totalVotes === 0) return 0;
      return Math.round((this.pollResults.employed / this.totalVotes) * 100);
    },
    unemployedPercentage() {
      if (this.totalVotes === 0) return 0;
      return Math.round((this.pollResults.unemployed / this.totalVotes) * 100);
    }
  },
  mounted() {
    // 组件挂载时获取统计数据
    this.getVoteStats();
    
    // 检查登录状态
    this.checkLoginStatus();
  },
  methods: {
    toggleStats() {
      this.showStats = !this.showStats;
      if (this.showStats) {
        this.$nextTick(() => {
          this.initCharts();
        });
      }
    },
    // 检查登录状态
    checkLoginStatus() {
      // 实际项目中应该从localStorage或store中获取
      const user = localStorage.getItem('user');
      if (user) {
        this.isLoggedIn = true;
        this.user = JSON.parse(user);
      }
    },
    
    // 登出功能
    logout() {
      // 清除localStorage中的用户信息和token
      localStorage.removeItem('user');
      localStorage.removeItem('token');
      // 更新登录状态
      this.isLoggedIn = false;
      this.user = null;
      // 跳转到登录页面
      this.$router.push('/login');
    },
    // 获取投票统计数据
    getVoteStats() {
      const token = localStorage.getItem('token');
      const headers = {
        'Content-Type': 'application/json'
      };
      
      if (token) {
        headers['Authorization'] = `Bearer ${token}`;
      }
      
      fetch(buildApiUrl(API_ENDPOINTS.VOTE_STATS), {
        headers: headers
      })
        .then(response => response.json())
        .then(data => {
          if (data.code === 0) {
            const stats = data.respBody;
            this.pollResults.employed = stats.employedCount;
            this.pollResults.unemployed = stats.unemployedCount;
            
            // 处理城市统计数据
            this.unemploymentData.cities.labels = Object.keys(stats.cityStats);
            this.unemploymentData.cities.data = Object.values(stats.cityStats);
            
            // 处理行业统计数据
            this.unemploymentData.industries.labels = Object.keys(stats.industryStats);
            this.unemploymentData.industries.data = Object.values(stats.industryStats);
            
            // 处理职业统计数据
            this.unemploymentData.occupations.labels = Object.keys(stats.occupationStats);
            this.unemploymentData.occupations.data = Object.values(stats.occupationStats);
            
            // 处理年龄统计数据
            this.unemploymentData.ages.labels = Object.keys(stats.ageStats);
            this.unemploymentData.ages.data = Object.values(stats.ageStats);
            
            // 处理性别统计数据
            this.unemploymentData.genders.labels = Object.keys(stats.genderStats);
            this.unemploymentData.genders.data = Object.values(stats.genderStats);
            
            // 如果统计面板是打开的，重新初始化图表
            if (this.showStats) {
              this.$nextTick(() => {
                this.initCharts();
              });
            }
          }
        })
        .catch(error => {
          console.error('获取统计数据失败:', error);
        });
    },
    // 提交投票
    vote(option) {
      if (this.voted) return;
      
      // 检查登录状态
      if (!this.isLoggedIn) {
        alert('请登录后再投票');
        return;
      }
      
      // 准备投票数据
      const voteData = {
        status: option,
        city: '北京', // 实际项目中应该从用户信息或表单中获取
        industry: '互联网', // 实际项目中应该从用户信息或表单中获取
        occupation: '技术岗', // 实际项目中应该从用户信息或表单中获取
        age: 35, // 实际项目中应该从用户信息或表单中获取
        gender: 'male' // 实际项目中应该从用户信息或表单中获取
      };
      
      // 提交投票
      const token = localStorage.getItem('token');
      fetch(buildApiUrl(API_ENDPOINTS.VOTE_SUBMIT), {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`
        },
        body: JSON.stringify(voteData)
      })
      .then(response => response.json())
      .then(data => {
        if (data.code === 0) {
          // 投票成功，更新本地数据
          this.selectedOption = option;
          this.pollResults[option]++;
          this.voted = true;
          // 重新获取统计数据
          this.getVoteStats();
        } else {
          // 投票失败，显示错误信息
          alert(data.msg);
        }
      })
      .catch(error => {
        console.error('投票失败:', error);
        alert('投票失败，请稍后重试');
      });
    },
    initCharts() {
      // 销毁旧图表
      Object.values(this.charts).forEach(chart => {
        if (chart) chart.destroy();
      });
      
      // 初始化城市图表
      if (this.$refs.cityChart) {
        this.charts.cityChart = new Chart(this.$refs.cityChart, {
          type: 'bar',
          data: {
            labels: this.unemploymentData.cities.labels,
            datasets: [{
              label: '失业人数',
              data: this.unemploymentData.cities.data,
              backgroundColor: 'rgba(54, 162, 235, 0.6)',
              borderColor: 'rgba(54, 162, 235, 1)',
              borderWidth: 1
            }]
          },
          options: {
            responsive: true,
            maintainAspectRatio: false,
            scales: {
              y: {
                beginAtZero: true,
                ticks: {
                  color: '#ddd'
                },
                grid: {
                  color: 'rgba(255, 255, 255, 0.1)'
                }
              },
              x: {
                ticks: {
                  color: '#ddd'
                },
                grid: {
                  color: 'rgba(255, 255, 255, 0.1)'
                }
              }
            },
            plugins: {
              legend: {
                labels: {
                  color: '#ddd'
                }
              }
            }
          }
        });
      }
      
      // 初始化行业图表
      if (this.$refs.industryChart) {
        this.charts.industryChart = new Chart(this.$refs.industryChart, {
          type: 'pie',
          data: {
            labels: this.unemploymentData.industries.labels,
            datasets: [{
              data: this.unemploymentData.industries.data,
              backgroundColor: [
                'rgba(255, 99, 132, 0.6)',
                'rgba(54, 162, 235, 0.6)',
                'rgba(255, 206, 86, 0.6)',
                'rgba(75, 192, 192, 0.6)',
                'rgba(153, 102, 255, 0.6)'
              ],
              borderColor: [
                'rgba(255, 99, 132, 1)',
                'rgba(54, 162, 235, 1)',
                'rgba(255, 206, 86, 1)',
                'rgba(75, 192, 192, 1)',
                'rgba(153, 102, 255, 1)'
              ],
              borderWidth: 1
            }]
          },
          options: {
            responsive: true,
            maintainAspectRatio: false,
            plugins: {
              legend: {
                labels: {
                  color: '#ddd'
                }
              }
            }
          }
        });
      }
      
      // 初始化职业图表
      if (this.$refs.occupationChart) {
        this.charts.occupationChart = new Chart(this.$refs.occupationChart, {
          type: 'doughnut',
          data: {
            labels: this.unemploymentData.occupations.labels,
            datasets: [{
              data: this.unemploymentData.occupations.data,
              backgroundColor: [
                'rgba(255, 99, 132, 0.6)',
                'rgba(54, 162, 235, 0.6)',
                'rgba(255, 206, 86, 0.6)',
                'rgba(75, 192, 192, 0.6)',
                'rgba(153, 102, 255, 0.6)'
              ],
              borderColor: [
                'rgba(255, 99, 132, 1)',
                'rgba(54, 162, 235, 1)',
                'rgba(255, 206, 86, 1)',
                'rgba(75, 192, 192, 1)',
                'rgba(153, 102, 255, 1)'
              ],
              borderWidth: 1
            }]
          },
          options: {
            responsive: true,
            maintainAspectRatio: false,
            plugins: {
              legend: {
                labels: {
                  color: '#ddd'
                }
              }
            }
          }
        });
      }
      
      // 初始化年龄图表
      if (this.$refs.ageChart) {
        this.charts.ageChart = new Chart(this.$refs.ageChart, {
          type: 'bar',
          data: {
            labels: this.unemploymentData.ages.labels,
            datasets: [{
              label: '失业人数',
              data: this.unemploymentData.ages.data,
              backgroundColor: 'rgba(255, 206, 86, 0.6)',
              borderColor: 'rgba(255, 206, 86, 1)',
              borderWidth: 1
            }]
          },
          options: {
            responsive: true,
            maintainAspectRatio: false,
            scales: {
              y: {
                beginAtZero: true,
                ticks: {
                  color: '#ddd'
                },
                grid: {
                  color: 'rgba(255, 255, 255, 0.1)'
                }
              },
              x: {
                ticks: {
                  color: '#ddd'
                },
                grid: {
                  color: 'rgba(255, 255, 255, 0.1)'
                }
              }
            },
            plugins: {
              legend: {
                labels: {
                  color: '#ddd'
                }
              }
            }
          }
        });
      }
      
      // 初始化性别图表
      if (this.$refs.genderChart) {
        this.charts.genderChart = new Chart(this.$refs.genderChart, {
          type: 'pie',
          data: {
            labels: this.unemploymentData.genders.labels,
            datasets: [{
              data: this.unemploymentData.genders.data,
              backgroundColor: [
                'rgba(54, 162, 235, 0.6)',
                'rgba(255, 99, 132, 0.6)'
              ],
              borderColor: [
                'rgba(54, 162, 235, 1)',
                'rgba(255, 99, 132, 1)'
              ],
              borderWidth: 1
            }]
          },
          options: {
            responsive: true,
            maintainAspectRatio: false,
            plugins: {
              legend: {
                labels: {
                  color: '#ddd'
                }
              }
            }
          }
        });
      }
    }
  }
}
</script>

<style scoped>
/* 用户信息和登出 */
.user-info {
  background-color: #f8f9fa;
  padding: 10px 20px;
  text-align: right;
  border-bottom: 1px solid #e9ecef;
}

.welcome-text {
  margin-right: 20px;
  font-weight: bold;
  color: #2c3e50;
}

.logout-btn {
  background-color: #e74c3c;
  color: white;
  border: none;
  padding: 8px 16px;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
  font-weight: bold;
  transition: background-color 0.3s ease;
}

.logout-btn:hover {
  background-color: #c0392b;
}

/* 英雄区 */
.hero {
  background: linear-gradient(135deg, #2c3e50 0%, #34495e 100%);
  color: #fff;
  padding: 60px 0;
  text-align: center;
}

.hero h1 {
  font-size: 36px;
  margin-bottom: 15px;
  font-weight: bold;
}

.hero p {
  font-size: 18px;
  margin-bottom: 30px;
  opacity: 0.9;
}

/* 投票区 */
.poll-container {
  background-color: rgba(255, 255, 255, 0.1);
  padding: 30px;
  border-radius: 8px;
  max-width: 600px;
  margin: 0 auto;
}

.poll-container h2 {
  font-size: 24px;
  margin-bottom: 15px;
  color: #3498db;
}

.poll-question {
  font-size: 18px;
  margin-bottom: 30px;
  color: #fff;
}

/* VS投票样式 */
.poll-vs {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 40px;
  margin-bottom: 30px;
}

.poll-circle-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  transition: all 0.3s ease;
}

.poll-circle-container.active {
  transform: scale(1.05);
}

.poll-circle {
  width: 120px;
  height: 120px;
  border: 3px solid #3498db;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.3s ease;
  background-color: rgba(52, 152, 219, 0.1);
}

.poll-circle:hover {
  background-color: rgba(52, 152, 219, 0.2);
  transform: scale(1.05);
}

.poll-circle-container.active .poll-circle {
  background-color: rgba(52, 152, 219, 0.3);
  border-color: #2980b9;
}

.circle-icon {
  font-size: 48px;
  color: #3498db;
  transition: all 0.3s ease;
}

.poll-circle-container.active .circle-icon {
  color: #2980b9;
  transform: scale(1.1);
}

.poll-label {
  font-size: 18px;
  font-weight: bold;
  margin-top: 15px;
  color: #fff;
}

.poll-count {
  font-size: 16px;
  color: #3498db;
  font-weight: bold;
  margin-top: 5px;
}

.poll-vs-text {
  font-size: 24px;
  font-weight: bold;
  color: #fff;
}

.poll-note {
  font-size: 14px;
  color: #3498db;
  font-weight: bold;
  text-align: center;
  margin-bottom: 20px;
}

/* 统计部分 */
.stats-section {
  margin-top: 20px;
}

.stats-btn {
  background-color: rgba(52, 152, 219, 0.2);
  color: #3498db;
  border: 2px solid #3498db;
  border-radius: 4px;
  padding: 10px 20px;
  font-size: 16px;
  font-weight: bold;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  gap: 10px;
  margin: 0 auto;
}

.stats-btn:hover {
  background-color: rgba(52, 152, 219, 0.3);
  transform: translateY(-2px);
}

.stats-arrow {
  transition: transform 0.3s ease;
}

.stats-arrow.rotated {
  transform: rotate(180deg);
}

.stats-dropdown {
  background-color: rgba(255, 255, 255, 0.05);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 4px;
  padding: 20px;
  margin-top: 10px;
  animation: slideDown 0.3s ease;
}

@keyframes slideDown {
  from {
    opacity: 0;
    transform: translateY(-10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.stats-dropdown h4 {
  font-size: 18px;
  margin-bottom: 20px;
  color: #3498db;
  text-align: center;
}

.stats-item {
  margin-bottom: 30px;
  padding: 15px;
  background-color: rgba(255, 255, 255, 0.05);
  border-radius: 4px;
}

.stats-item h5 {
  font-size: 16px;
  margin-bottom: 15px;
  color: #fff;
  text-align: center;
}

/* 图表容器 */
.chart-container {
  position: relative;
  height: 300px;
  width: 100%;
  margin: 0 auto;
}


.btn-link {
  color: #3498db;
  text-decoration: none;
  font-weight: bold;
  transition: color 0.3s ease;
}

.btn-link:hover {
  color: #2980b9;
}

/* 内容预览 */
.content-preview {
  padding: 80px 0;
  background-color: #f5f5f5;
}

.content-preview h2 {
  text-align: center;
  font-size: 36px;
  margin-bottom: 60px;
  color: #2c3e50;
}

.content-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 30px;
}

.content-item {
  background-color: #fff;
  border-radius: 8px;
  overflow: hidden;
  transition: all 0.3s ease;
}

.content-item:hover {
  transform: translateY(-5px);
  box-shadow: 0 10px 20px rgba(0, 0, 0, 0.1);
}

.content-img {
  height: 200px;
  background-color: #e0e0e0;
  margin-bottom: 20px;
}

.content-item h3 {
  padding: 0 20px;
  font-size: 18px;
  margin-bottom: 10px;
  color: #2c3e50;
}

.content-item p {
  padding: 0 20px;
  margin-bottom: 20px;
  color: #666;
}

.content-item .btn-link {
  padding: 0 20px 20px;
  display: block;
  text-align: left;
}

/* 会员优势 */
.membership {
  padding: 80px 0;
  background-color: #2c3e50;
  color: #fff;
  text-align: center;
}

.membership h2 {
  font-size: 36px;
  margin-bottom: 20px;
}

.membership p {
  font-size: 18px;
  margin-bottom: 60px;
  opacity: 0.9;
}

.membership-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 30px;
  margin-bottom: 40px;
}

.membership-item {
  padding: 30px;
  background-color: rgba(255, 255, 255, 0.1);
  border-radius: 8px;
}

.membership-item h3 {
  font-size: 20px;
  margin-bottom: 15px;
}

.membership-item p {
  font-size: 16px;
  margin-bottom: 0;
  opacity: 0.8;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .hero h1 {
    font-size: 36px;
  }
  
  .hero p {
    font-size: 18px;
  }
  
  .hero-buttons {
    flex-direction: column;
    align-items: center;
  }
  
  .btn {
    width: 200px;
    text-align: center;
  }
  
  .features h2,
  .content-preview h2,
  .membership h2 {
    font-size: 28px;
  }
}
</style>
