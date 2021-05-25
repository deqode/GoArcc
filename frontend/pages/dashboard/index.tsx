import { Button, CircularProgress, Paper } from '@material-ui/core'
import { withSentry } from '@sentry/nextjs'
import useTranslation from 'next-translate/useTranslation'
import { useRouter } from 'next/router'
import { ReactElement, useEffect, useState } from 'react'

import { getGithubVCSConnection } from '../../api/rest/fetchUrls'
import BasicLayout from '../../components/layouts/BasicLayout'
import PageHead, { Titles } from '../../components/PageHead'
import { IronSessionRequest, UserResponse } from '../../intefaces/interface'
import { RedirectReturn, redirectToLandingPage } from '../../utils/redirects'
import { validateUser, sessionPropsWrapper } from '../../utils/user'

interface DashboardProps {
  user: UserResponse
}
export const Dashboard = ({ user }: DashboardProps): ReactElement => {
  const [url, setUrl] = useState<string>('')
  const router = useRouter()
  const { t } = useTranslation()

  useEffect(() => {
    if (user.idToken) {
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
        heading={t('dashboard-index:GREETING')}
        subHeading={t('dashboard-index:CONNECT_GITHUB')}
        component={
          !url.length ? (
            <CircularProgress />
          ) : (
            <Button variant="contained" color="primary" href={url}>
              {t('dashboard-index:CONNECT_GITHUB')}
            </Button>
          )
        }
      />
    </Paper>
  )
}

export const handler = async ({
  req,
}: {
  req: IronSessionRequest
}): Promise<RedirectReturn | { props: DashboardProps }> => {
  if (validateUser(req)) {
    return { props: { user: req.session.get('user') } }
  }
  return redirectToLandingPage()
}

export const getServerSideProps = withSentry(sessionPropsWrapper(handler))
export default Dashboard
