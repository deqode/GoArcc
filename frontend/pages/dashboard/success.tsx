import { Paper } from '@material-ui/core'
import { withSentry } from '@sentry/nextjs'
import { ReactElement } from 'react'

import BasicLayout from '../../components/layouts/BasicLayout'
import PageHead, { Titles } from '../../components/PageHead'
import { IronSessionRequest } from '../../intefaces/interface'
import { redirectToLandingPage } from '../../utils/redirects'
import { sessionPropsWrapper, validateUser } from '../../utils/user'

const Success = (): ReactElement => {
  return (
    <Paper elevation={0}>
      <PageHead title={Titles.DASHBOARD} />

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
export default Success
