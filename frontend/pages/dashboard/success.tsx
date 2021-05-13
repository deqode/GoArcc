import { withSentry } from '@sentry/nextjs'
import { withIronSession } from 'next-iron-session'
import Head from 'next/head'
import { sessionCongfig } from '../../utils/constants'
import { redirectToLandingPage } from '../../utils/redirects'
import { validateUser } from '../../utils/user'

export default function Success() {
  return (
    <div>
      <Head>
        <title>Congratulation !!</title>
        <link rel="icon" href="assets/alfred.png" />
      </Head>
      <section className="content_section">
        <div className="container">
          <div className="row">
            <div className="col-md-6">
              <h1 className="page-head">Alfred has cloned your app repository</h1>
            </div>
            <div className="col-md-6 my-auto login-right"></div>
          </div>
        </div>
      </section>
    </div>
  )
}
export const getServerSideProps = withSentry(
  withIronSession(async ({ req }) => {
    if (validateUser(req)) {
      return { props: { user: req.session.get('user') } }
    }
    return redirectToLandingPage()
  }, sessionCongfig)
)
