export const getStorage = (item:string) => (JSON.parse(<string>sessionStorage.getItem(item)))
export const setStorage = (item:string, data:any) => (sessionStorage.setItem(item, JSON.stringify(data)))
export const removeStorage = (item:string) => (sessionStorage.removeItem(item))