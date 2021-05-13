import { Paper } from '@material-ui/core'
import Head from 'next/head'
import { sessionCongfig } from '../utils/constants'
import { withIronSession } from 'next-iron-session'
import { getLoginURL } from '../api/rest/fetchUrls'
import { validateUser } from '../utils/user'
import { ReactElement } from 'react'
import { redirectToErrorPage, redirectToDashboard } from '../utils/redirects'
import { withSentry } from '@sentry/nextjs'
import BasicLayout from '../components/layouts/BasicLayout'
import Link from 'next/link'

const Landing = ({ url }: { url: string }): ReactElement => {
  return (
    <Paper elevation={0}>
      <Head>
        <title>Login to Alfred</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <BasicLayout
        heading={'Welcome to Alfred'}
        subHeading={'Provide your repo, Select your cloud, and let Alfred do the heavy lifting '}
        component={<Link href={url}>SignUP</Link>}
      />
    </Paper>
  )
}

export const getServerSideProps = withSentry(
  withIronSession(async ({ req }) => {
    if (!validateUser(req)) {
      const res = await getLoginURL()
      if (!res.error) return { props: { url: res.url } }
      return redirectToErrorPage('Network Error')
    }
    return redirectToDashboard()
  }, sessionCongfig)
)

export default Landing
