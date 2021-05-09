import { withIronSession } from 'next-iron-session'
import { sessionCongfig } from '../../../utils/constants'
// todo:validate
async function handler(req: any, res: any) {
  console.log(req.body)
  req.session.set('user', req.body)
  await req.session.save()
  res.json({ susscess: true })
}

export default withIronSession(handler, sessionCongfig)
