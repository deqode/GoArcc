import { withIronSession } from 'next-iron-session'
import { sessionCongfig } from '../../../utils/constants'
import { withSentry } from '@sentry/nextjs'

// TODO: notify platform to invalidate
// TODO:validate
const handler = async (req: any, res: any) => {
  req.session.destroy('user')
  res.json({ susscess: true })
}

export default withSentry(withIronSession(handler, sessionCongfig))
