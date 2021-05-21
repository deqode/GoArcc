import axios from 'axios'

import { ResponseError } from '../../intefaces/interface'

const defaultHeaders = { Accept: 'application/json', 'Content-Type': 'application/json' }
interface Header {
  [key: string]: string
}
export const get = async <Type>(link: string, header?: Header): Promise<ResponseError<Type>> => {
  try {
    const response = await axios.get<Type>(link, {
      headers: {
        ...defaultHeaders,
        ...header,
      },
    })
    const data = response.data
    return {
      data: data,
      error: false,
      message: '',
    }
  } catch (err) {
    // TODO: check error type
    return {
      error: true,
      message: 'Network Error',
      data: null,
    }
  }
}
