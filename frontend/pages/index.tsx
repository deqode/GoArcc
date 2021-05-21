import { Button, Paper } from '@material-ui/core'
import { withSentry } from '@sentry/nextjs'
import { NextApiRequest } from 'next'
import useTranslation from 'next-translate/useTranslation'
import { ReactElement } from 'react'

import { getLoginURL } from '../api/rest/fetchUrls'
import BasicLayout from '../components/layouts/BasicLayout'
import PageHead, { Titles } from '../components/PageHead'
import { redirectToErrorPage, redirectToDashboard } from '../utils/redirects'
import { sessionPropsWrapper, validateUser } from '../utils/user'

const Landing = ({ url }: { url: string }): ReactElement => {
  const { t } = useTranslation()

  return (
    <Paper elevation={0}>
      <PageHead title={Titles.WELCOME} />
      <BasicLayout
        heading={t('landing:Welcome to Alfred')}
        subHeading={t(
          'landing:Provide your repo, Select your cloud, and let Alfred do the heavy lifting'
        )}
        component={
          <Button variant="contained" color="primary" href={url}>
            {t('landing:Sign UP')}
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
