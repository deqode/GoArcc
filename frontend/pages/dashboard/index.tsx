import { Button, CircularProgress, Paper } from '@material-ui/core'
import { withSentry } from '@sentry/nextjs'
import { useRouter } from 'next/router'
import { ReactElement, useEffect, useState } from 'react'

import { getGithubVCSConnection } from '../../api/rest/fetchUrls'
import BasicLayout from '../../components/layouts/BasicLayout'
import PageHead, { Titles } from '../../components/PageHead'
import { IronSessionRequest, UserResponse } from '../../interface'
import { redirectToLandingPage } from '../../utils/redirects'
import { validateUser, sessionPropsWrapper } from '../../utils/user'

export const Dashboard = ({ user }: { user: UserResponse }): ReactElement => {
  const [url, setUrl] = useState<string>('')
  const router = useRouter()
  useEffect(() => {
    if (user.idToken !== '') {
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
      <PageHead title={Titles.DASHBOARD} />
      <BasicLayout
        heading={'Let us make your cloud work for you'}
        subHeading={'Connect with Github '}
        component={
          !url.length ? (
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
