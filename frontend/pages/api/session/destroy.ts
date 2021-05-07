import { withIronSession } from "next-iron-session";
import { sessionCongfig } from "../../../utils/constants";

async function handler(req: any, res: any) {
    console.log(req.body)
    req.session.destroy("user");
    res.json({ susscess: true });
}

export default withIronSession(handler, sessionCongfig);