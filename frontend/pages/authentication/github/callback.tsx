import { CircularProgress } from '@material-ui/core'
import { withSentry } from '@sentry/nextjs'
import { withIronSession } from 'next-iron-session'
import { useRouter } from 'next/router'
import React, { Fragment, ReactElement, useEffect } from 'react'

import { getVCSConnectionGitHubCallback } from '../../../api/rest/callbacks'
import { getAllUserAccounts } from '../../../api/rest/user'
import BasicLayout from '../../../components/layouts/BasicLayout'
import PageHead, { Titles } from '../../../components/PageHead'
import { UserResponse } from '../../../intefaces/interface'
import { sessionCongfig } from '../../../utils/constants'
import { validateUser } from '../../../utils/user'

interface CallbackProps {
  user: UserResponse
}

const Callback = ({ user }: CallbackProps): ReactElement => {
  const router = useRouter()
  const { query } = router
  useEffect(() => {
    if (query.code && query.code.length > 0) {
      ;(async () => {
        if (query.code && query.code && user.userId) {
          const res = await getAllUserAccounts(user.userId, user.idToken)
          if (res.error) {
            router.push('/error?message=user not found')
            return
          }
          const vcs = await getVCSConnectionGitHubCallback(
            query.code as string,
            res.accounts[0].id as string,
            user.idToken
          )
          if (!vcs.error) router.push('/dashboard/tell-us-more')
        } else router.push('/dashboard')
      })()
    } else router.push('/')
  }, [query, router, user])

  return (
    <Fragment>
      <PageHead title={Titles.LOADING} />
      <BasicLayout heading={'Alfred connecting to github'} component={<CircularProgress />} />
    </Fragment>
  )
}

export const getServerSideProps = withSentry(
  withIronSession(async ({ req }) => {
    if (validateUser(req)) {
      return { props: { user: req.session.get('user') } }
    }
    return {
      redirect: {
        permanent: false,
        destination: '/',
      },
    }
  }, sessionCongfig)
)

export default Callback
