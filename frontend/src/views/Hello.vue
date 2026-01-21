<template>
  <div class="hello-container">
    <canvas ref="nebulaCanvas" class="nebula-bg"></canvas>
    <div class="content fade-in-up">
      <div class="card">
        <div class="icon-wrapper">
          <svg xmlns="http://www.w3.org/2000/svg" width="64" height="64" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-smile"><circle cx="12" cy="12" r="10"></circle><path d="M8 14s1.5 2 4 2 4-2 4-2"></path><line x1="9" y1="9" x2="9.01" y2="9"></line><line x1="15" y1="9" x2="15.01" y2="9"></line></svg>
        </div>
        <h1>Welcome to the Galaxy</h1>
        <p class="welcome-text">欢迎来到星云空间！你已登录成功。</p>
        
        <button @click="logout" class="btn-primary">
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-log-out"><path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"></path><polyline points="16 17 21 12 16 7"></polyline><line x1="21" y1="12" x2="9" y2="12"></line></svg>
          安全登出
        </button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Hello',
  data() {
    return {
      animationId: null
    }
  },
  mounted() {
    this.initNebula();
    window.addEventListener('resize', this.handleResize);
  },
  beforeUnmount() {
    if (this.animationId) {
      cancelAnimationFrame(this.animationId);
    }
    window.removeEventListener('resize', this.handleResize);
  },
  methods: {
    logout() {
      localStorage.removeItem('token')
      this.$router.push('/login')
    },
    handleResize() {
      const canvas = this.$refs.nebulaCanvas;
      if (canvas) {
        canvas.width = window.innerWidth;
        canvas.height = window.innerHeight;
      }
    },
    initNebula() {
      const canvas = this.$refs.nebulaCanvas;
      const ctx = canvas.getContext('2d');
      
      canvas.width = window.innerWidth;
      canvas.height = window.innerHeight;

      const stars = [];
      const starCount = 800;
      const centerX = canvas.width / 2;
      const centerY = canvas.height / 2;

      // 初始化星星
      for (let i = 0; i < starCount; i++) {
        stars.push({
          x: Math.random() * canvas.width,
          y: Math.random() * canvas.height,
          z: Math.random() * canvas.width, // 深度
          o: '0.' + Math.floor(Math.random() * 99) + 1 // 透明度
        });
      }

      const animate = () => {
        // 黑色背景，带一点透明度产生拖尾效果
        ctx.fillStyle = "rgba(0, 0, 0, 0.1)"; 
        ctx.fillRect(0, 0, canvas.width, canvas.height);

        for (let i = 0; i < starCount; i++) {
          const star = stars[i];
          
          // 移动逻辑：向屏幕中心移动，或者旋转
          // 这里实现一个向中心吸入并旋转的效果，类似穿越虫洞
          
          star.z -= 10; // 速度
          
          if (star.z <= 0) {
            star.z = canvas.width;
            star.x = Math.random() * canvas.width;
            star.y = Math.random() * canvas.height;
          }

          const k = 128.0 / star.z;
          const px = (star.x - centerX) * k + centerX;
          const py = (star.y - centerY) * k + centerY;

          if (px >= 0 && px <= canvas.width && py >= 0 && py <= canvas.height) {
            const size = (1 - star.z / canvas.width) * 2.5;
            const shade = parseInt((1 - star.z / canvas.width) * 255);
            
            // 颜色动态变化
            ctx.fillStyle = "rgb(" + shade + "," + shade + "," + 255 + ")";
            ctx.beginPath();
            ctx.arc(px, py, size, 0, Math.PI * 2);
            ctx.fill();
          }
        }
        
        // 绘制星云雾气 (可选，为了性能简单点可以先不加太复杂的)
        
        this.animationId = requestAnimationFrame(animate);
      };

      animate();
    }
  }
}
</script>

<style scoped>
.hello-container {
  position: relative;
  width: 100vw;
  height: 100vh;
  overflow: hidden;
  background-color: #000;
}

.nebula-bg {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: 1;
}

.content {
  position: relative;
  z-index: 2;
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
}

.card {
  background: rgba(255, 255, 255, 0.1); /* 更加透明 */
  padding: 3rem;
  border-radius: 20px;
  box-shadow: 0 0 50px rgba(108, 92, 231, 0.3); /* 发光效果 */
  width: 100%;
  max-width: 450px;
  text-align: center;
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.1);
  color: white;
}

.icon-wrapper {
  margin-bottom: 1.5rem;
  color: #a29bfe;
  animation: float 3s ease-in-out infinite;
  filter: drop-shadow(0 0 10px #6c5ce7);
}

h1 {
  margin: 0 0 0.5rem;
  color: #fff;
  font-size: 2.5rem;
  font-weight: 700;
  text-shadow: 0 0 20px rgba(108, 92, 231, 0.8);
}

.welcome-text {
  color: #dfe6e9;
  margin-bottom: 2rem;
  font-size: 1.1rem;
}

.btn-primary {
  width: 100%;
  padding: 12px;
  background: linear-gradient(135deg, #6c5ce7, #a29bfe);
  color: white;
  border: none;
  border-radius: 12px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 10px;
  box-shadow: 0 5px 15px rgba(108, 92, 231, 0.4);
}

.btn-primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(108, 92, 231, 0.6);
}

.btn-primary:active {
  transform: translateY(0);
}

@keyframes float {
  0% { transform: translateY(0px); }
  50% { transform: translateY(-10px); }
  100% { transform: translateY(0px); }
}

.fade-in-up {
  animation: fadeInUp 1s ease-out forwards;
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
