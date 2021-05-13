import { CircularProgress, Grid, Paper, Typography } from '@material-ui/core'
import Head from 'next/head'
import { sessionCongfig } from '../../utils/constants'
import Link from 'next/link'
import { ReactElement, useContext, useEffect, useState } from 'react'
import { withIronSession } from 'next-iron-session'
import UserContext from '../../contexts/UserContext'

import { UserResponse } from '../../interface'
import { getGithubVCSConnection } from '../../api/rest/fetchUrls'
import { validateUser } from '../../utils/user'
import { useRouter } from 'next/router'
import { redirectToLandingPage } from '../../utils/redirects'
import { withSentry } from '@sentry/nextjs'

export const Dashboard = ({ user }: { user: UserResponse }): ReactElement => {
  const [url, setUrl] = useState<string>('')
  const { setUser } = useContext(UserContext)
  const router = useRouter()
  useEffect(() => {
    if (user.idToken !== '') {
      setUser({
        loggedIn: true,
        idToken: user.idToken,
      })
      ;(async () => {
        const res = await getGithubVCSConnection(user.idToken)
        if (res.error) {
          router.push('/error')
        } else {
          setUrl(res.redirectUrl)
        }
      })()
    } else {
      router.push('/')
    }
  }, [user])

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
            <div className="sign_up_head">Connect</div>
            {url === '' ? (
              <CircularProgress />
            ) : (
              <Link href={url}>
                <a className="btn github_btn">Connect with github</a>
              </Link>
            )}
          </div>
        </Grid>
      </Grid>

      <Paper />
    </div>
  )
}

// TODO:Create common for withSentry & withIronSession
export const getServerSideProps = withSentry(
  withIronSession(async ({ req }) => {
    if (validateUser(req)) {
      return { props: { user: req.session.get('user') } }
    }
    return redirectToLandingPage()
  }, sessionCongfig)
)
export default Dashboard
