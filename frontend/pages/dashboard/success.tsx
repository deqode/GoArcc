import { Paper } from '@material-ui/core'
import { withSentry } from '@sentry/nextjs'
import { withIronSession } from 'next-iron-session'
import Head from 'next/head'
import BasicLayout from '../../components/layouts/BasicLayout'
import { sessionCongfig } from '../../utils/constants'
import { redirectToLandingPage } from '../../utils/redirects'
import { validateUser } from '../../utils/user'

export default function Success() {
  return (
    <Paper elevation={0}>
      <Head>
        <title>Login to Alfred</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <BasicLayout
        heading={'Success'}
        subHeading={'Alfred has cloned your app repository'}
        component={''}
      />
    </Paper>
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
