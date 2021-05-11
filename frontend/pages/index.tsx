import { Grid, Paper, Typography } from '@material-ui/core'
import Head from 'next/head'
import { sessionCongfig } from '../utils/constants'
import { withIronSession } from 'next-iron-session'
import { getLoginURL } from '../api/rest/fetchUrls'
import { validateUser } from '../utils/user'
import { ReactElement } from 'react'

const Login = ({ url }: { url: string }): ReactElement => {
  return (
    <div>
      <Head>
        <title>Login to Alfred</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Paper elevation={0} style={{ padding: '100px' }} />
      <Grid justify="center" alignItems="center" spacing={1} container>
        <Grid item md={4}>
          <Typography variant="h1" component="h1" style={{ fontSize: '50px' }}>
            Let us make your cloud work for you
          </Typography>
          <Typography variant="subtitle2" style={{ fontSize: '15px' }}>
            provide your repo, select your cloud, and let Alfred do the heavy lifting.
          </Typography>
        </Grid>
        <Grid item>
          <div className="text-center">
            <div className="sign_up_head">Sign Up</div>
            <a href={url} className="btn github_btn">
              Login with github
              <img src="/assets/github_icon.png" alt="Login with github" />
            </a>
          </div>{' '}
        </Grid>
      </Grid>
      <Paper />
    </div>
  )
}

export const getServerSideProps = withIronSession(async ({ req }) => {
  if (!validateUser(req)) {
    const res = await getLoginURL()
    if (!res.error) return { props: { url: res.url } }
    // todo:redirect to error page
    return {
      redirect: {
        permanent: false,
        destination: '/error',
      },
    }
  }
  return {
    redirect: {
      permanent: false,
      destination: '/dashboard',
    },
  }
}, sessionCongfig)

export default Login
