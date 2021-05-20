import { ResponseError, UserResponse } from '../../intefaces/interface'
import { post } from '../../services/rest/post'

interface Session {
  success: boolean
}
export const setUserSession = async (data: UserResponse): Promise<ResponseError<Session>> => {
  // TODO:handle data
  const response = await post<Session, UserResponse>(`/api/session/set`, data)
  if (response.error) {
    return { ...response }
  }
  if (response.data && response.data.success) {
    return {
      error: false,
      message: '',
    }
  } else
    return {
      error: true,
      message: 'Server Error',
      // TODO : integrate with error message from backend
    }
}

export const destroyUserSession = async (): Promise<ResponseError<Session>> => {
  const response = await post<Session, null>(`/api/session/destroy`, null)
  if (response.error) {
    return { ...response }
  }
  if (response.data && response.data.success) {
    return {
      error: false,
      message: '',
    }
  } else
    return {
      error: true,
      message: 'Server Error',
      // TODO : integrate with error message from backend
    }
}
