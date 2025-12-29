<template>
  <div style="max-width:420px;margin:40px auto">
    <h2>注册</h2>
    <form @submit.prevent="onSubmit">
      <div>
        <label>用户名</label>
        <input v-model="username" required />
      </div>
      <div>
        <label>邮箱</label>
        <input v-model="email" type="email" required />
      </div>
      <div>
        <label>密码</label>
        <input type="password" v-model="password" required />
      </div>
      <div>
        <label>确认密码</label>
        <input type="password" v-model="confirmPassword" required />
      </div>
      <div style="margin-top:12px">
        <button type="submit">注册</button>
        <button type="button" @click="goLogin">返回登录</button>
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
      email: '',
      password: '',
      confirmPassword: '',
      error: ''
    }
  },
  methods: {
    goLogin() {
      this.$router.push('/login')
    },
    async onSubmit() {
      this.error = ''
      if (this.password !== this.confirmPassword) {
        this.error = '两次密码不一致'
        return
      }
      try {
        const res = await api.post('/register', {
          username: this.username,
          password: this.password,
          confirm_password: this.confirmPassword,
          email: this.email
        })
        const data = res.data
        if (data && data.code === 0) {
          this.$router.push('/login')
        } else {
          this.error = data.message || '注册失败'
        }
      } catch (e) {
        this.error = e.response?.data?.message || '无法连接到后台'
      }
    }
  }
}
</script>
