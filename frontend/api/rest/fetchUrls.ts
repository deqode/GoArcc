import { ResponseError } from '../../intefaces/interface'
import { get } from '../../services/rest'
import { SERVER } from '../../utils/constants'

interface GithubVCSConnection extends ResponseError<GithubVCSConnection> {
  redirectUrl: string
}
interface LoginURL extends ResponseError<LoginURL> {
  url: string
}

export const getLoginURL = async (): Promise<LoginURL> => {
  const response = await get<LoginURL>(`${SERVER}/authentication/login`)
  if (response.error) {
    return { url: '', ...response }
  }
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
}

export const getGithubVCSConnection = async (idToken: string): Promise<GithubVCSConnection> => {
  //TODO:validate idToken
  const response = await get<GithubVCSConnection>(`${SERVER}/vcs-connection/authorize/GITHUB`, {
    Authorization: `Bearer ${idToken}`,
  })
  if (response.error) {
    return { redirectUrl: '', ...response }
  }
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
}
