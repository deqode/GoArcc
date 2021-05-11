import axios from 'axios'
import { ResponseError, UserResponse } from '../../interface'

export const setUserSession = async (data: UserResponse): Promise<ResponseError> => {
  // todo:handle data
  try {
    const response = await axios.post(`/api/session/set`, data, {
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
      },
    })
    if (response.data.susscess) {
      return {
        error: false,
        message: '',
      }
    } else
      return {
        error: true,
        message: 'Data Not found',
        // todo : integrate with error message from backend
      }
  } catch (e) {
    return {
      error: true,
      message: 'Network Error',
      // todo : integrate with error message from backend
    }
  }
}

export const destroyUserSession = async (): Promise<ResponseError> => {
  try {
    const response = await axios.post(
      `/api/session/destroy`,
      {},
      {
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
      }
    )
    if (response.data.susscess) {
      return {
        error: false,
        message: '',
      }
    } else
      return {
        error: true,
        message: 'Data Not found',
        // todo : integrate with error message from backend
      }
  } catch (e) {
    return {
      error: true,
      message: 'Network Error',
      // todo : integrate with error message from backend
    }
  }
}
