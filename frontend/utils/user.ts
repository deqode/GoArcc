import { Handler, withIronSession } from 'next-iron-session'
import { sessionCongfig } from './constants'

export const validateUser = (req: any): boolean => {
  return req.session.get('user')
  // TODO : add try catch
  // TODO: need to add validation
}

export const sessionPropsWrapper = (handler: Handler) => withIronSession(handler, sessionCongfig)
// TODO: need to fix withSentry
