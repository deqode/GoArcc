import { withIronSession } from 'next-iron-session'
import { sessionCongfig } from '../../../utils/constants'
// todo: notify platform to invalidate
// todo:validate
const handler = async (req: any, res: any) => {
  req.session.destroy('user')
  res.json({ susscess: true })
}

export default withIronSession(handler, sessionCongfig)
