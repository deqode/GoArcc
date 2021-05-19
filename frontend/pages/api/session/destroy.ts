import { withIronSession } from 'next-iron-session'
import { sessionCongfig } from '../../../utils/constants'
import { withSentry } from '@sentry/nextjs'

// TODO: notify platform to invalidate
// TODO:validate
export const handler = async (req: any, res: any) => {
  req.session.destroy('user')
  res.json({ success: true })
}

export default withSentry(withIronSession(handler, sessionCongfig))
