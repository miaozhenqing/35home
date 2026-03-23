<template>
  <div class="pressure">
    <div class="container">
      <h1>压力管理</h1>
      <p>了解并应对35岁左右中年男性面临的各种压力</p>

      <!-- 压力类型 -->
      <section class="pressure-types">
        <h2>压力类型</h2>
        <div class="types-grid">
          <div class="type-item">
            <h3>事业压力</h3>
            <p>职业发展瓶颈、职场竞争、工作与生活平衡</p>
            <router-link to="/content" class="btn-link">了解更多</router-link>
          </div>
          <div class="type-item">
            <h3>家庭压力</h3>
            <p>婚姻关系、子女教育、父母养老、经济负担</p>
            <router-link to="/content" class="btn-link">了解更多</router-link>
          </div>
          <div class="type-item">
            <h3>情感压力</h3>
            <p>中年危机、社交焦虑、自我认同</p>
            <router-link to="/content" class="btn-link">了解更多</router-link>
          </div>
          <div class="type-item">
            <h3>健康压力</h3>
            <p>身体亚健康、中年健康管理、压力与健康的关系</p>
            <router-link to="/content" class="btn-link">了解更多</router-link>
          </div>
        </div>
      </section>

      <!-- 压力管理策略 -->
      <section class="management-strategies">
        <h2>压力管理策略</h2>
        <div class="strategies-grid">
          <div class="strategy-item">
            <div class="strategy-icon">
              <i class="icon-breathing"></i>
            </div>
            <h3>呼吸练习</h3>
            <p>通过深呼吸和冥想缓解压力，提高专注力</p>
          </div>
          <div class="strategy-item">
            <div class="strategy-icon">
              <i class="icon-exercise"></i>
            </div>
            <h3>运动锻炼</h3>
            <p>定期运动可以释放内啡肽，改善心情</p>
          </div>
          <div class="strategy-item">
            <div class="strategy-icon">
              <i class="icon-time"></i>
            </div>
            <h3>时间管理</h3>
            <p>合理安排时间，避免过度承诺和拖延</p>
          </div>
          <div class="strategy-item">
            <div class="strategy-icon">
              <i class="icon-social"></i>
            </div>
            <h3>社交支持</h3>
            <p>与朋友和家人保持联系，分享感受</p>
          </div>
        </div>
      </section>

      <!-- 专家建议 -->
      <section class="expert-advice">
        <h2>专家建议</h2>
        <div class="advice-item">
          <div class="advice-content">
            <h3>如何识别压力信号</h3>
            <p>压力会通过身体、情绪和行为表现出来。身体上可能出现头痛、失眠、疲劳等症状；情绪上可能表现为焦虑、易怒、抑郁；行为上可能出现暴饮暴食、吸烟饮酒增加等。了解这些信号，及时采取措施。</p>
          </div>
        </div>
        <div class="advice-item">
          <div class="advice-content">
            <h3>建立健康的生活方式</h3>
            <p>保持规律的作息时间，充足的睡眠，均衡的饮食，适量的运动。这些都是缓解压力的基础。同时，避免过度依赖咖啡因和酒精等刺激性物质。</p>
          </div>
        </div>
        <div class="advice-item">
          <div class="advice-content">
            <h3>学会说"不"</h3>
            <p>35岁的男性往往承担着多重责任，学会合理拒绝一些超出自己能力范围的请求，避免过度压力。同时，为自己留出一些私人时间，做自己喜欢的事情。</p>
          </div>
        </div>
      </section>

      <!-- 测评入口 -->
      <section class="test-entry">
        <div class="test-banner">
          <h2>想了解你的压力水平吗？</h2>
          <p>通过我们的专业测评工具，了解自己的压力状态，获取个性化的压力管理建议</p>
          <router-link to="/test" class="btn primary">开始测评</router-link>
        </div>
      </section>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Pressure'
}
</script>

<style scoped>
.pressure {
  padding: 40px 0;
}

.pressure h1 {
  font-size: 36px;
  margin-bottom: 20px;
  color: #2c3e50;
  text-align: center;
}

.pressure p {
  text-align: center;
  font-size: 18px;
  margin-bottom: 60px;
  color: #666;
}

/* 压力类型 */
.pressure-types {
  margin-bottom: 80px;
}

.pressure-types h2 {
  font-size: 28px;
  margin-bottom: 40px;
  color: #2c3e50;
  text-align: center;
}

.types-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 30px;
}

.type-item {
  background-color: #f9f9f9;
  padding: 30px;
  border-radius: 8px;
  transition: all 0.3s ease;
}

.type-item:hover {
  transform: translateY(-5px);
  box-shadow: 0 10px 20px rgba(0, 0, 0, 0.1);
}

.type-item h3 {
  font-size: 20px;
  margin-bottom: 15px;
  color: #3498db;
}

.type-item p {
  text-align: left;
  font-size: 16px;
  margin-bottom: 20px;
  color: #666;
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

/* 压力管理策略 */
.management-strategies {
  margin-bottom: 80px;
}

.management-strategies h2 {
  font-size: 28px;
  margin-bottom: 40px;
  color: #2c3e50;
  text-align: center;
}

.strategies-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 30px;
}

.strategy-item {
  text-align: center;
  padding: 30px;
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
  transition: all 0.3s ease;
}

.strategy-item:hover {
  transform: translateY(-5px);
  box-shadow: 0 10px 20px rgba(0, 0, 0, 0.1);
}

.strategy-icon {
  width: 60px;
  height: 60px;
  background-color: #3498db;
  border-radius: 50%;
  margin: 0 auto 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #fff;
  font-size: 24px;
}

.strategy-item h3 {
  font-size: 18px;
  margin-bottom: 15px;
  color: #2c3e50;
}

.strategy-item p {
  text-align: center;
  font-size: 16px;
  color: #666;
}

/* 专家建议 */
.expert-advice {
  margin-bottom: 80px;
}

.expert-advice h2 {
  font-size: 28px;
  margin-bottom: 40px;
  color: #2c3e50;
  text-align: center;
}

.advice-item {
  background-color: #f9f9f9;
  padding: 30px;
  border-radius: 8px;
  margin-bottom: 20px;
  transition: all 0.3s ease;
}

.advice-item:hover {
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
}

.advice-content h3 {
  font-size: 20px;
  margin-bottom: 15px;
  color: #3498db;
}

.advice-content p {
  text-align: left;
  font-size: 16px;
  color: #666;
}

/* 测评入口 */
.test-entry {
  margin-bottom: 40px;
}

.test-banner {
  background: linear-gradient(135deg, #3498db 0%, #2980b9 100%);
  color: #fff;
  padding: 60px;
  border-radius: 8px;
  text-align: center;
}

.test-banner h2 {
  font-size: 28px;
  margin-bottom: 20px;
}

.test-banner p {
  font-size: 18px;
  margin-bottom: 30px;
  color: rgba(255, 255, 255, 0.9);
}

.btn {
  display: inline-block;
  padding: 12px 30px;
  border-radius: 4px;
  text-decoration: none;
  font-weight: bold;
  transition: all 0.3s ease;
}

.btn.primary {
  background-color: #fff;
  color: #3498db;
}

.btn.primary:hover {
  background-color: rgba(255, 255, 255, 0.9);
  transform: translateY(-2px);
}

/* 响应式设计 */
@media (max-width: 768px) {
  .pressure h1 {
    font-size: 28px;
  }
  
  .pressure-types h2,
  .management-strategies h2,
  .expert-advice h2 {
    font-size: 24px;
  }
  
  .test-banner {
    padding: 40px 20px;
  }
  
  .test-banner h2 {
    font-size: 24px;
  }
}
</style>
