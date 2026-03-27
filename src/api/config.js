// API配置文件
export const API_BASE_URL = 'http://localhost:9001/api';

// API端点配置
export const API_ENDPOINTS = {
  // 用户相关
  USER_LOGIN: '/user/login',
  USER_REGISTER: '/user/register',
  
  // 投票相关
  VOTE_STATS: '/vote/stats',
  VOTE_SUBMIT: '/vote/submit',
  
  // 压力测试相关
  TEST_QUESTIONS: '/test/questions',
  TEST_SUBMIT: '/test/submit',
  
  // 内容相关
  ARTICLES: '/content/articles',
  ARTICLE_BY_ID: '/content/article/:id'
};

// 构建完整的API URL
export const buildApiUrl = (endpoint) => {
  return `${API_BASE_URL}${endpoint}`;
};