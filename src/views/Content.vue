<template>
  <div class="content">
    <div class="container">
      <h1>内容中心</h1>
      <p>探索压力管理、生活品质提升等相关内容</p>

      <!-- 内容分类 -->
      <div class="content-categories">
        <button class="category-btn active">全部</button>
        <button class="category-btn">文章</button>
        <button class="category-btn">视频</button>
        <button class="category-btn">音频</button>
        <button class="category-btn">工具</button>
      </div>

      <!-- 内容列表 -->
      <div class="content-list">
        <div class="content-item">
          <div class="content-img"></div>
          <div class="content-info">
            <div class="content-meta">
              <span class="content-type">文章</span>
              <span class="content-date">2026-03-20</span>
            </div>
            <h3>如何平衡工作与生活</h3>
            <p>35岁是事业的关键期，如何在追求成功的同时保持生活的平衡？本文将为你提供实用的建议和方法。</p>
            <a href="#" class="content-link">阅读全文</a>
          </div>
        </div>
        <div class="content-item">
          <div class="content-img"></div>
          <div class="content-info">
            <div class="content-meta">
              <span class="content-type">视频</span>
              <span class="content-date">2026-03-18</span>
            </div>
            <h3>中年情感管理技巧</h3>
            <p>面对家庭、婚姻和社交关系的挑战，如何保持健康的情感状态？专家为你解读。</p>
            <a href="#" class="content-link">观看视频</a>
          </div>
        </div>
        <div class="content-item">
          <div class="content-img"></div>
          <div class="content-info">
            <div class="content-meta">
              <span class="content-type">文章</span>
              <span class="content-date">2026-03-15</span>
            </div>
            <h3>健康管理策略</h3>
            <p>35岁后，如何关注自己的身体健康，预防亚健康状态？本文为你提供全面的健康管理方案。</p>
            <a href="#" class="content-link">阅读全文</a>
          </div>
        </div>
        <div class="content-item">
          <div class="content-img"></div>
          <div class="content-info">
            <div class="content-meta">
              <span class="content-type">音频</span>
              <span class="content-date">2026-03-10</span>
            </div>
            <h3>冥想指导：缓解压力</h3>
            <p>通过冥想练习，缓解工作和生活压力，提高专注力和幸福感。</p>
            <a href="#" class="content-link">收听音频</a>
          </div>
        </div>
      </div>

      <!-- 分页 -->
      <div class="pagination">
        <button class="page-btn">上一页</button>
        <button class="page-btn active">1</button>
        <button class="page-btn">2</button>
        <button class="page-btn">3</button>
        <button class="page-btn">下一页</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Content'
}
</script>

<style scoped>
.content {
  padding: 40px 0;
}

.content h1 {
  font-size: 36px;
  margin-bottom: 20px;
  color: #2c3e50;
  text-align: center;
}

.content p {
  text-align: center;
  font-size: 18px;
  margin-bottom: 60px;
  color: #666;
}

/* 内容分类 */
.content-categories {
  display: flex;
  justify-content: center;
  gap: 20px;
  margin-bottom: 40px;
}

.category-btn {
  padding: 10px 20px;
  border: 1px solid #3498db;
  background-color: #fff;
  color: #3498db;
  border-radius: 20px;
  cursor: pointer;
  transition: all 0.3s ease;
}

.category-btn:hover {
  background-color: #3498db;
  color: #fff;
}

.category-btn.active {
  background-color: #3498db;
  color: #fff;
}

/* 内容列表 */
.content-list {
  margin-bottom: 40px;
}

.content-item {
  display: flex;
  background-color: #fff;
  border-radius: 8px;
  overflow: hidden;
  margin-bottom: 20px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
  transition: all 0.3s ease;
}

.content-item:hover {
  transform: translateY(-5px);
  box-shadow: 0 10px 20px rgba(0, 0, 0, 0.1);
}

.content-img {
  width: 200px;
  height: 150px;
  background-color: #e0e0e0;
}

.content-info {
  flex: 1;
  padding: 20px;
}

.content-meta {
  display: flex;
  gap: 20px;
  margin-bottom: 10px;
  font-size: 14px;
  color: #999;
}

.content-type {
  background-color: #f0f0f0;
  padding: 2px 10px;
  border-radius: 10px;
}

.content-info h3 {
  font-size: 20px;
  margin-bottom: 10px;
  color: #2c3e50;
}

.content-info p {
  text-align: left;
  font-size: 16px;
  margin-bottom: 15px;
  color: #666;
}

.content-link {
  color: #3498db;
  text-decoration: none;
  font-weight: bold;
  transition: color 0.3s ease;
}

.content-link:hover {
  color: #2980b9;
}

/* 分页 */
.pagination {
  display: flex;
  justify-content: center;
  gap: 10px;
}

.page-btn {
  padding: 8px 16px;
  border: 1px solid #ddd;
  background-color: #fff;
  color: #333;
  border-radius: 4px;
  cursor: pointer;
  transition: all 0.3s ease;
}

.page-btn:hover {
  background-color: #f0f0f0;
}

.page-btn.active {
  background-color: #3498db;
  color: #fff;
  border-color: #3498db;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .content h1 {
    font-size: 28px;
  }
  
  .content-categories {
    flex-wrap: wrap;
  }
  
  .content-item {
    flex-direction: column;
  }
  
  .content-img {
    width: 100%;
    height: 200px;
  }
  
  .content-info {
    padding: 20px;
  }
}
</style>
