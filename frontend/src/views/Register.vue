<template>
  <div class="register-container">
    <div class="card fade-in-up">
      <div class="header">
        <h2>Create Account</h2>
        <p>Join us today</p>
      </div>

      <form @submit.prevent="onSubmit">
        <div class="form-group">
          <label>Username</label>
          <input 
            v-model="username" 
            type="text" 
            placeholder="Choose a username"
            required 
            class="form-input"
          />
        </div>

        <div class="form-group">
          <label>Email</label>
          <input 
            v-model="email" 
            type="email" 
            placeholder="Enter your email"
            required 
            class="form-input"
          />
        </div>

        <div class="form-group">
          <label>Password</label>
          <div class="password-wrapper">
            <input 
              :type="showPassword ? 'text' : 'password'" 
              v-model="password" 
              placeholder="Create a password"
              required 
              class="form-input"
            />
            <button type="button" class="toggle-password" @click="showPassword = !showPassword">
              <svg v-if="showPassword" xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="#666" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"></path><circle cx="12" cy="12" r="3"></circle></svg>
              <svg v-else xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="#666" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M17.94 17.94A10.07 10.07 0 0 1 12 20c-7 0-11-8-11-8a18.45 18.45 0 0 1 5.06-5.94M9.9 4.24A9.12 9.12 0 0 1 12 4c7 0 11 8 11 8a18.5 18.5 0 0 1-2.16 3.19m-6.72-1.07a3 3 0 1 1-4.24-4.24"></path><line x1="1" y1="1" x2="23" y2="23"></line></svg>
            </button>
          </div>
        </div>

        <div class="form-group">
          <label>Confirm Password</label>
          <div class="password-wrapper">
            <input 
              :type="showConfirmPassword ? 'text' : 'password'" 
              v-model="confirmPassword" 
              placeholder="Confirm your password"
              required 
              class="form-input"
            />
            <button type="button" class="toggle-password" @click="showConfirmPassword = !showConfirmPassword">
              <svg v-if="showConfirmPassword" xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="#666" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"></path><circle cx="12" cy="12" r="3"></circle></svg>
              <svg v-else xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="#666" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M17.94 17.94A10.07 10.07 0 0 1 12 20c-7 0-11-8-11-8a18.45 18.45 0 0 1 5.06-5.94M9.9 4.24A9.12 9.12 0 0 1 12 4c7 0 11 8 11 8a18.5 18.5 0 0 1-2.16 3.19m-6.72-1.07a3 3 0 1 1-4.24-4.24"></path><line x1="1" y1="1" x2="23" y2="23"></line></svg>
            </button>
          </div>
        </div>

        <div class="actions">
          <button type="submit" class="btn btn-primary" :disabled="isLoading">
            <span v-if="isLoading">Creating Account...</span>
            <span v-else>Register</span>
          </button>
          <button type="button" class="btn btn-secondary" @click="goLogin">Back to Login</button>
        </div>

        <transition name="fade">
          <p v-if="error" class="error-msg">{{ error }}</p>
        </transition>
      </form>
    </div>
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
      error: '',
      showPassword: false,
      showConfirmPassword: false,
      isLoading: false
    }
  },
  methods: {
    goLogin() {
      this.$router.push('/login')
    },
    async onSubmit() {
      this.error = ''
      if (this.password !== this.confirmPassword) {
        this.error = 'Passwords do not match'
        return
      }
      
      this.isLoading = true
      try {
        const res = await api.post('/register', {
          username: this.username,
          password: this.password,
          confirm_password: this.confirmPassword,
          email: this.email
        })
        const data = res.data
        if (data && data.code === 200) {
          // Success, maybe show a toast or auto login?
          // For now, redirect to login
          this.$router.push('/login')
        } else {
          this.error = data.message || 'Registration failed'
        }
      } catch (e) {
        this.error = e.response?.data?.message || 'Cannot connect to server'
      } finally {
        this.isLoading = false
      }
    }
  }
}
</script>

<style scoped>
.register-container {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  padding: 20px;
}

.card {
  background: var(--card-bg);
  padding: 40px;
  border-radius: 16px;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.1);
  width: 100%;
  max-width: 420px;
  backdrop-filter: blur(10px);
}

.header {
  text-align: center;
  margin-bottom: 30px;
}

.header h2 {
  color: var(--primary-color);
  margin-bottom: 8px;
  font-weight: 700;
}

.header p {
  color: #666;
  font-size: 0.95rem;
}

.form-group {
  margin-bottom: 20px;
}

.form-group label {
  display: block;
  margin-bottom: 8px;
  color: var(--text-color);
  font-weight: 500;
  font-size: 0.9rem;
}

.form-input {
  width: 100%;
  padding: 12px 16px;
  border: 2px solid transparent;
  background: var(--input-bg);
  border-radius: 8px;
  font-size: 1rem;
  transition: all 0.3s ease;
}

.form-input:focus {
  background: #fff;
  border-color: var(--primary-color);
  box-shadow: 0 0 0 3px rgba(108, 92, 231, 0.1);
}

.password-wrapper {
  position: relative;
}

.toggle-password {
  position: absolute;
  right: 12px;
  top: 50%;
  transform: translateY(-50%);
  background: none;
  display: flex;
  align-items: center;
  padding: 4px;
  border-radius: 4px;
  transition: background 0.2s;
}

.toggle-password:hover {
  background: rgba(0,0,0,0.05);
}

.actions {
  display: flex;
  flex-direction: column;
  gap: 12px;
  margin-top: 30px;
}

.btn {
  width: 100%;
  padding: 12px;
  border-radius: 8px;
  font-weight: 600;
  font-size: 1rem;
  transition: transform 0.2s, box-shadow 0.2s;
}

.btn:active {
  transform: scale(0.98);
}

.btn-primary {
  background: var(--primary-color);
  color: white;
  box-shadow: 0 4px 6px rgba(108, 92, 231, 0.25);
}

.btn-primary:hover {
  background: var(--primary-hover);
  box-shadow: 0 6px 12px rgba(108, 92, 231, 0.3);
}

.btn-primary:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.btn-secondary {
  background: transparent;
  color: var(--primary-color);
  border: 2px solid transparent;
}

.btn-secondary:hover {
  background: rgba(108, 92, 231, 0.05);
}

.error-msg {
  color: var(--error-color);
  text-align: center;
  margin-top: 16px;
  font-size: 0.9rem;
  background: rgba(214, 48, 49, 0.1);
  padding: 8px;
  border-radius: 6px;
}

.fade-in-up {
  animation: fadeInUp 0.5s ease-out;
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
