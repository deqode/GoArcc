import { Button, Paper } from '@material-ui/core'
import Head from 'next/head'
import { getLoginURL } from '../api/rest/fetchUrls'
import { sessionPropsWrapper, validateUser } from '../utils/user'
import { ReactElement } from 'react'
import { redirectToErrorPage, redirectToDashboard } from '../utils/redirects'
import BasicLayout from '../components/layouts/BasicLayout'
import { NextApiRequest } from 'next'
import { withSentry } from '@sentry/nextjs'

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
        component={
          <Button variant="contained" color="primary" href={url}>
            SignUP
          </Button>
        }
      />
    </Paper>
  )
}

export const handler = async ({ req }: { req: NextApiRequest }) => {
  if (!validateUser(req)) {
    const res = await getLoginURL()
    if (!res.error) return { props: { url: res.url } }
    return redirectToErrorPage('Network Error')
  }
  return redirectToDashboard()
}

export const getServerSideProps = withSentry(sessionPropsWrapper(handler))

export default Landing
