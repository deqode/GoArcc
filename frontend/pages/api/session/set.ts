import { withIronSession } from 'next-iron-session'
import { sessionCongfig } from '../../../utils/constants'
import { withSentry } from '@sentry/nextjs'

// TODO:validate
export const handler = async (req: any, res: any) => {
  req.session.set('user', req.body)
  await req.session.save()
  res.json({ success: true })
}

export default withSentry(withIronSession(handler, sessionCongfig))
