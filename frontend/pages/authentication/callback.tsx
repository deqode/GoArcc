import { CircularProgress } from '@material-ui/core'
import { useRouter } from 'next/router'
import React, { Fragment, ReactElement, useEffect } from 'react'

import { getAuth0Callback } from '../../api/rest/callbacks'
import { setUserSession } from '../../api/rest/session'
import BasicLayout from '../../components/layouts/BasicLayout'
import PageHead, { Titles } from '../../components/PageHead'

const Callback = (): ReactElement => {
  const router = useRouter()
  const { query } = router
  useEffect(() => {
    ;(async () => {
      if (query && query.code && query.state) {
        const res = await getAuth0Callback(query.code as string, query.state as string)
        if (!res.error) {
          const sessionRes = await setUserSession(res)
          if (sessionRes.error) {
            router.push('/')
            return
          }
        }
        router.push('/dashboard')
      }
    })()
  }, [router, query])

  return (
    <Fragment>
      <PageHead title={Titles.LOADING} />
      <BasicLayout heading={'You are Logging-In to Alfred'} component={<CircularProgress />} />
    </Fragment>
  )
}
export default Callback
