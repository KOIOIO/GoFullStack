<template>
  <div style="text-align:center;margin-top:80px">
    <h1>Hello World</h1>
    <p>欢迎！你已登录。</p>
    <button @click="logout">登出</button>
  </div>
</template>

<script>
export default {
  methods: {
    logout() {
      localStorage.removeItem('token')
      this.$router.push('/login')
    }
  }
}
</script>
