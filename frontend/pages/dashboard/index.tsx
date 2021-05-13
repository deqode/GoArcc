import { CircularProgress, Paper } from '@material-ui/core'
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
import BasicLayout from '../../components/layouts/BasicLayout'

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
    <Paper elevation={0}>
      <Head>
        <title>Login to Alfred</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <BasicLayout
        heading={'Let us make your cloud work for you'}
        subHeading={'Connect with Github '}
        component={url === '' ? <CircularProgress /> : <Link href={url}>Connect with github</Link>}
      />
    </Paper>
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
