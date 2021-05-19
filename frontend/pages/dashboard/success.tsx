import { Paper } from '@material-ui/core'
import { withSentry } from '@sentry/nextjs'
import Head from 'next/head'
import BasicLayout from '../../components/layouts/BasicLayout'
import { IronSessionRequest } from '../../interface'
import { redirectToLandingPage } from '../../utils/redirects'
import { sessionPropsWrapper, validateUser } from '../../utils/user'

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

export const handler = async ({ req }: { req: IronSessionRequest }) => {
  if (validateUser(req)) {
    return { props: { user: req.session.get('user') } }
  }
  return redirectToLandingPage()
}

export const getServerSideProps = withSentry(sessionPropsWrapper(handler))
