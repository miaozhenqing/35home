<template>
  <div class="home">
    <!-- 投票区 -->
    <section class="hero">
      <div class="container">
        <p>轻松交流，分享生活，一起面对中年的挑战</p>
        <div class="poll-container">
          <h2>当前热门话题</h2>
          <p class="poll-question">你目前的就业状态是？</p>
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
              <span>统计</span>
              <span class="stats-arrow" :class="{ rotated: showStats }">▼</span>
            </button>
            <div class="stats-dropdown" v-if="showStats">
              <h4>失业统计</h4>
              <div class="stats-grid">
                <div class="stats-item">
                  <h5>按城市</h5>
                  <ul>
                    <li>北京: 12人</li>
                    <li>上海: 18人</li>
                    <li>广州: 8人</li>
                    <li>深圳: 10人</li>
                    <li>其他: 5人</li>
                  </ul>
                </div>
                <div class="stats-item">
                  <h5>按行业</h5>
                  <ul>
                    <li>互联网: 25人</li>
                    <li>金融: 8人</li>
                    <li>制造业: 6人</li>
                    <li>教育: 3人</li>
                    <li>其他: 3人</li>
                  </ul>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- 内容预览 -->
    <section class="content-preview">
      <div class="container">
        <h2>热门内容</h2>
        <div class="content-grid">
          <div class="content-item">
            <div class="content-img"></div>
            <h3>如何平衡工作与生活</h3>
            <p>35岁是事业的关键期，如何在追求成功的同时保持生活的平衡？</p>
            <router-link to="/content" class="btn-link">阅读更多</router-link>
          </div>
          <div class="content-item">
            <div class="content-img"></div>
            <h3>中年情感管理</h3>
            <p>面对家庭、婚姻和社交关系的挑战，如何保持健康的情感状态？</p>
            <router-link to="/content" class="btn-link">阅读更多</router-link>
          </div>
          <div class="content-item">
            <div class="content-img"></div>
            <h3>健康管理策略</h3>
            <p>35岁后，如何关注自己的身体健康，预防亚健康状态？</p>
            <router-link to="/content" class="btn-link">阅读更多</router-link>
          </div>
        </div>
      </div>
    </section>


  </div>
</template>

<script>
export default {
  name: 'Home',
  data() {
    return {
      pollResults: {
        employed: 128,
        unemployed: 45
      },
      selectedOption: null,
      voted: false,
      showStats: false
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
  methods: {
    toggleStats() {
      this.showStats = !this.showStats;
    },
    vote(option) {
      if (this.voted) return;
      
      this.selectedOption = option;
      this.pollResults[option]++;
      this.voted = true;
      
      // 模拟登录状态检查
      const isLoggedIn = false; // 实际项目中应该从用户状态中获取
      if (!isLoggedIn) {
        alert('请登录后再投票');
        this.selectedOption = null;
        this.pollResults[option]--;
        this.voted = false;
      }
    }
  }
}
</script>

<style scoped>
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
  margin-bottom: 15px;
  color: #3498db;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 20px;
}

.stats-item h5 {
  font-size: 16px;
  margin-bottom: 10px;
  color: #fff;
}

.stats-item ul {
  list-style: none;
  padding: 0;
}

.stats-item li {
  font-size: 14px;
  margin-bottom: 8px;
  color: #ddd;
  display: flex;
  justify-content: space-between;
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
