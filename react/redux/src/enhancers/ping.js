/*eslint-disable */
export const ping = store => next => action => {
    console.log('ping')
    return next(action)
  }
  /*eslint-enable */