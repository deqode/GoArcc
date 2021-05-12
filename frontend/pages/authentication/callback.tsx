import React, { ReactElement, useEffect } from 'react'
import { useRouter } from 'next/router'

import { getAuth0Callback } from '../../api/rest/callbacks'
import { setUserSession } from '../../api/rest/session'
const Callback = (): ReactElement => {
  const router = useRouter()
  const { query } = router
  useEffect(() => {
    ;(async () => {
      if (query && typeof query.code === 'string' && typeof query.state === 'string') {
        const res = await getAuth0Callback(query.code, query.state)
        if (!res.error) {
          const sessionRes = await setUserSession(res)
          if (!sessionRes.error) {
            router.push('/')
            return
          }
        }
        router.push('/dashboard')
      }
    })()
  }, [router, query])

  return <div></div>
}
export default Callback
