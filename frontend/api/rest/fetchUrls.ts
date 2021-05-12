import axios from 'axios'
import { ResponseError } from '../../interface'

import { SERVER } from '../../utils/constants'

export const getLoginURL = async (): Promise<LoginURL> => {
  try {
    const response = await axios.get(`${SERVER}/authentication/login`)
    if (response.data && response.data.url) {
      return {
        url: response.data.url,
        error: false,
        message: '',
      }
    } else
      return {
        url: '',
        error: true,
        message: 'Data Not found',
        // TODO : integrate with error message from backend
      }
  } catch (e) {
    return {
      url: '',
      error: true,
      message: 'Network Error',
      // TODO : integrate with error message from backend
    }
  }
}

export const getGithubVCSConnection = async (idToken: string): Promise<GithubVCSConnection> => {
  try {
    const response = await axios.get(`${SERVER}/vcs-connection/authorize/GITHUB`, {
      headers: { Authorization: `Bearer ${idToken}` },
    })
    if (response.data && response.data.redirectUrl) {
      return {
        redirectUrl: response.data.redirectUrl,
        error: false,
        message: '',
      }
    } else
      return {
        redirectUrl: '',
        error: true,
        message: 'Data Not found',
        // TODO : integrate with error message from backend
      }
  } catch (e) {
    return {
      redirectUrl: '',
      error: true,
      message: 'Network Error',
      // TODO : integrate with error message from backend
    }
  }
}

interface GithubVCSConnection extends ResponseError {
  redirectUrl: string
}
interface LoginURL extends ResponseError {
  url: string
}
