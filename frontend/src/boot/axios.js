import { boot } from 'quasar/wrappers'
import axios from 'axios'
import { handleErrors } from 'src/utils/api'

// Be careful when using SSR for cross-request state pollution
// due to creating a Singleton instance here;
// If any client changes this (global) instance, it might be a
// good idea to move this instance creation inside of the
// "export default () => {}" function below (which runs individually
// for each client)
// console.log(new URL('api', process.env.BACKEND_URL).href)
// console.log(process.env)

const api = axios.create({
  baseURL: new URL('api', process.env.BACKEND_URL).href,
  timeout: 180000,
  headers: {
    'X-Requested-With': 'XMLHttpRequest',
    'Content-Type': 'Application/json',
  }
})

const blockchain = axios.create({
  baseURL: new URL(process.env.BLOCKCHAIN_URL).href,
  timeout: 180000,
  headers: {
    'Content-Type': 'Application/json',
  }
})

export default boot(({ app }) => {
  api.interceptors.response.use((response) => {
    return response
  }, (error) => {
    if (error.config && !error.config.ignoreErrorHandling) {
      return handleErrors(error, { customErrorHandlers: error.config?.customErrorHandlers })
    }

    return Promise.reject(error)
  })

  blockchain.interceptors.response.use((response) => {
    return response
  }, (error) => {
    if (error.config && !error.config.ignoreErrorHandling) {
      return handleErrors(error, { customErrorHandlers: error.config?.customErrorHandlers })
    }

    return Promise.reject(error)
  })

  // for use inside Vue files (Options API) through this.$axios and this.$api

  app.config.globalProperties.$axios = axios
  // ^ ^ ^ this will allow you to use this.$axios (for Vue Options API form)
  //       so you won't necessarily have to import axios in each vue file

  app.config.globalProperties.$api = api
  // ^ ^ ^ this will allow you to use this.$api (for Vue Options API form)
  //       so you can easily perform requests against your app's API
})

export { api, blockchain }
