import React, { ReactElement, useEffect } from 'react'
import { useRouter } from 'next/router'
import { sessionCongfig } from '../../../utils/constants'
import { withIronSession } from 'next-iron-session'
import { getVCSConnectionGitHubCallback } from '../../../api/rest/callbacks'
import { validateUser } from '../../../utils/user'
import { UserResponse } from '../../../interface'
import { getAllUserAccounts } from '../../../api/rest/user'

function Callback({ user }: { user: UserResponse }): ReactElement {
  const router = useRouter()

  const { query } = router
  useEffect(() => {
    if (query.code && query.code.length > 0) {
      ;(async () => {
        if (query.code && typeof query.code === 'string' && user.userId !== '') {
          const res = await getAllUserAccounts(user.userId, user.idToken)
          // TODO: gql
          const vcs = await getVCSConnectionGitHubCallback(
            query.code,
            res.accounts[0].id,
            user.idToken
          )
          if (!vcs.error) router.push('/dashboard/tell-us-more')
        } else router.push('/dashboard')
      })()
    } else router.push('/')
  }, [query, router, user])

  return <div></div>
}

export default Callback

export const getServerSideProps = withIronSession(async ({ req }) => {
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
