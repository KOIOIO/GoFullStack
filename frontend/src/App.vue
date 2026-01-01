<template>
  <router-view v-slot="{ Component }">
    <transition name="fade" mode="out-in">
      <component :is="Component" />
    </transition>
  </router-view>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
:root {
  --primary-color: #6c5ce7;
  --primary-hover: #5b4cc4;
  --secondary-color: #a29bfe;
  --bg-gradient: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  --text-color: #2d3436;
  --error-color: #d63031;
  --input-bg: #f5f6fa;
  --card-bg: rgba(255, 255, 255, 0.95);
}

* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

body {
  font-family: 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  background: var(--bg-gradient);
  min-height: 100vh;
  color: var(--text-color);
  -webkit-font-smoothing: antialiased;
}

#app {
  min-height: 100vh;
}

/* 全局过渡动画 */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

/* 通用按钮样式 */
button {
  cursor: pointer;
  border: none;
  outline: none;
  font-family: inherit;
}

/* 通用输入框样式 */
input {
  font-family: inherit;
}
</style>
