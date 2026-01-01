import axios from 'axios'

const api = axios.create({
  timeout: 5000,
  headers: { 'Content-Type': 'application/json' }
})

// Request interceptor to set base URL dynamically
api.interceptors.request.use(config => {
  const baseURL = localStorage.getItem('apiBaseUrl') || 'http://localhost:8888';
  config.baseURL = baseURL;
  return config;
}, error => {
  return Promise.reject(error);
});

export default api
