<template>
  <div class="selection-container">
    <div class="content fade-in-up">
      <div class="header">
        <h1>Select Backend</h1>
        <p>Choose your preferred service provider</p>
      </div>
      
      <div class="options">
        <div class="card option-card" @click="selectBackend('http://localhost:8000', 'Python')">
          <div class="icon-wrapper python-bg">
            <span>Py</span>
          </div>
          <h3>Python Backend</h3>
          <p class="port">Port: 8000</p>
          <div class="status-indicator">
            <span class="dot"></span> Online
          </div>
        </div>

        <div class="card option-card" @click="selectBackend('http://localhost:8888', 'Go')">
          <div class="icon-wrapper go-bg">
            <span>Go</span>
          </div>
          <h3>Go Backend</h3>
          <p class="port">Port: 8888</p>
          <div class="status-indicator">
            <span class="dot"></span> Online
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  methods: {
    selectBackend(url, name) {
      localStorage.setItem('apiBaseUrl', url);
      localStorage.setItem('backendName', name);
      this.$router.push('/login');
    }
  }
}
</script>

<style scoped>
.selection-container {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  padding: 20px;
}

.content {
  width: 100%;
  max-width: 800px;
}

.header {
  text-align: center;
  margin-bottom: 50px;
  color: white;
  text-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.header h1 {
  font-size: 2.5rem;
  font-weight: 800;
  margin-bottom: 10px;
}

.header p {
  font-size: 1.1rem;
  opacity: 0.9;
}

.options {
  display: flex;
  gap: 30px;
  justify-content: center;
  flex-wrap: wrap;
}

.option-card {
  background: var(--card-bg);
  padding: 40px;
  border-radius: 20px;
  box-shadow: 0 10px 30px rgba(0,0,0,0.1);
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  text-align: center;
  width: 280px;
  border: 2px solid transparent;
  backdrop-filter: blur(10px);
}

.option-card:hover {
  transform: translateY(-10px);
  box-shadow: 0 20px 40px rgba(0,0,0,0.2);
  border-color: rgba(255,255,255,0.5);
}

.icon-wrapper {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 20px;
  font-size: 1.8rem;
  font-weight: 700;
  color: white;
  box-shadow: 0 8px 16px rgba(0,0,0,0.15);
}

.python-bg {
  background: linear-gradient(135deg, #306998 0%, #FFD43B 100%);
}

.go-bg {
  background: linear-gradient(135deg, #00ADD8 0%, #5DC9E2 100%);
}

h3 {
  color: var(--text-color);
  font-size: 1.4rem;
  margin-bottom: 8px;
}

.port {
  color: #888;
  margin-bottom: 20px;
  font-family: monospace;
  background: rgba(0,0,0,0.05);
  display: inline-block;
  padding: 4px 12px;
  border-radius: 20px;
  font-size: 0.9rem;
}

.status-indicator {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 6px;
  font-size: 0.9rem;
  color: #27ae60;
  font-weight: 600;
}

.dot {
  width: 8px;
  height: 8px;
  background: #27ae60;
  border-radius: 50%;
  display: inline-block;
  box-shadow: 0 0 0 2px rgba(39, 174, 96, 0.2);
  animation: pulse 2s infinite;
}

@keyframes pulse {
  0% { box-shadow: 0 0 0 0 rgba(39, 174, 96, 0.4); }
  70% { box-shadow: 0 0 0 6px rgba(39, 174, 96, 0); }
  100% { box-shadow: 0 0 0 0 rgba(39, 174, 96, 0); }
}

.fade-in-up {
  animation: fadeInUp 0.6s ease-out;
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
