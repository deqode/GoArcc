export const getStorage = (item) => (JSON.parse(sessionStorage.getItem(item)))
export const setStorage = (item, data) => (sessionStorage.setItem(item, JSON.stringify(data)))
export const removeStorage = (item) => (sessionStorage.removeItem(item))