import { Notify } from "quasar"

export const handleErrors = async (error, { router, store, customErrorHandlers }) => {
  const defaultErrorHandlers = {
    401: async (data) => {
      Notify.create({
        type: 'negative',
        message: 'Unauthorized.'
      })
    },
    404: data => {
      Notify.create({
        type: 'negative',
        message: data.message ?? 'Not found.'
      })
    },
    422: data => {
      const error = data.errors ? Object.values(data.errors).flat().shift() : data.message
      Notify.create({
        type: 'negative',
        message: error ?? 'Invalid data.'
      })
    },
    500: (/* data */) => {
      Notify.create({
        type: 'negative',
        message: 'Server error.'
      })
    }
  }

  const errorHandler = { ...defaultErrorHandlers }

  Object.assign(errorHandler, customErrorHandlers)

  const { status } = error.response

  return (
    (
      errorHandler[status] &&
      (await errorHandler[status](error.response.data, {
        config: error.config,
        defaultHandler: (data = error.response.data) => defaultErrorHandlers[status](data)
      }))
    ) ||
    Promise.reject(error)
  )
}