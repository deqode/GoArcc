import axios from 'axios'
import { ResponseError, UserResponse } from '../../interface'

import { SERVER } from '../../utils/constants'

export const getAuth0Callback = async (code: string, state: string): Promise<UserAuthResponse> => {
  if (code !== '' && state !== '')
    try {
      const response = await axios.get(
        `${SERVER}/authentication/callback?code=${code}&state=${state}`
      )
      if (response.data && response.data.userId) {
        return {
          userId: response.data.userId,
          idToken: response.data.idToken,
          accessToken: response.data.idToken,
          error: false,
          message: '',
        }
      } else
        return {
          userId: '',
          idToken: '',
          accessToken: '',
          error: true,
          message: 'Data Not found',
          // todo : integrate with error message from backend
        }
    } catch (e) {
      return {
        userId: '',
        idToken: '',
        accessToken: '',
        error: true,
        message: 'Network Error',
        // todo : integrate with error message from backend
      }
    }
  else
    return {
      userId: '',
      idToken: '',
      accessToken: '',
      error: true,
      message: 'Invalid types',
      // todo : integrate with error message from backend
    }
}

export const getVCSConnectionGitHubCallback = async (
  code: string,
  accountId: string,
  idToken: string
): Promise<ResponseError> => {
  if (code !== '' && accountId !== '')
    try {
      const response = await axios.get(
        `${SERVER}/vcs-connection/GITHUB/callback?code=${code}&account_id=${accountId}`,
        { headers: { Authorization: `Bearer ${idToken}` } }
      )
      if (response.data && response.data.accountId) {
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
  else
    return {
      error: true,
      message: 'Invalid types',
      // todo : integrate with error message from backend
    }
}

interface UserAuthResponse extends ResponseError, UserResponse {}
