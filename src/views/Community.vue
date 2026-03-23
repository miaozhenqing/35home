<template>
  <div class="community">
    <div class="container">
      <h1>社区</h1>
      <p>与志同道合的朋友交流，分享经验，互相支持</p>

      <!-- 社区导航 -->
      <div class="community-nav">
        <button class="nav-btn active">全部</button>
        <button class="nav-btn">热门</button>
        <button class="nav-btn">最新</button>
        <button class="nav-btn">我的话题</button>
      </div>

      <!-- 发布话题 -->
      <div class="post-topic">
        <textarea placeholder="分享你的故事或问题..."></textarea>
        <div class="post-actions">
          <button class="post-btn">发布话题</button>
        </div>
      </div>

      <!-- 话题列表 -->
      <div class="topic-list">
        <div class="topic-item">
          <div class="topic-author">
            <div class="author-avatar"></div>
            <div class="author-info">
              <h4>用户123</h4>
              <p>2026-03-20 14:30</p>
            </div>
          </div>
          <div class="topic-content">
            <h3>工作压力大，如何调整心态？</h3>
            <p>最近项目 deadline 临近，每天加班到很晚，感觉压力很大，睡眠质量也不好。想问问大家都是如何调整心态的？</p>
          </div>
          <div class="topic-actions">
            <button class="action-btn">
              <i class="icon-like"></i>
              <span>28</span>
            </button>
            <button class="action-btn">
              <i class="icon-comment"></i>
              <span>15</span>
            </button>
            <button class="action-btn">
              <i class="icon-share"></i>
              <span>分享</span>
            </button>
          </div>
          <!-- 评论 -->
          <div class="topic-comments">
            <div class="comment-item">
              <div class="comment-author">
                <div class="author-avatar small"></div>
                <div class="author-info">
                  <h5>用户456</h5>
                  <p>2026-03-20 15:00</p>
                </div>
              </div>
              <div class="comment-content">
                <p>我之前也有类似的情况，建议试试冥想和深呼吸，每天抽出10分钟时间，会有很大帮助。</p>
              </div>
            </div>
            <div class="comment-item">
              <div class="comment-author">
                <div class="author-avatar small"></div>
                <div class="author-info">
                  <h5>用户789</h5>
                  <p>2026-03-20 16:00</p>
                </div>
              </div>
              <div class="comment-content">
                <p>工作再忙也要保证充足的睡眠，建议制定一个合理的作息时间表，提高工作效率，避免拖延。</p>
              </div>
            </div>
            <div class="comment-input">
              <input type="text" placeholder="写下你的评论...">
              <button class="comment-btn">发表</button>
            </div>
          </div>
        </div>
        <div class="topic-item">
          <div class="topic-author">
            <div class="author-avatar"></div>
            <div class="author-info">
              <h4>用户ABC</h4>
              <p>2026-03-19 10:15</p>
            </div>
          </div>
          <div class="topic-content">
            <h3>中年危机，感觉人生没有方向</h3>
            <p>35岁了，事业上遇到瓶颈，家庭压力也大，感觉人生没有方向，不知道该如何突破。</p>
          </div>
          <div class="topic-actions">
            <button class="action-btn">
              <i class="icon-like"></i>
              <span>45</span>
            </button>
            <button class="action-btn">
              <i class="icon-comment"></i>
              <span>23</span>
            </button>
            <button class="action-btn">
              <i class="icon-share"></i>
              <span>分享</span>
            </button>
          </div>
          <!-- 评论 -->
          <div class="topic-comments">
            <div class="comment-item">
              <div class="comment-author">
                <div class="author-avatar small"></div>
                <div class="author-info">
                  <h5>用户DEF</h5>
                  <p>2026-03-19 11:00</p>
                </div>
              </div>
              <div class="comment-content">
                <p>我也有过类似的感受，建议尝试一些新的兴趣爱好，或者参加一些行业活动，拓展人脉，可能会找到新的方向。</p>
              </div>
            </div>
            <div class="comment-input">
              <input type="text" placeholder="写下你的评论...">
              <button class="comment-btn">发表</button>
            </div>
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
  name: 'Community'
}
</script>

<style scoped>
.community {
  padding: 40px 0;
}

.community h1 {
  font-size: 36px;
  margin-bottom: 20px;
  color: #2c3e50;
  text-align: center;
}

.community p {
  text-align: center;
  font-size: 18px;
  margin-bottom: 60px;
  color: #666;
}

/* 社区导航 */
.community-nav {
  display: flex;
  gap: 20px;
  margin-bottom: 30px;
}

.nav-btn {
  padding: 10px 20px;
  border: 1px solid #ddd;
  background-color: #fff;
  color: #333;
  border-radius: 20px;
  cursor: pointer;
  transition: all 0.3s ease;
}

.nav-btn:hover {
  background-color: #f0f0f0;
}

.nav-btn.active {
  background-color: #3498db;
  color: #fff;
  border-color: #3498db;
}

/* 发布话题 */
.post-topic {
  background-color: #f9f9f9;
  padding: 20px;
  border-radius: 8px;
  margin-bottom: 30px;
}

.post-topic textarea {
  width: 100%;
  height: 100px;
  padding: 15px;
  border: 1px solid #ddd;
  border-radius: 4px;
  resize: none;
  font-size: 16px;
  margin-bottom: 15px;
}

.post-actions {
  display: flex;
  justify-content: flex-end;
}

.post-btn {
  padding: 10px 30px;
  background-color: #3498db;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: all 0.3s ease;
}

.post-btn:hover {
  background-color: #2980b9;
}

/* 话题列表 */
.topic-list {
  margin-bottom: 40px;
}

.topic-item {
  background-color: #fff;
  border-radius: 8px;
  padding: 20px;
  margin-bottom: 20px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
  transition: all 0.3s ease;
}

.topic-item:hover {
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
}

.topic-author {
  display: flex;
  gap: 15px;
  margin-bottom: 15px;
}

.author-avatar {
  width: 50px;
  height: 50px;
  background-color: #e0e0e0;
  border-radius: 50%;
}

.author-avatar.small {
  width: 30px;
  height: 30px;
}

.author-info h4 {
  font-size: 16px;
  margin-bottom: 5px;
  color: #2c3e50;
}

.author-info h5 {
  font-size: 14px;
  margin-bottom: 3px;
  color: #2c3e50;
}

.author-info p {
  text-align: left;
  font-size: 14px;
  margin-bottom: 0;
  color: #999;
}

.topic-content h3 {
  font-size: 18px;
  margin-bottom: 10px;
  color: #2c3e50;
}

.topic-content p {
  text-align: left;
  font-size: 16px;
  margin-bottom: 15px;
  color: #666;
}

.topic-actions {
  display: flex;
  gap: 30px;
  margin-bottom: 20px;
  padding-bottom: 20px;
  border-bottom: 1px solid #f0f0f0;
}

.action-btn {
  display: flex;
  align-items: center;
  gap: 5px;
  background: none;
  border: none;
  color: #666;
  cursor: pointer;
  transition: color 0.3s ease;
}

.action-btn:hover {
  color: #3498db;
}

/* 评论 */
.topic-comments {
  margin-top: 20px;
}

.comment-item {
  display: flex;
  gap: 15px;
  margin-bottom: 15px;
  padding: 15px;
  background-color: #f9f9f9;
  border-radius: 4px;
}

.comment-content p {
  text-align: left;
  font-size: 14px;
  margin-bottom: 0;
  color: #666;
}

.comment-input {
  display: flex;
  gap: 10px;
  margin-top: 15px;
}

.comment-input input {
  flex: 1;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 14px;
}

.comment-btn {
  padding: 0 20px;
  background-color: #3498db;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: all 0.3s ease;
}

.comment-btn:hover {
  background-color: #2980b9;
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
  .community h1 {
    font-size: 28px;
  }
  
  .community-nav {
    flex-wrap: wrap;
  }
  
  .topic-actions {
    gap: 20px;
  }
}
</style>
