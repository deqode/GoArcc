import { Grid, Paper, Typography } from '@material-ui/core'
import Head from 'next/head'
import { SERVER, sessionCongfig } from '../utils/constants'
import { useContext } from 'react'
import { withIronSession } from 'next-iron-session'
import { UserContext } from '../Contexts/UserContext'

export default function Login(props: any) {
  const { setUser } = useContext(UserContext)
  const { auth, user } = props
  console.log(user)
  setUser(user)

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
            <a href={auth.url} className="btn github_btn">
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
  const result = await fetch(`${SERVER}/authentication/login`)
  const data = await result.json()
  const auth = data
  if (!req.session.get('user')) {
    return { props: { user: null, auth } }
  }
  return {
    redirect: {
      permanent: false,
      destination: '/dashboard',
    },
  }
}, sessionCongfig)
