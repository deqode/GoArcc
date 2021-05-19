import { Button, CircularProgress, Paper } from '@material-ui/core'
import Head from 'next/head'
import { ReactElement, useContext, useEffect, useState } from 'react'
import UserContext from '../../contexts/UserContext'

import { IronSessionRequest, UserResponse } from '../../interface'
import { getGithubVCSConnection } from '../../api/rest/fetchUrls'
import { validateUser, sessionPropsWrapper } from '../../utils/user'
import { useRouter } from 'next/router'
import { redirectToLandingPage } from '../../utils/redirects'
import BasicLayout from '../../components/layouts/BasicLayout'
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
    <Paper elevation={0}>
      <Head>
        <title>Login to Alfred</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <BasicLayout
        heading={'Let us make your cloud work for you'}
        subHeading={'Connect with Github '}
        component={
          url === '' ? (
            <CircularProgress />
          ) : (
            <Button variant="contained" color="primary" href={url}>
              Connect with github
            </Button>
          )
        }
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
export default Dashboard
