import { withIronSession } from 'next-iron-session'
import { sessionCongfig } from '../../../utils/constants'
// TODO:validate
const handler = async (req: any, res: any) => {
  req.session.set('user', req.body)
  await req.session.save()
  res.json({ susscess: true })
}

export default withIronSession(handler, sessionCongfig)
