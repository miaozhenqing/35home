<template>
  <div class="test">
    <div class="container">
      <h1>压力测评</h1>
      <p>通过专业的测评工具，了解你的压力水平</p>

      <!-- 测评说明 -->
      <div class="test-intro">
        <h2>测评说明</h2>
        <p>本测评包含20个问题，每个问题有5个选项，请根据你的实际情况选择最符合的答案。测评完成后，我们会为你生成详细的压力分析报告。</p>
        <button class="start-btn" @click="startTest">开始测评</button>
      </div>

      <!-- 测评问卷 -->
      <div class="test-questions" v-if="showQuestions">
        <div class="question-progress">
          <span>问题 {{ currentQuestion + 1 }} / {{ questions.length }}</span>
          <div class="progress-bar">
            <div class="progress-fill" :style="{ width: ((currentQuestion + 1) / questions.length) * 100 + '%' }"></div>
          </div>
        </div>

        <div class="question-item">
          <h3>{{ questions[currentQuestion].question }}</h3>
          <div class="options">
            <div 
              v-for="(option, index) in questions[currentQuestion].options" 
              :key="index"
              class="option-item"
              :class="{ active: selectedOption === index }"
              @click="selectOption(index)"
            >
              <span class="option-text">{{ option }}</span>
            </div>
          </div>
        </div>

        <div class="question-actions">
          <button class="prev-btn" @click="prevQuestion" :disabled="currentQuestion === 0">上一题</button>
          <button class="next-btn" @click="nextQuestion" :disabled="selectedOption === -1">
            {{ currentQuestion === questions.length - 1 ? '提交' : '下一题' }}
          </button>
        </div>
      </div>

      <!-- 测评结果 -->
      <div class="test-result" v-if="showResult">
        <h2>测评结果</h2>
        <div class="result-card">
          <div class="result-score">
            <h3>你的压力指数</h3>
            <div class="score-value">{{ totalScore }}</div>
            <div class="score-level" :class="scoreLevelClass">{{ scoreLevel }}</div>
          </div>
          <div class="result-analysis">
            <h4>分析报告</h4>
            <p>{{ analysis }}</p>
          </div>
          <div class="result-suggestions">
            <h4>建议</h4>
            <ul>
              <li v-for="(suggestion, index) in suggestions" :key="index">{{ suggestion }}</li>
            </ul>
          </div>
        </div>
        <button class="restart-btn" @click="restartTest">重新测评</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Test',
  data() {
    return {
      showQuestions: false,
      showResult: false,
      currentQuestion: 0,
      selectedOption: -1,
      answers: [],
      questions: [
        {
          question: '你是否经常感到紧张或焦虑？',
          options: ['从不', '偶尔', '有时', '经常', '总是']
        },
        {
          question: '你是否经常感到疲劳或精力不足？',
          options: ['从不', '偶尔', '有时', '经常', '总是']
        },
        {
          question: '你是否经常失眠或睡眠质量差？',
          options: ['从不', '偶尔', '有时', '经常', '总是']
        },
        {
          question: '你是否经常感到工作压力过大？',
          options: ['从不', '偶尔', '有时', '经常', '总是']
        },
        {
          question: '你是否经常感到家庭压力过大？',
          options: ['从不', '偶尔', '有时', '经常', '总是']
        },
        {
          question: '你是否经常感到情绪低落或抑郁？',
          options: ['从不', '偶尔', '有时', '经常', '总是']
        },
        {
          question: '你是否经常感到身体不适（如头痛、胃痛等）？',
          options: ['从不', '偶尔', '有时', '经常', '总是']
        },
        {
          question: '你是否经常感到时间不够用？',
          options: ['从不', '偶尔', '有时', '经常', '总是']
        },
        {
          question: '你是否经常感到无法应对生活中的挑战？',
          options: ['从不', '偶尔', '有时', '经常', '总是']
        },
        {
          question: '你是否经常感到社交压力？',
          options: ['从不', '偶尔', '有时', '经常', '总是']
        },
        {
          question: '你是否经常感到经济压力？',
          options: ['从不', '偶尔', '有时', '经常', '总是']
        },
        {
          question: '你是否经常感到职业发展瓶颈？',
          options: ['从不', '偶尔', '有时', '经常', '总是']
        },
        {
          question: '你是否经常感到与家人或朋友的关系紧张？',
          options: ['从不', '偶尔', '有时', '经常', '总是']
        },
        {
          question: '你是否经常感到自我怀疑或缺乏自信？',
          options: ['从不', '偶尔', '有时', '经常', '总是']
        },
        {
          question: '你是否经常感到无法放松？',
          options: ['从不', '偶尔', '有时', '经常', '总是']
        },
        {
          question: '你是否经常感到对未来担忧？',
          options: ['从不', '偶尔', '有时', '经常', '总是']
        },
        {
          question: '你是否经常感到工作与生活不平衡？',
          options: ['从不', '偶尔', '有时', '经常', '总是']
        },
        {
          question: '你是否经常感到需要控制一切？',
          options: ['从不', '偶尔', '有时', '经常', '总是']
        },
        {
          question: '你是否经常感到孤独或孤立？',
          options: ['从不', '偶尔', '有时', '经常', '总是']
        },
        {
          question: '你是否经常感到身体亚健康？',
          options: ['从不', '偶尔', '有时', '经常', '总是']
        }
      ],
      totalScore: 0,
      scoreLevel: '',
      analysis: '',
      suggestions: []
    }
  },
  computed: {
    scoreLevelClass() {
      if (this.totalScore <= 20) return 'low';
      if (this.totalScore <= 40) return 'medium';
      if (this.totalScore <= 60) return 'high';
      return 'very-high';
    }
  },
  methods: {
    startTest() {
      this.showQuestions = true;
      this.showResult = false;
      this.currentQuestion = 0;
      this.selectedOption = -1;
      this.answers = [];
    },
    selectOption(index) {
      this.selectedOption = index;
    },
    prevQuestion() {
      if (this.currentQuestion > 0) {
        this.currentQuestion--;
        this.selectedOption = this.answers[this.currentQuestion] || -1;
      }
    },
    nextQuestion() {
      if (this.selectedOption === -1) return;

      this.answers[this.currentQuestion] = this.selectedOption;

      if (this.currentQuestion === this.questions.length - 1) {
        this.calculateResult();
        this.showQuestions = false;
        this.showResult = true;
      } else {
        this.currentQuestion++;
        this.selectedOption = this.answers[this.currentQuestion] || -1;
      }
    },
    calculateResult() {
      this.totalScore = this.answers.reduce((sum, option) => sum + option, 0);

      if (this.totalScore <= 20) {
        this.scoreLevel = '压力水平低';
        this.analysis = '你的压力水平较低，能够较好地应对生活和工作中的挑战。继续保持良好的生活习惯和心态。';
        this.suggestions = [
          '继续保持规律的生活作息',
          '定期进行体育锻炼',
          '保持积极的社交活动',
          '学习新技能，保持个人成长'
        ];
      } else if (this.totalScore <= 40) {
        this.scoreLevel = '压力水平中等';
        this.analysis = '你的压力水平适中，有时会感到压力，但能够应对。建议采取一些压力管理策略，保持身心健康。';
        this.suggestions = [
          '学习冥想和深呼吸技巧',
          '合理安排工作和休息时间',
          '培养兴趣爱好，放松心情',
          '与朋友和家人保持良好沟通'
        ];
      } else if (this.totalScore <= 60) {
        this.scoreLevel = '压力水平较高';
        this.analysis = '你的压力水平较高，需要关注自己的身心健康，采取积极的压力管理措施。';
        this.suggestions = [
          '寻求专业心理咨询',
          '制定合理的目标和计划',
          '学习时间管理技巧',
          '增加体育锻炼，释放压力',
          '减少咖啡因和酒精的摄入'
        ];
      } else {
        this.scoreLevel = '压力水平很高';
        this.analysis = '你的压力水平很高，已经影响到了你的身心健康，建议立即采取行动，寻求专业帮助。';
        this.suggestions = [
          '立即寻求专业心理咨询',
          '调整工作和生活节奏',
          '学习放松技巧，如冥想、瑜伽等',
          '与家人和朋友分享你的感受',
          '考虑调整职业规划或生活方式'
        ];
      }
    },
    restartTest() {
      this.startTest();
    }
  }
}
</script>

<style scoped>
.test {
  padding: 40px 0;
}

.test h1 {
  font-size: 36px;
  margin-bottom: 20px;
  color: #2c3e50;
  text-align: center;
}

.test p {
  text-align: center;
  font-size: 18px;
  margin-bottom: 60px;
  color: #666;
}

/* 测评说明 */
.test-intro {
  background-color: #f9f9f9;
  padding: 40px;
  border-radius: 8px;
  text-align: center;
  margin-bottom: 40px;
}

.test-intro h2 {
  font-size: 24px;
  margin-bottom: 20px;
  color: #2c3e50;
}

.test-intro p {
  text-align: center;
  font-size: 16px;
  margin-bottom: 30px;
  color: #666;
}

.start-btn {
  padding: 15px 40px;
  background-color: #3498db;
  color: #fff;
  border: none;
  border-radius: 4px;
  font-size: 18px;
  cursor: pointer;
  transition: all 0.3s ease;
}

.start-btn:hover {
  background-color: #2980b9;
  transform: translateY(-2px);
}

/* 测评问卷 */
.test-questions {
  background-color: #fff;
  padding: 40px;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
  margin-bottom: 40px;
}

.question-progress {
  margin-bottom: 30px;
}

.question-progress span {
  display: block;
  margin-bottom: 10px;
  color: #666;
}

.progress-bar {
  width: 100%;
  height: 10px;
  background-color: #f0f0f0;
  border-radius: 5px;
  overflow: hidden;
}

.progress-fill {
  height: 100%;
  background-color: #3498db;
  transition: width 0.3s ease;
}

.question-item {
  margin-bottom: 30px;
}

.question-item h3 {
  font-size: 18px;
  margin-bottom: 20px;
  color: #2c3e50;
}

.options {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.option-item {
  padding: 15px;
  border: 1px solid #ddd;
  border-radius: 4px;
  cursor: pointer;
  transition: all 0.3s ease;
}

.option-item:hover {
  background-color: #f0f0f0;
}

.option-item.active {
  background-color: #3498db;
  color: #fff;
  border-color: #3498db;
}

.question-actions {
  display: flex;
  justify-content: space-between;
  margin-top: 30px;
}

.prev-btn,
.next-btn {
  padding: 12px 30px;
  border: 1px solid #ddd;
  background-color: #fff;
  color: #333;
  border-radius: 4px;
  cursor: pointer;
  transition: all 0.3s ease;
}

.prev-btn:hover,
.next-btn:hover {
  background-color: #f0f0f0;
}

.next-btn {
  background-color: #3498db;
  color: #fff;
  border-color: #3498db;
}

.next-btn:hover {
  background-color: #2980b9;
}

.prev-btn:disabled,
.next-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

/* 测评结果 */
.test-result {
  background-color: #fff;
  padding: 40px;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
  text-align: center;
}

.test-result h2 {
  font-size: 24px;
  margin-bottom: 30px;
  color: #2c3e50;
}

.result-card {
  background-color: #f9f9f9;
  padding: 30px;
  border-radius: 8px;
  margin-bottom: 30px;
  text-align: left;
}

.result-score {
  text-align: center;
  margin-bottom: 30px;
}

.result-score h3 {
  font-size: 20px;
  margin-bottom: 10px;
  color: #2c3e50;
}

.score-value {
  font-size: 48px;
  font-weight: bold;
  margin-bottom: 10px;
  color: #3498db;
}

.score-level {
  font-size: 20px;
  font-weight: bold;
  padding: 5px 20px;
  border-radius: 20px;
  display: inline-block;
}

.score-level.low {
  background-color: #27ae60;
  color: #fff;
}

.score-level.medium {
  background-color: #f39c12;
  color: #fff;
}

.score-level.high {
  background-color: #e67e22;
  color: #fff;
}

.score-level.very-high {
  background-color: #e74c3c;
  color: #fff;
}

.result-analysis {
  margin-bottom: 30px;
}

.result-analysis h4,
.result-suggestions h4 {
  font-size: 18px;
  margin-bottom: 15px;
  color: #2c3e50;
}

.result-analysis p {
  text-align: left;
  font-size: 16px;
  color: #666;
  margin-bottom: 0;
}

.result-suggestions ul {
  list-style: none;
  padding: 0;
}

.result-suggestions li {
  padding: 10px 0;
  border-bottom: 1px solid #ddd;
  color: #666;
}

.result-suggestions li:last-child {
  border-bottom: none;
}

.restart-btn {
  padding: 12px 30px;
  background-color: #3498db;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: all 0.3s ease;
}

.restart-btn:hover {
  background-color: #2980b9;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .test h1 {
    font-size: 28px;
  }
  
  .test-intro,
  .test-questions,
  .test-result {
    padding: 20px;
  }
  
  .question-actions {
    flex-direction: column;
    gap: 10px;
  }
  
  .prev-btn,
  .next-btn {
    width: 100%;
  }
}
</style>
