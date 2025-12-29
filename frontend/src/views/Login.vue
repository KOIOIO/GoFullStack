<template>
  <div style="max-width:400px;margin:40px auto">
    <h2>登录</h2>
    <form @submit.prevent="onSubmit">
      <div>
        <label>用户名</label>
        <input v-model="username" required />
      </div>
      <div>
        <label>密码</label>
        <input type="password" v-model="password" required />
      </div>
      <div style="margin-top:12px">
        <button type="submit">登录</button>
        <button type="button" @click="goRegister">去注册</button>
      </div>
      <p v-if="error" style="color:red">{{ error }}</p>
    </form>
  </div>
</template>

<script>
import api from '../api'

export default {
  data() {
    return {
      username: '',
      password: '',
      error: ''
    }
  },
  methods: {
    goRegister() {
      this.$router.push('/register')
    },
    async onSubmit() {
      this.error = ''
      try {
        const res = await api.post('/login', { username: this.username, password: this.password })
        const data = res.data
        if (data && data.token) {
          localStorage.setItem('token', data.token)
          this.$router.push('/hello')
        } else {
          this.error = data.message || '登录失败'
        }
      } catch (e) {
        this.error = e.response?.data?.message || '无法连接到后台'
      }
    }
  }
}
</script>
